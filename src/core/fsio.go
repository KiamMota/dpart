package core

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type FileMetadata struct {
	Name        string
	Type        string
	Size        string
	ModifiedAt  string
	Permissions string
	Absolute    string
}

func Put(name string, param string) error {
	if strings.Contains(name, "/") {

		return (fsIoDirCreate(name))
	}
	_, err := fsIoFileCreate(name)
	return err
}

func Remove(name string) error {
	if fsIoExists(name) {
		return os.RemoveAll(name)
	}
	return errors.New("file doenst exists")
}

func Get(name string) ([]FileMetadata, error) {

	if name == "./" {
		return fsIoListDir(InterState.CurrentDirectory)
	}

	if !fsIoExists(name) {
		return nil, nil
	}

	handle := []FileMetadata{}

	fs, e := EntryInspect(name)
	if e != nil {
		return nil, e
	}
	return append(handle, fs), nil
}

func ChangeDirectory(name string) error {
 if !fsIoExists(name) {
	 return errors.New("Directory doenst exists: " + name)
 }

 abs, err := filepath.Abs(name)
 if err != nil {return err}

 InterState.SetCurrentDirectory(abs)
 return nil
}

func fsIoExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
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

	if _, err := f.WriteString(content); err != nil {
		return err
	}
	return nil
}

func fsIoListDir(name string) ([]FileMetadata, error) {

	e, err := os.Stat(name)
	if err != nil {
		return nil, err
	}
	if !e.IsDir() {
		return nil, errors.New("this is not dir")
	}
	entries, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}
	handle := make([]FileMetadata, 0, len(entries))

	for _, e := range entries {
		meta, _ := EntryInspect(e.Name())

		handle = append(handle, meta)

	}
	return handle, nil
}

func EntryInspect(name string) (FileMetadata, error) {
	f, err := os.Stat(name)
	if err != nil {
		return FileMetadata{}, err
	}

	abs, err := filepath.Abs(name)
	if err != nil {
		abs = ""
	}


	return FileMetadata{
		Name:        f.Name(),
		Type: fileTypeEntry(name),
		Absolute:    abs,
		Size:        strconv.FormatInt(f.Size(), 10),
		ModifiedAt:  f.ModTime().String(),
	}, nil
}

func fileTypeEntry(filePath string) string {
	info, err := os.Stat(filePath)
	if err != nil {
		return "Unknown"
	}

	if info.IsDir() {
		return "Directory"
	}

	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {

	case ".go":
		return "Golang"
	case ".c", ".h":
		return "C"
	case ".cpp", ".cc", ".cxx", ".hpp":
		return "C++"
	case ".js":
		return "JavaScript"
	case ".ts":
		return "TypeScript"
	case ".py":
		return "Python"
	case ".java":
		return "Java"
	case ".rs":
		return "Rust"
	case ".php":
		return "PHP"
	case ".rb":
		return "Ruby"
	case ".cs":
		return "C#"

	case ".txt":
		return "Text"
	case ".md":
		return "Markdown"
	case ".json":
		return "JSON"
	case ".xml":
		return "XML"
	case ".yaml", ".yml":
		return "YAML"
	case ".toml":
		return "TOML"

	case ".html", ".htm":
		return "HTML"
	case ".css":
		return "CSS"

	case ".png":
		return "PNG Image"
	case ".jpg", ".jpeg":
		return "JPEG"
	case ".gif":
		return "GIF"
	case ".webp":
		return "WebP"
	case ".svg":
		return "SVG"

	case ".mp3":
		return "MP3"
	case ".wav":
		return "WAV"
	case ".ogg":
		return "OGG"

	case ".mp4":
		return "MP4"
	case ".mkv":
		return "MKV" 
	case ".avi":
		return "AVI"

	case ".zip":
		return "ZIP Archive"
	case ".rar":
		return "RAR Archive"
	case ".tar":
		return "TAR Archive"
	case ".gz":
		return "GZIP Archive"

	case ".exe":
		return "Executable"
	case ".bin":
		return "Binary"
	case ".so":
		return "Shared Library"
	case ".dll":
		return "DLL"

	case ".csv":
		return "CSV"
	case ".db", ".sqlite":
		return "Database"

	default:
		return "Unknown"
	}
}
