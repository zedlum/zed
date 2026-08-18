// Package vmctl wraps virsh for a single, already-existing libvirt domain.
// It never creates or provisions a domain.
package vmctl

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func run(args ...string) (string, error) {
	cmd := exec.Command("virsh", args...)
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

func notFoundErr(domain string, out string, err error) error {
	if strings.Contains(out, "failed to get domain") {
		return fmt.Errorf("domain %q does not exist; create it manually first (virt-install or virt-manager), zed only manages existing domains", domain)
	}
	return fmt.Errorf("%w\n%s", err, out)
}

// Start starts an existing domain.
func Start(domain string) error {
	out, err := run("start", domain)
	if err != nil {
		return notFoundErr(domain, out, err)
	}
	fmt.Println(out)
	return nil
}

// Stop gracefully shuts down a running domain.
func Stop(domain string) error {
	out, err := run("shutdown", domain)
	if err != nil {
		return notFoundErr(domain, out, err)
	}
	fmt.Println(out)
	return nil
}

var ipv4Re = regexp.MustCompile(`ipv4\s+([0-9.]+)/\d+`)

// IP returns the domain's first reported IPv4 address.
func IP(domain string) (string, error) {
	out, err := run("domifaddr", domain)
	if err != nil {
		return "", notFoundErr(domain, out, err)
	}
	m := ipv4Re.FindStringSubmatch(out)
	if m == nil {
		return "", fmt.Errorf("no IPv4 address reported for %q yet (still booting?)", domain)
	}
	return m[1], nil
}

// Shell opens an interactive SSH session to the domain.
func Shell(domain, user string) error {
	ip, err := IP(domain)
	if err != nil {
		return err
	}
	cmd := exec.Command("ssh", fmt.Sprintf("%s@%s", user, ip))
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	return cmd.Run()
}

// SnapshotCreate creates a named snapshot.
func SnapshotCreate(domain, name string) error {
	out, err := run("snapshot-create-as", domain, name)
	if err != nil {
		return notFoundErr(domain, out, err)
	}
	fmt.Println(out)
	return nil
}

// SnapshotRevert reverts to a named snapshot.
func SnapshotRevert(domain, name string) error {
	out, err := run("snapshot-revert", domain, name)
	if err != nil {
		return notFoundErr(domain, out, err)
	}
	fmt.Println(out)
	return nil
}
