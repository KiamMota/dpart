package core

const (
	StrTrue = "True"
	StrFalse = "False"
	StrNil = "Nil"
)

func MacroTable(macro string) string {
	switch macro {
	case "pwd":
		return InterState.GetCurrentDirectory()
	case "home":
		return InterState.GetUserHomeDirectory()
	default: return macro
	}
}

const (
	TypeDirectory = "Directory"
	TypeUnknown   = "Unknown"

	TypeGolang     = "Golang"
	TypeC          = "C"
	TypeCPP        = "C++"
	TypeJS         = "JavaScript"
	TypeTS         = "TypeScript"
	TypePython     = "Python"
	TypeJava       = "Java"
	TypeRust       = "Rust"
	TypePHP        = "PHP"
	TypeRuby       = "Ruby"
	TypeCSharp     = "C#"

	TypeText       = "Text"
	TypeMarkdown   = "Markdown"
	TypeJSON       = "JSON"
	TypeXML        = "XML"
	TypeYAML       = "YAML"
	TypeTOML       = "TOML"

	TypeHTML       = "HTML"
	TypeCSS        = "CSS"

	TypePNG        = "PNG"
	TypeJPEG       = "JPEG"
	TypeGIF        = "GIF"
	TypeWEBP       = "WebP"
	TypeSVG        = "SVG"

	TypeMP3        = "MP3"
	TypeWAV        = "WAV"
	TypeOGG        = "OGG"

	TypeMP4        = "MP4"
	TypeMKV        = "MKV"
	TypeAVI        = "AVI"

	TypeZIP        = "ZIP Archive"
	TypeRAR        = "RAR Archive"
	TypeTAR        = "TAR Archive"
	TypeGZ         = "GZIP Archive"

	TypeEXE        = "Executable"
	TypeBIN        = "Binary"
	TypeSO         = "Linux Shared Library"
	TypeDLL        = "DLL"

	TypeCSV        = "CSV"
	TypeDatabase   = "Database"
)
