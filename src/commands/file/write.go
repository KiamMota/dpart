package file

import (
	"path/filepath"
	"dpart/core"
)

func FileWrite(path string) string {
	f, err := core.FsIoFileCreate(path)
	f.Close()
	if err != nil {
		return err.Error()
	}
	abs, _ := filepath.Abs(f.Name())
	return abs
}
