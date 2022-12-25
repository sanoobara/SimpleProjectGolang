package files

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

func JoinRunPath(name string) string {
	return path.Join(runPath(), name)
}

func runPath() string {
	ex, er := os.Executable()
	if er != nil {
		panic(er)
	}
	dir := filepath.ToSlash(filepath.Dir(ex))
	return strings.TrimPrefix(dir, filepath.VolumeName(dir))
}
