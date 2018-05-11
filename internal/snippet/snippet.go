package snippet

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
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
	return uuid.Must(uuid.NewRandom()).String()
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

// Random number state.
// We generate random temporary file names so that there's a good
// chance the file doesn't exist yet - keeps the number of tries in
// TempFile to a minimum.
var rand uint32
var randmu sync.Mutex

func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}

func nextSuffix() string {
	randmu.Lock()
	r := rand
	if r == 0 {
		r = reseed()
	}
	r = r*1664525 + 1013904223 // constants from Numerical Recipes
	rand = r
	randmu.Unlock()
	return strconv.Itoa(int(1e9 + r%1e9))[1:]
}

func TempFile(dir, prefix, fileType string) (f *os.File, err error) {
	if dir == "" {
		dir = os.TempDir()
	}

	nconflict := 0
	for i := 0; i < 10000; i++ {
		fname := fmt.Sprintf("%s%s.%s", prefix, nextSuffix(), fileType)
		name := filepath.Join(dir, fname)
		f, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
		if os.IsExist(err) {
			if nconflict++; nconflict > 10 {
				randmu.Lock()
				rand = reseed()
				randmu.Unlock()
			}
			continue
		}
		break
	}
	return
}
