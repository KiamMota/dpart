package core

import "os"
type InternalState struct {
	CurrentDirectory string
	UserDir          string
}

var InterState InternalState = InternalState{}

func (s* InternalState) StartInternalState() {
	env, _ := os.Getwd()
	home, _ := os.UserHomeDir()
	s.CurrentDirectory = env
	s.UserDir = home
}
