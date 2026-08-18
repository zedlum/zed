package vmctl

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// candidate locations checked in addition to PATH, since virtiofsd is
// often installed under libexec rather than a PATH dir.
var virtiofsdFallbackPaths = []string{
	"/usr/libexec/virtiofsd",
	"/usr/lib/qemu/virtiofsd",
}

// CheckVirtiofsd reports whether virtiofsd is present. It never installs
// anything.
func CheckVirtiofsd() error {
	if _, err := exec.LookPath("virtiofsd"); err == nil {
		return nil
	}
	for _, p := range virtiofsdFallbackPaths {
		if _, err := os.Stat(p); err == nil {
			return nil
		}
	}
	return fmt.Errorf("virtiofsd not found on PATH or in %v; install it yourself (e.g. apt install virtiofsd), zed will not do this for you", virtiofsdFallbackPaths)
}

const mountTag = "zedshare"

// injectVirtiofs adds the memoryBacking + filesystem elements needed for a
// virtiofs share of hostDir. Pure string transform, no I/O, so it can be
// tested against real domain XML without touching any actual domain.
// Returns the unmodified doc and alreadyPresent=true if the share already
// exists.
func injectVirtiofs(xmlDoc, hostDir string) (result string, alreadyPresent bool) {
	if strings.Contains(xmlDoc, `dir='`+mountTag+`'`) {
		return xmlDoc, true
	}

	if !strings.Contains(xmlDoc, "<memoryBacking>") {
		xmlDoc = strings.Replace(xmlDoc, "<devices>",
			"<memoryBacking>\n    <source type='memfd'/>\n    <access mode='shared'/>\n  </memoryBacking>\n  <devices>", 1)
	}

	fsDevice := fmt.Sprintf(
		"  <filesystem type='mount' accessmode='passthrough'>\n"+
			"    <driver type='virtiofs' queue='1024'/>\n"+
			"    <source dir='%s'/>\n"+
			"    <target dir='%s'/>\n"+
			"  </filesystem>\n</devices>", hostDir, mountTag)
	xmlDoc = strings.Replace(xmlDoc, "</devices>", fsDevice, 1)

	return xmlDoc, false
}

// SyncInit adds a virtiofs share of hostDir to the domain's XML, if not
// already present. Idempotent. Requires the domain to already exist and
// virtiofsd to already be installed.
func SyncInit(domain, hostDir string) error {
	if err := CheckVirtiofsd(); err != nil {
		return err
	}

	out, err := run("dumpxml", domain)
	if err != nil {
		return notFoundErr(domain, out, err)
	}

	newXML, already := injectVirtiofs(out, hostDir)
	if already {
		fmt.Println("virtiofs share already configured, nothing to do")
		return nil
	}

	tmp, err := os.CreateTemp("", "zed-domain-*.xml")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.WriteString(newXML); err != nil {
		return err
	}
	tmp.Close()

	defOut, err := run("define", tmp.Name())
	if err != nil {
		return fmt.Errorf("virsh define: %w\n%s", err, defOut)
	}
	fmt.Println(defOut)
	fmt.Printf("added virtiofs share %s -> %s. Restart the domain, then inside the guest:\n  mount -t virtiofs %s /mnt/zedlum\n", hostDir, mountTag, mountTag)
	return nil
}
