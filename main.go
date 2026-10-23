package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/rhythmcreative/multi-roblox-macos/internal/deeplink"
	"github.com/rhythmcreative/multi-roblox-macos/internal/macosapp"
	"github.com/rhythmcreative/multi-roblox-macos/internal/robloxapp"
	"github.com/rhythmcreative/multi-roblox-macos/internal/syncbreaker"
	"github.com/rhythmcreative/multi-roblox-macos/internal/urlhandler"
)

func init() {
	runtime.LockOSThread()
}

func findEnclosingAppBundle() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	realExe, err := filepath.EvalSymlinks(exe)
	if err == nil {
		exe = realExe
	}
	idx := strings.Index(exe, ".app/Contents/MacOS/")
	if idx != -1 {
		return exe[:idx+4]
	}
	return ""
}

func main() {
	appBundle := findEnclosingAppBundle()
	bundleID := urlhandler.DEFAULT_APP_BUNDLE_IDENTIFIER

	if appBundle != "" {
		_ = urlhandler.RegisterBundle(appBundle)
	}

	// Register as default handler for roblox-player & roblox
	urlhandler.Set(bundleID, urlhandler.SCHEME_ROBLOX_PLAYER)
	urlhandler.Set(bundleID, urlhandler.SCHEME_ROBLOX)

	// If the bundle ID is still roblox and not ours, try legacy bundle ID as well
	if urlhandler.Check(urlhandler.ROBLOX_BUNDLE_IDENTIFIER, urlhandler.SCHEME_ROBLOX_PLAYER) {
		urlhandler.Set(urlhandler.LEGACY_APP_BUNDLE_IDENTIFIER, urlhandler.SCHEME_ROBLOX_PLAYER)
		urlhandler.Set(urlhandler.LEGACY_APP_BUNDLE_IDENTIFIER, urlhandler.SCHEME_ROBLOX)
	}

	// Continuous background watcher to unlink /RobloxPlayerUniq
	watcher := syncbreaker.StartWatcher(250 * time.Millisecond)

	var copiesMu sync.Mutex
	var activeCopies []io.Closer

	addCopy := func(c io.Closer) {
		copiesMu.Lock()
		activeCopies = append(activeCopies, c)
		copiesMu.Unlock()
	}

	cleanupDone := false
	var cleanupMu sync.Mutex
	cleanup := func() {
		cleanupMu.Lock()
		if cleanupDone {
			cleanupMu.Unlock()
			return
		}
		cleanupDone = true
		cleanupMu.Unlock()

		log.Println("Restoring default URL handler and cleaning up...")
		watcher.Stop()
		urlhandler.RestoreDefault()

		copiesMu.Lock()
		for _, c := range activeCopies {
			_ = c.Close()
		}
		activeCopies = nil
		copiesMu.Unlock()
	}

	// Handle termination signals gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		<-sigChan
		cleanup()
		os.Exit(0)
	}()

	// Listen for Roblox URLs triggered from the browser
	go func() {
		for url := range deeplink.Handler() {
			log.Printf("Received Roblox launch URL: %s\n", url)

			// Pulse syncbreaker before and during startup
			syncbreaker.Break()
			syncbreaker.Pulse(4*time.Second, 100*time.Millisecond)

			appCopy, err := robloxapp.NewCopy()
			if err != nil {
				macosapp.Notification(fmt.Sprintf("Failed to clone Roblox: %v", err))
				continue
			}

			addCopy(appCopy)

			// Launch with -n (new instance) and pass the game URL
			cmd := exec.Command("open", "-n", "-a", appCopy.Path(), url)
			if err := cmd.Run(); err != nil {
				macosapp.Notification(fmt.Sprintf("Failed to launch Roblox: %v", err))
				continue
			}

			// Pulse again to ensure the newly spawned instance's semaphore is unlinked
			syncbreaker.Pulse(3*time.Second, 100*time.Millisecond)
		}
	}()

	macosapp.Run(cleanup)
}
