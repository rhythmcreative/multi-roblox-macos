package infoplist

import (
	"bytes"
	"os"

	"howett.net/plist"
)

// ConfigureCopy updates both the bundle identifier and the multiple instances prohibition.
func ConfigureCopy(path string, bundleIdentifier string, prohibited bool) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var val map[string]interface{}
	dec := plist.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&val); err != nil {
		return err
	}

	if val == nil {
		val = make(map[string]interface{})
	}

	if bundleIdentifier != "" {
		val["CFBundleIdentifier"] = bundleIdentifier
	}
	val["LSMultipleInstancesProhibited"] = prohibited

	buf := &bytes.Buffer{}
	format := dec.Format
	if format == 0 {
		format = plist.BinaryFormat
	}
	enc := plist.NewEncoderForFormat(buf, format)
	if err := enc.Encode(val); err != nil {
		return err
	}

	return os.WriteFile(path, buf.Bytes(), 0644)
}

func SetMultipleInstancesProhibition(path string, prohibited bool) error {
	return ConfigureCopy(path, "", prohibited)
}

func SetBundleIdentifier(path string, bundleIdentifier string) error {
	return ConfigureCopy(path, bundleIdentifier, false)
}
