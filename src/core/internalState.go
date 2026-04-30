package core

import "os"
type InternalState struct {
	CurrentDirectory string
	UserHomeDir          string
	Clipboard 			 string
}

var InterState InternalState = InternalState{}

func (s* InternalState) StartInternalState() {
	env, _ := os.Getwd()
	home, _ := os.UserHomeDir()
	s.CurrentDirectory = env
	s.UserHomeDir = home
	s.Clipboard = ""
}

func (s InternalState) GetCurrentDirectory() string {
	return s.CurrentDirectory
}
func (s InternalState) GetUserHomeDirectory() string {
	return s.UserHomeDir
}
func (s InternalState) SetCurrentDirectory(pwd string) {
	s.CurrentDirectory = pwd
}
func (s InternalState) SetClipboard(clip string) {
	s.Clipboard = clip
}
