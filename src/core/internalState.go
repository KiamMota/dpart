package core
type InternalState struct {
	CurrentDirectory string
	UserDir          string
}

var InterState InternalState = InternalState{}
