package commands

import (
	"dpart/core"
	"path/filepath"
)

func ChangeCurrentDirectory(path string) string {

	if path[0] == '$' {
		switch path {
		case "$user":
			core.InterState.CurrentDirectory = core.InterState.UserDir

		default:
			return "'" + path + "' is not a macro."
		}

		return core.InterState.CurrentDirectory
	}
	if !core.FsIoExists(path)	{
		return "path dont exists."
	}

	newPath := filepath.Join(core.InterState.CurrentDirectory, path)
	core.InterState.CurrentDirectory = newPath

	if path == ".." {
		parent := filepath.Dir(core.InterState.CurrentDirectory)

		if parent == core.InterState.CurrentDirectory {
			return core.InterState.CurrentDirectory
		}

		core.InterState.CurrentDirectory = parent
		return core.InterState.CurrentDirectory
	}

	return core.InterState.CurrentDirectory
}


