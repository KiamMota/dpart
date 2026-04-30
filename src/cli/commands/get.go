package commands

import (
	"dpart/core"
	"encoding/json"
)

func Get(param string) string {
	f, err := core.Get(param)

	if f == nil && err == nil { return core.StrNil}
	if err != nil {
		return err.Error()
	}

	if len(f) == 0 {
		return core.StrNil
	}

	data, err := json.MarshalIndent(f, "", " ")
	if err != nil {
		return err.Error()
	}

	return string(data)
}
