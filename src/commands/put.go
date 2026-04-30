package commands

import "dpart/core"

func Put(param string, other string) string {
	e := core.Put(param, other)
	if e != nil {return e.Error()}
	return "" 
}
