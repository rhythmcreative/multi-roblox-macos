package urlhandler

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreServices
#include <CoreServices/CoreServices.h>
#include <stdlib.h>

int setHandler(const char *bundleIdentifier, const char *urlScheme) {
    CFStringRef cfUrlScheme = CFStringCreateWithCString(NULL, urlScheme, kCFStringEncodingUTF8);
    CFStringRef cfBundleIdentifier = CFStringCreateWithCString(NULL, bundleIdentifier, kCFStringEncodingUTF8);

    OSStatus status = LSSetDefaultHandlerForURLScheme(cfUrlScheme, cfBundleIdentifier);

    CFRelease(cfUrlScheme);
    CFRelease(cfBundleIdentifier);

    return (status == noErr) ? 0 : 1;
}

int checkHandler(const char *bundleIdentifier, const char *urlScheme) {
    CFStringRef cfUrlScheme = CFStringCreateWithCString(NULL, urlScheme, kCFStringEncodingUTF8);
    CFStringRef cfBundleIdentifier = CFStringCreateWithCString(NULL, bundleIdentifier, kCFStringEncodingUTF8);

    CFStringRef currentHandler = LSCopyDefaultHandlerForURLScheme(cfUrlScheme);
    if (!currentHandler) {
        CFRelease(cfUrlScheme);
        CFRelease(cfBundleIdentifier);
        return 1;
    }

    Boolean isDefault = CFStringCompare(currentHandler, cfBundleIdentifier, 0) == kCFCompareEqualTo;

    CFRelease(currentHandler);
    CFRelease(cfUrlScheme);
    CFRelease(cfBundleIdentifier);

    return isDefault ? 0 : 1;
}
*/
import "C"
import (
	"os/exec"
	"unsafe"
)

const (
	ROBLOX_BUNDLE_IDENTIFIER      = "com.roblox.RobloxPlayer"
	DEFAULT_APP_BUNDLE_IDENTIFIER = "com.rhythmcreative.multiroblox"
	LEGACY_APP_BUNDLE_IDENTIFIER  = "com.insadem.multiroblox"
	SCHEME_ROBLOX_PLAYER          = "roblox-player"
	SCHEME_ROBLOX                 = "roblox"
)

func Set(bundleIdentifier, urlScheme string) bool {
	cBundleIdentifier := C.CString(bundleIdentifier)
	defer C.free(unsafe.Pointer(cBundleIdentifier))

	cURLScheme := C.CString(urlScheme)
	defer C.free(unsafe.Pointer(cURLScheme))

	result := C.setHandler(cBundleIdentifier, cURLScheme)
	return result == 0
}

func Check(bundleIdentifier, urlScheme string) bool {
	cBundleIdentifier := C.CString(bundleIdentifier)
	defer C.free(unsafe.Pointer(cBundleIdentifier))

	cURLScheme := C.CString(urlScheme)
	defer C.free(unsafe.Pointer(cURLScheme))

	result := C.checkHandler(cBundleIdentifier, cURLScheme)
	return result == 0
}

func SetAll(bundleIdentifier string) bool {
	p1 := Set(bundleIdentifier, SCHEME_ROBLOX_PLAYER)
	p2 := Set(bundleIdentifier, SCHEME_ROBLOX)
	return p1 || p2
}

func RestoreDefault() {
	Set(ROBLOX_BUNDLE_IDENTIFIER, SCHEME_ROBLOX_PLAYER)
	Set(ROBLOX_BUNDLE_IDENTIFIER, SCHEME_ROBLOX)
}

// RegisterBundle registers an application bundle with LaunchServices via lsregister.
func RegisterBundle(bundlePath string) error {
	lsregister := "/System/Library/Frameworks/CoreServices.framework/Frameworks/LaunchServices.framework/Support/lsregister"
	return exec.Command(lsregister, "-f", bundlePath).Run()
}
