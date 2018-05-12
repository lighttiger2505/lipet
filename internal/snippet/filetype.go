package snippet

import (
	"strings"
)

type Filetype struct {
	Type      string
	Extension string
}

var Filetypes = []*Filetype{
	&Filetype{Type: "Bash", Extension: "sh"},
	&Filetype{Type: "C", Extension: "c"},
	&Filetype{Type: "C#", Extension: "cs"},
	&Filetype{Type: "C++", Extension: "cpp"},
	&Filetype{Type: "CSS", Extension: "css"},
	&Filetype{Type: "Clojure", Extension: "clj"},
	&Filetype{Type: "CoffeeScript", Extension: "coffee"},
	&Filetype{Type: "D", Extension: "d"},
	&Filetype{Type: "Dockerfile", Extension: ""},
	&Filetype{Type: "Elixir", Extension: "exs"},
	&Filetype{Type: "Erlang", Extension: "erl"},
	&Filetype{Type: "Fish", Extension: "fish"},
	&Filetype{Type: "Go", Extension: "go"},
	&Filetype{Type: "HTML", Extension: "html"},
	&Filetype{Type: "Haml", Extension: "haml"},
	&Filetype{Type: "Haskell", Extension: "hs"},
	&Filetype{Type: "JSON", Extension: "json"},
	&Filetype{Type: "Java", Extension: "java"},
	&Filetype{Type: "JavaScript", Extension: "js"},
	&Filetype{Type: "Kotlin", Extension: "kt"},
	&Filetype{Type: "LaTeX", Extension: "tx"},
	&Filetype{Type: "Lua", Extension: "lua"},
	&Filetype{Type: "Make", Extension: "Makefile"},
	&Filetype{Type: "Markdown", Extension: "md"},
	&Filetype{Type: "Nim", Extension: "nim"},
	&Filetype{Type: "OCaml", Extension: "cma"},
	&Filetype{Type: "Objective-C", Extension: "m"},
	&Filetype{Type: "PHP", Extension: "php"},
	&Filetype{Type: "Perl", Extension: "pl"},
	&Filetype{Type: "Python", Extension: "py"},
	&Filetype{Type: "R", Extension: "r"},
	&Filetype{Type: "Ruby", Extension: "rb"},
	&Filetype{Type: "Rust", Extension: "rs"},
	&Filetype{Type: "SASS", Extension: "sass"},
	&Filetype{Type: "SCSS", Extension: "scss"},
	&Filetype{Type: "SQL", Extension: "sql"},
	&Filetype{Type: "Scala", Extension: "scala"},
	&Filetype{Type: "Shell", Extension: "sh"},
	&Filetype{Type: "Slim", Extension: "slim"},
	&Filetype{Type: "Swift", Extension: "swift"},
	&Filetype{Type: "Terraform", Extension: "tf"},
	&Filetype{Type: "TypeScript", Extension: "ts"},
	&Filetype{Type: "Vimscript", Extension: "vim"},
	&Filetype{Type: "XHTML", Extension: "xhtml"},
	&Filetype{Type: "XML", Extension: "xml"},
	&Filetype{Type: "YAML", Extension: "yml"},
	&Filetype{Type: "ZSH", Extension: "zsh"},
	&Filetype{Type: "reStructuredText", Extension: "rst"},
}

func lowerTrimSpace(val string) string {
	return strings.ToLower(strings.Replace(val, " ", "", -1))
}

func generateFiletypeMap() map[string]*Filetype {
	result := map[string]*Filetype{}
	for _, filetype := range Filetypes {
		lowerTrimFiletype := lowerTrimSpace(filetype.Type)
		result[lowerTrimFiletype] = filetype
	}
	return result
}

func generateFileExtensionMap() map[string]*Filetype {
	result := map[string]*Filetype{}
	for _, filetype := range Filetypes {
		result[filetype.Extension] = filetype
	}
	return result
}

func matchFiletype(val string) *Filetype {
	mapping := generateFiletypeMap()
	lowerTrimFiletype := lowerTrimSpace(val)

	_, has := mapping[lowerTrimFiletype]
	if !has {
		return nil
	}
	return mapping[lowerTrimFiletype]
}

func matchFileExtension(val string) *Filetype {
	mapping := generateFileExtensionMap()
	lowerTrimFiletype := lowerTrimSpace(val)

	_, has := mapping[lowerTrimFiletype]
	if !has {
		return nil
	}
	return mapping[lowerTrimFiletype]
}

func GetFiletype(val string) string {
	matchedFiletype := matchFiletype(val)
	if matchedFiletype != nil {
		return matchedFiletype.Type
	}

	matchedExtension := matchFileExtension(val)
	if matchedExtension != nil {
		return matchedExtension.Type
	}
	return ""
}

func GetFileExtension(val string) string {
	matchedFiletype := matchFiletype(val)
	if matchedFiletype != nil {
		return matchedFiletype.Extension
	}
	return lowerTrimSpace(val)
}
