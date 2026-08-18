// Package build compiles component repos and launches the ISO repack step.
// v1 is intentionally light: no automated image pipeline yet, Cubic is
// interactive and zed only launches it.
package build

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/zedlum/zed/internal/manifest"
)

// Components runs `make build` in every manifest repo that has a Makefile.
// Repos without one are skipped, not treated as an error.
func Components(m *manifest.Manifest) error {
	if len(m.Repos) == 0 {
		fmt.Println("no repos in manifest, nothing to build")
		return nil
	}
	for _, r := range m.Repos {
		mk := filepath.Join(r.Path, "Makefile")
		if _, err := os.Stat(mk); os.IsNotExist(err) {
			fmt.Printf("%s: no Makefile, skipping\n", r.Path)
			continue
		}
		cmd := exec.Command("make", "build")
		cmd.Dir = r.Path
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		fmt.Printf("%s: make build\n", r.Path)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("%s: %w", r.Path, err)
		}
	}
	return nil
}

// LaunchCubic opens Cubic against baseISO for the manual repack step. Cubic
// is interactive/GUI; zed does not script it and does not install it.
func LaunchCubic(baseISO string) error {
	if _, err := exec.LookPath("cubic"); err != nil {
		return fmt.Errorf(
			"cubic not installed; install it yourself (see https://github.com/PJ-Singh-001/Cubic), " +
				"then run `cubic` and point it at your base Ubuntu Budgie ISO")
	}
	if baseISO != "" {
		if _, err := os.Stat(baseISO); err != nil {
			return fmt.Errorf("base ISO %q: %w", baseISO, err)
		}
	}
	cmd := exec.Command("cubic")
	if err := cmd.Start(); err != nil {
		return err
	}
	fmt.Println("cubic launched; repack the ISO manually, this step is not scripted yet")
	return nil
}
