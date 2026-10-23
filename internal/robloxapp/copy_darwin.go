package robloxapp

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/rhythmcreative/multi-roblox-macos/internal/infoplist"
)

func findRobloxApp() (string, error) {
	candidates := []string{
		"/Applications/Roblox.app",
	}

	if home, err := os.UserHomeDir(); err == nil {
		candidates = append(candidates, filepath.Join(home, "Applications", "Roblox.app"))
	}

	for _, p := range candidates {
		if fi, err := os.Stat(p); err == nil && fi.IsDir() {
			return p, nil
		}
	}

	return "", fmt.Errorf("Roblox.app was not found in /Applications or ~/Applications")
}

func NewCopy() (Copy, error) {
	srcApp, err := findRobloxApp()
	if err != nil {
		return Copy{}, err
	}

	destDir, err := copyDestDir()
	if err != nil {
		return Copy{}, fmt.Errorf("failed to create destination dir: %w", err)
	}

	destApp := filepath.Join(destDir, "Roblox.app")

	// 1. APFS clone (instant copy-on-write), falling back to recursive copy
	cloneCmd := exec.Command("cp", "-cR", srcApp, destApp)
	if err := cloneCmd.Run(); err != nil {
		copyCmd := exec.Command("cp", "-R", srcApp, destApp)
		if err := copyCmd.Run(); err != nil {
			_ = os.RemoveAll(destDir)
			return Copy{}, fmt.Errorf("failed to copy Roblox.app: %w", err)
		}
	}

	// 2. Assign unique bundle identifier & allow multiple instances in Info.plist
	uniqueBundleID := fmt.Sprintf("com.roblox.RobloxPlayer.instance.%d.%d", time.Now().UnixNano(), rand.Intn(100000))
	plistPath := filepath.Join(destApp, "Contents", "Info.plist")
	if err := infoplist.ConfigureCopy(plistPath, uniqueBundleID, false); err != nil {
		_ = os.RemoveAll(destDir)
		return Copy{}, fmt.Errorf("failed to configure Info.plist: %w", err)
	}

	// 3. Remove quarantine attributes so Gatekeeper doesn't block modified bundle
	_ = exec.Command("xattr", "-cr", destApp).Run()

	// 4. Crucial for modern macOS (14, 15, 27+): ad-hoc resign the modified app bundle.
	// Without re-signing, RunningBoard/AMFI rejects the modified bundle with Error 257.
	signCmd := exec.Command("codesign", "--force", "--deep", "--sign", "-", destApp)
	if err := signCmd.Run(); err != nil {
		_ = os.RemoveAll(destDir)
		return Copy{}, fmt.Errorf("failed to codesign app copy: %w", err)
	}

	return Copy{
		dir:     destDir,
		appPath: destApp,
	}, nil
}
