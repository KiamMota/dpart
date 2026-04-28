package commands

import "dpart/core"

func Exists(path string) string {
	if core.FsIoExists(path){
		return core.StrTrue
	}
	return core.StrFalse
}
