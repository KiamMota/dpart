package commands

import "dpart/core"

func Cd(name string) string {
	err := core.ChangeDirectory(name)
	if err != nil { return err.Error()}
	return core.InterState.CurrentDirectory
}
