package commands

import (
	"dpart/core"
	"path/filepath"
)

func ChangeCurrentDirectory(path string) string {
	if path == "" {
		return "empty path."
	}

	if path[0] == '$' {
		switch path {
		case "$user":
			core.InterState.CurrentDirectory = core.InterState.UserDir
			return core.InterState.CurrentDirectory
		default:
			return "'" + path + "' is not a macro."
		}
	}

	if path == ".." {
		parent := filepath.Dir(core.InterState.CurrentDirectory)
		core.InterState.CurrentDirectory = parent
		return parent
	}

	var newPath string
	if filepath.IsAbs(path) {
		newPath = path
	} else {
		newPath = filepath.Join(core.InterState.CurrentDirectory, path)
	}

	if !core.FsIoExists(newPath) {
		return "path does not exist."
	}

	core.InterState.CurrentDirectory = newPath
	return newPath
}


