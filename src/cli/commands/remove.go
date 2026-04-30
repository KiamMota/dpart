package commands

import "dpart/core"

func Remove(param string) string {
	err := core.Remove(param)
	if err != nil { return err.Error()}
	return core.StrTrue
}
