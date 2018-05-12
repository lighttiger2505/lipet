package snippet

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/lighttiger2505/lipet/internal/path"
	yaml "gopkg.in/yaml.v2"
)

// Snippet snippet data
type Snippet struct {
	Hash          string
	Title         string
	FileType      string
	FileExtension string
	Content       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (s *Snippet) GetHashShorthand() string {
	return snippetShorthand(s.Hash)
}

func snippetShorthand(hash string) string {
	return hash[:7]
}

func ValidateSnippetHash(hash string) (bool, error) {
	matched, err := regexp.MatchString("[0-9a-f]{7,40}", hash)
	if err != nil {
		return false, fmt.Errorf("Failed validate snippet hash. Hash:%s %s", hash, err)
	}
	return matched, nil
}

// NewSnippetHash generate snippet hash format of sha1
func NewSnippetHash() string {
	h := sha1.New()
	unixMilliSecStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	h.Write([]byte(unixMilliSecStr))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func isFileExist(fPath string) bool {
	_, err := os.Stat(fPath)
	return err == nil || !os.IsNotExist(err)
}

func write(writer io.Writer, snippet *Snippet) error {
	out, err := yaml.Marshal(snippet)
	if err != nil {
		return fmt.Errorf("Failed marshal snippet. %s", err)
	}

	if _, err = io.WriteString(writer, string(out)); err != nil {
		return fmt.Errorf("Failed write snippet. %s", err)
	}
	return nil
}

// Create new snippet
func Create(snippet *Snippet) error {
	snippetPath := path.SnippetPath(snippet.Hash)
	if isFileExist(snippetPath) {
		return fmt.Errorf("Dose exist snippet. Path:%s", snippetPath)
	}

	// Open new snippet file
	file, err := os.OpenFile(snippetPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Failed create new snippet. %s", err)
	}
	defer func() {
		if cerr := file.Close(); err != nil {
			err = cerr
		}
	}()

	// Write snippet
	if err := write(file, snippet); err != nil {
		return err
	}
	return nil
}

func read(r io.Reader) (*Snippet, error) {
	// Read snippet file
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("Failed read snippet. Error: %s", err.Error())
	}

	// Unmarshal snippet
	snippet := &Snippet{}
	if err := yaml.Unmarshal(b, snippet); err != nil {
		return nil, fmt.Errorf("Failed unmarshal snippet.\nFile:%s\nError:%s ", string(b), err.Error())
	}
	return snippet, nil
}

func Get(hash string) (*Snippet, error) {
	switch len(hash) {
	case 7:
		snip, err := GetShorthand(hash)
		if err != nil {
			return nil, err
		}
		return snip, nil
	case 40:
		snip, err := GetFull(hash)
		if err != nil {
			return nil, err
		}
		return snip, nil
	default:
		return nil, nil
	}
}

// GetShorthand get create snippet data with shorthand snippet hash
func GetShorthand(shorthandHash string) (*Snippet, error) {
	snipFiles, err := path.ListSnippetFiles()
	if err != nil {
		return nil, err
	}

	for _, snipFile := range snipFiles {
		if strings.HasPrefix(snipFile, shorthandHash) {
			snip, err := GetFull(snipFile)
			if err != nil {
				return nil, err
			}
			return snip, nil
		}
	}
	// TODO Return special error for not found snippet
	return nil, nil
}

// GetFull created snippet from specific snippet hash
func GetFull(hash string) (*Snippet, error) {
	// Check snippet is exist
	snippetPath := path.SnippetPath(hash)
	if !isFileExist(snippetPath) {
		// TODO Return special error for not found snippet
		return nil, nil
	}

	// Open snippet file
	file, err := os.OpenFile(snippetPath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := file.Close(); err != nil {
			err = cerr
		}
	}()

	// Read snippee
	snippet, err := read(file)
	if err != nil {
		return nil, err
	}
	return snippet, nil
}

// List created snipets
func List() ([]*Snippet, error) {
	snipFiles, err := path.ListSnippetFiles()
	if err != nil {
		return nil, err
	}

	var snips []*Snippet
	for _, snipFile := range snipFiles {
		snip, err := GetFull(snipFile)
		if err != nil {
			return nil, err
		}
		snips = append(snips, snip)
	}
	return snips, nil
}
