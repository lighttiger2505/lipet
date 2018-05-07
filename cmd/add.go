// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create new snippet",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: add,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func add(cmd *cobra.Command, args []string) error {
	fmt.Println("add called")

	targetTime := time.Now()
	targetPath, err := diaryPath(targetTime, diaryDirPath())
	if err != nil {
		return err
	}

	// Make directory
	targetDirPath := filepath.Dir(targetPath)
	if !isFileExist(targetDirPath) {
		if err := os.MkdirAll(targetDirPath, 0755); err != nil {
			return fmt.Errorf("Failed make diary dir. %s", err.Error())
		}
	}

	// // Make diary file
	// if !isFileExist(targetPath) {
	// 	if err := makeFile(targetPath); err != nil {
	// 		return fmt.Errorf("Failed make diary file. %s", err.Error())
	// 	}
	// }

	// Open text editor
	editorEnv := os.Getenv("EDITOR")
	if editorEnv == "" {
		editorEnv = "vim"
	}
	err = openEditor(editorEnv, targetPath)
	if err != nil {
		return fmt.Errorf("Failed open editor. %s", err.Error())
	}
	return nil
}

func diaryDirPath() string {
	home, _ := homedir.Dir()
	diaryDirPath := filepath.Join(home, ".config", "lipet", "_post")
	return diaryDirPath
}

func diaryPath(targetTime time.Time, dirPath string) (string, error) {
	year, month, day := targetTime.Date()
	// diaryDirPath := diaryDirPath()
	diaryPath := filepath.Join(
		dirPath,
		fmt.Sprintf("%02d", year),
		fmt.Sprintf("%02d", int(month)),
		fmt.Sprintf("%s.md", fmt.Sprintf("%02d", day)),
	)
	return diaryPath, nil
}

func isFileExist(fPath string) bool {
	_, err := os.Stat(fPath)
	return err == nil || !os.IsNotExist(err)
}

func openEditor(program string, args ...string) error {
	c := exec.Command(program, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

// func makeFile(fPath string) error {
// 	err := ioutil.WriteFile(fPath, []byte(""), 0644)
// 	if err != nil {
// 		return fmt.Errorf("Failed make file. %v", err.Error())
// 	}
// 	return nil
// }
