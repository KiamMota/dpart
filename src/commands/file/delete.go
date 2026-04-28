package file 

import (
	"dpart/core"
)

func FileDelete(file string) string {
	if(!core.FsIoExists(file)){
		return "file dont exists!"
	}

	err := core.FsIoDeleteEntry(file)
	if err != nil { return err.Error()}
	return core.StrTrue
}
