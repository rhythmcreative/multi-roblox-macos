package robloxapp

import (
	"io"
	"os"
	"path/filepath"

	"github.com/rhythmcreative/multi-roblox-macos/pkg/fspath"
)

type Copy struct {
	dir     string
	appPath string
}

func (b Copy) Path() string {
	if b.appPath != "" {
		return b.appPath
	}
	return filepath.Join(b.dir, "Roblox.app")
}

// Close releases all associated resources with the copy.
func (b Copy) Close() error {
	return os.RemoveAll(b.dir)
}

var _ io.Closer = (*Copy)(nil)

func copyDestDir() (string, error) {
	d, err := fspath.TMPDir.Get()
	if err != nil {
		d = os.TempDir()
	}

	subDir := filepath.Join(d, "multi-roblox")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		return "", err
	}

	return os.MkdirTemp(subDir, "inst_")
}
