package path

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

// NewSnippetPath get new snippet path
func NewSnippetPath(targetTime time.Time, title, fileType string) (string, error) {
	dirPath := StoreDirPath()
	createdFileName, err := listDirFileNames(dirPath)
	if err != nil {
		return "", err
	}
	fileName := snippetFileName(targetTime, createdFileName, title, fileType)
	return filepath.Join(dirPath, fileName), nil
}

// StoreDirPath get snippet store directory path
func StoreDirPath() string {
	home, _ := homedir.Dir()
	diaryDirPath := filepath.Join(home, ".config", "lipet", "_post")
	return diaryDirPath
}

// SnippetPath get snippet path to specific ID
func SnippetPath(snippetID string) string {
	fileName := fmt.Sprintf("%s.yaml", snippetID)
	return filepath.Join(StoreDirPath(), fileName)
}

func snippetFileName(targetTime time.Time, createdNames []string, suffix string, fileType string) string {
	dateString := targetTime.Format("2006-01-02")

	var overlapPrefixies []string
	for _, name := range createdNames {
		if strings.HasPrefix(name, dateString) {
			overlapPrefixies = append(overlapPrefixies, name)
		}
	}
	renban := 1 + len(overlapPrefixies)

	snippetFileName := fmt.Sprintf("%s_%s_%s", dateString, fmt.Sprintf("%02d", renban), suffix)
	if fileType != "" {
		snippetFileName = fmt.Sprintf("%s.%s", snippetFileName, fileType)
	}

	return snippetFileName
}

func ListSnippetFiles() ([]string, error) {
	return listDirFileNames(StoreDirPath())
}

func listDirFileNames(dirPath string) ([]string, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("Failed read dir files. %s", err)
	}

	var names []string
	for _, file := range files {
		filename := file.Name()
		name := strings.TrimSuffix(filename, filepath.Ext(filename))
		names = append(names, name)
	}
	return names, nil
}
