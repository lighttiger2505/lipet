package path

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

func NewSnippetPath(targetTime time.Time) (string, error) {
	dirPath := storeDirPath()
	createdFileName, err := listDirFileNames(dirPath)
	if err != nil {
		return "", err
	}
	fileName := snippetFileName(targetTime, createdFileName, "", "")
	return filepath.Join(dirPath, fileName), nil
}

func storeDirPath() string {
	home, _ := homedir.Dir()
	diaryDirPath := filepath.Join(home, ".config", "lipet", "_post")
	return diaryDirPath
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

func listDirFileNames(dirPath string) ([]string, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("Failed read dir files. %s", err)
	}

	var names []string
	for _, file := range files {
		names = append(names, file.Name())
	}
	return names, nil
}
