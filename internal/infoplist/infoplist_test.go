package infoplist

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"howett.net/plist"
)

func TestConfigureCopy(t *testing.T) {
	tmpDir := t.TempDir()
	plistPath := filepath.Join(tmpDir, "Info.plist")

	initial := map[string]interface{}{
		"CFBundleIdentifier":           "com.roblox.RobloxPlayer",
		"LSMultipleInstancesProhibited": true,
	}

	buf := &bytes.Buffer{}
	if err := plist.NewEncoder(buf).Encode(initial); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(plistPath, buf.Bytes(), 0644); err != nil {
		t.Fatal(err)
	}

	targetBundleID := "com.roblox.RobloxPlayer.instance.12345"
	if err := ConfigureCopy(plistPath, targetBundleID, false); err != nil {
		t.Fatalf("ConfigureCopy failed: %v", err)
	}

	data, err := os.ReadFile(plistPath)
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]interface{}
	if _, err := plist.Unmarshal(data, &result); err != nil {
		t.Fatal(err)
	}

	if result["CFBundleIdentifier"] != targetBundleID {
		t.Errorf("expected bundle ID %s, got %v", targetBundleID, result["CFBundleIdentifier"])
	}

	if result["LSMultipleInstancesProhibited"] != false {
		t.Errorf("expected LSMultipleInstancesProhibited false, got %v", result["LSMultipleInstancesProhibited"])
	}
}
