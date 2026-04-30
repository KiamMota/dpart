package core

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type FileMetadata struct {
	Name string
	Absolute string
	Size string 
	ModifiedAt string
	IsDirectory string
	Permissions string
}

func Put(name string, param string) (error) {
	if(strings.Contains(name, "/")){

		return (fsIoDirCreate(name))
	}
	_, err := fsIoFileCreate(name)
	return err
}

func Remove(name string) error {
	if fsIoExists(name){
		return os.RemoveAll(name)
	}
	return errors.New("file doenst exists")
}

func Get(name string) ([]FileMetadata, error) {

	if name == "./"{
		return fsIoListDir(InterState.CurrentDirectory)
	}

	if(!fsIoExists(name)){
		return nil, nil 
 	 }

	handle := []FileMetadata{}

	fs, e := EntryInspect(name)
	if e != nil { return nil, e} 
	return append(handle, fs), nil

}

func fsIoExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil{
		return false 
	}
		return true
}


func fsIoFileCreate(filename string) (*os.File, error) {
    return os.Create(filename)
}

func fsIoDirCreate(filename string) error {
	err := os.MkdirAll(filename, 0755)
	if err != nil {
		return err
	}
	return err
}

func fsIoDeleteEntry(name string) error {
	err := os.Remove(name)
	if err != nil { 
		return err
	}
	return nil
}



func NormalizePath(base, input string) string {
	var p string

	if filepath.IsAbs(input) {
		p = input
	} else {
		p = filepath.Join(base, input)
	}

	return filepath.Clean(p)
}

func fsIoAppendFile(name string, content string) error {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return err 
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil { return err }
	return nil
}

func fsIoListDir(name string) ([]FileMetadata, error){
	
	e, err := os.Stat(name)
	if err != nil {
		return nil, err
	}
	if(!e.IsDir()) {
		return nil, errors.New("this is not dir") 
	}
	entries, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}
	handle := make([]FileMetadata, 0, len(entries))



	for _, e := range entries {
		meta,_ := EntryInspect(e.Name())

		handle = append(handle, meta)
			

	}
	return handle, nil
}

func EntryInspect(name string) (FileMetadata, error) {
	f, err := os.Stat(name)
	if err != nil {
		return FileMetadata{}, err
	}

	var isDir string 
	abs, err := filepath.Abs(name)
	if err != nil { abs = ""}
	if f.IsDir() {
		isDir = StrTrue 
	}

	return FileMetadata{
		Name:    f.Name(),
		Absolute: abs,
		Size:    strconv.FormatInt(f.Size(), 10),
		IsDirectory:   isDir,
		ModifiedAt: f.ModTime().String(),
	}, nil
}
