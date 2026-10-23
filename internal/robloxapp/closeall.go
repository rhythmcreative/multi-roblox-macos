package robloxapp

import (
	"os"
	"os/exec"

	"github.com/rhythmcreative/multi-roblox-macos/pkg/ps"
)

func CloseAll() {
	processes, err := ps.Processes()
	if err == nil {
		for _, process := range processes {
			pName := process.Executable()
			if pName == "RobloxPlayer" || pName == "RobloxCrashHandler" {
				p, err := os.FindProcess(process.Pid())
				if err == nil {
					_ = p.Kill()
				}
			}
		}
	}

	// Backup killall to catch any edge cases
	_ = exec.Command("killall", "-9", "RobloxPlayer").Run()
}
