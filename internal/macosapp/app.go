package macosapp

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "app.h"
*/
import "C"

import (
	"github.com/rhythmcreative/multi-roblox-macos/internal/robloxapp"
)

var cleanCallback func()

//export goAppClean
func goAppClean() {
	if cleanCallback != nil {
		cleanCallback()
	}
}

//export goCloseAll
func goCloseAll() {
	robloxapp.CloseAll()
}

func Run(clean func()) {
	cleanCallback = clean
	C.RunNativeApp()
}
