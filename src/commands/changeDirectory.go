package commands

import (
	"dpart/core"
	"path/filepath"
)

func ChangeCurrentDirectory(path string) string {
	if !core.FsIoExists(path)	{
		return "path dont exists."
	}
	if path[0] == '$' {
		switch path {
		case "$user":
			core.InterState.CurrentDirectory = core.InterState.UserDir

		default:
			return "'" + path + "' is not a macro."
		}

		return core.InterState.CurrentDirectory
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


