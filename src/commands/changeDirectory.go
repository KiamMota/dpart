package commands

import (
	"dpart/core"
	"strings"
)

func ChangeCurrentDirectory(path string) string {
	if path == "" {
		return "empty path"
	}

	// macros
	if strings.HasPrefix(path, "$") {
		switch path {
		case "$user":
			core.InterState.CurrentDirectory = core.InterState.UserDir
			return core.InterState.CurrentDirectory
		default:
			return "unknown macro: " + path
		}
	}

	newPath := core.NormalizePath(core.InterState.CurrentDirectory, path)

	if !core.FsIoExists(newPath) {
		return "path does not exist"
	}

	core.InterState.CurrentDirectory = newPath
	return newPath
}
