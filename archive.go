// Package archive provides a collection of archive extensions
// and a utility function to check if given path is an archive
package archive

import (
	"path"
	"strings"
)

var extensions = []string{
	"7z", "a", "apk", "ar", "cab", "cpio", "deb",
	"dmg", "egg", "epub", "iso", "jar", "mar", "pea",
	"rar", "s7z", "shar", "tar", "tbz2", "tgz", "tlz",
	"war", "whl", "xpi", "zip", "zipx",
}

// Extensions is the extensions for different archive types
var Extensions map[string]bool

func init() {
	Extensions = map[string]bool{}
	for _, ext := range extensions {
		Extensions[ext] = true
	}
}

// Is checks if a path is an archive
func Is(p string) bool {
	ext := strings.TrimLeft(path.Ext(p), ".")
	return Extensions[strings.ToLower(ext)]
}
