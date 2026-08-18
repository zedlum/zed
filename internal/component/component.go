// Package component clones and checks the status of manifest repos.
// Paths are resolved relative to the current working directory.
package component

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/zedlum/zed/internal/manifest"
)

// Status is one repo's sync state.
type Status struct {
	Repo    manifest.Repo
	State   string // missing, clean, diverged
	HeadRef string
}

func git(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

// Sync clones missing repos and checks out each to its pinned version.
func Sync(m *manifest.Manifest) error {
	for _, r := range m.Repos {
		if _, err := os.Stat(r.Path); os.IsNotExist(err) {
			fmt.Printf("cloning %s -> %s\n", r.URL, r.Path)
			if out, err := git(".", "clone", r.URL, r.Path); err != nil {
				return fmt.Errorf("clone %s: %w\n%s", r.Path, err, out)
			}
		} else {
			if out, err := git(r.Path, "fetch", "--all"); err != nil {
				return fmt.Errorf("fetch %s: %w\n%s", r.Path, err, out)
			}
		}
		if out, err := git(r.Path, "checkout", r.Version); err != nil {
			return fmt.Errorf("checkout %s@%s: %w\n%s", r.Path, r.Version, err, out)
		}
		fmt.Printf("%s @ %s\n", r.Path, r.Version)
	}
	return nil
}

// Check reports each repo's sync state without changing anything.
func Check(m *manifest.Manifest) ([]Status, error) {
	var statuses []Status
	for _, r := range m.Repos {
		if _, err := os.Stat(r.Path); os.IsNotExist(err) {
			statuses = append(statuses, Status{Repo: r, State: "missing"})
			continue
		}
		head, err := git(r.Path, "rev-parse", "HEAD")
		if err != nil {
			return nil, fmt.Errorf("rev-parse %s: %w", r.Path, err)
		}
		pinned, err := git(r.Path, "rev-parse", r.Version)
		if err != nil {
			// pinned ref not resolvable locally; report as diverged rather than failing the whole check
			statuses = append(statuses, Status{Repo: r, State: "diverged", HeadRef: head})
			continue
		}
		state := "clean"
		if head != pinned {
			state = "diverged"
		}
		statuses = append(statuses, Status{Repo: r, State: state, HeadRef: head})
	}
	return statuses, nil
}
