package commands

import (
	"dpart/core"
	"encoding/json"
)

func Inspect(entry string) string {
	info, err := core.EntryInspect(entry)
	if err != nil {
		return err.Error()
	}

	b, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(b)
}
