// Package manifest holds the embedded component repo list.
package manifest

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed zed.yaml
var raw []byte

// Repo is one entry in the manifest.
type Repo struct {
	Path    string `yaml:"path"`
	URL     string `yaml:"url"`
	Version string `yaml:"version"`
}

// Manifest is the full embedded repo list.
type Manifest struct {
	Repos []Repo `yaml:"repos"`
}

// Load parses the embedded zed.yaml.
func Load() (*Manifest, error) {
	var m Manifest
	if err := yaml.Unmarshal(raw, &m); err != nil {
		return nil, err
	}
	return &m, nil
}
