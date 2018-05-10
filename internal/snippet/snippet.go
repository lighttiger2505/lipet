package snippet

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/lighttiger2505/lipet/internal/path"
	yaml "gopkg.in/yaml.v2"
)

// Snippet snippet data
type Snippet struct {
	ID        string
	Title     string
	FileType  string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewSnippetID generate snippet ID format of sha256
func NewSnippetID() string {
	h := sha256.New()
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
	snippetPath := path.SnippetPath(snippet.ID)
	if isFileExist(snippetPath) {
		return fmt.Errorf("Dose exist snippet file. Path:%s", snippetPath)
	}

	// Open new snippet file
	file, err := os.OpenFile(snippetPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Failed create new snippet file. %s", err)
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

// Get created snippet from specific ID
func Get(snippetID string) (*Snippet, error) {
	// Check snippet is exist
	snippetPath := path.SnippetPath(snippetID)
	if !isFileExist(snippetPath) {
		return nil, fmt.Errorf("Not found snippet. Path:%s", snippetPath)
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
func List() []*Snippet {
	return []*Snippet{}
}
