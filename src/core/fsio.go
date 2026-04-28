package core

import (
	"os"
	"path/filepath"
	"strconv"
)

type FileMetadata struct {
	Name string
	Size string 
	ModifiedAt string
	IsDirectory string
	Permissions string
}

func FsIoExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil{
		return false 
	}
		return true
}

func FsIoFileCreate(filename string) (*os.File, error) {
    return os.Create(filename)
}

func FsIoDirCreate(filename string) error {
	err := os.MkdirAll(filename, 0755)
	if err != nil {
		return err
	}
	return err
}

func FsIoDeleteEntry(name string) error {
	err := os.Remove(name)
	if err != nil { 
		return err
	}
	return nil
}

func FsIoMoveToTrash(name string) error {
	trashDir := filepath.Join(os.Getenv("HOME"), ".trash")

	if err := FsIoDirCreate(trashDir); err != nil {
		return err
	}

	base := filepath.Base(name)
	dest := filepath.Join(trashDir, base)

	if err := os.Rename(name, dest); err != nil {
		return nil
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

func FsIoAppendFile(name string, content string) error {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return err 
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil { return err }
	return nil
}

func EntryInspect(name string) (FileMetadata, error) {
	f, err := os.Stat(name)
	if err != nil {
		return FileMetadata{}, err
	}

	var isDir string 
	if f.IsDir() {
		isDir = StrTrue 
	}

	return FileMetadata{
		Name:    f.Name(),
		Size:    strconv.FormatInt(f.Size(), 10),
		IsDirectory:   isDir,
		ModifiedAt: f.ModTime().String(),
	}, nil
}
