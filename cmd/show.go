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

	"github.com/lighttiger2505/lipet/internal/snippet"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show snippet description",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: show,
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func show(cmd *cobra.Command, args []string) error {
	// Check required
	if len(args) < 1 {
		return fmt.Errorf("Requirements arg. Prease input snippet hash")
	}

	// Validate snippet hash
	hash := args[0]
	result, err := snippet.ValidateSnippetHash(hash)
	if err != nil {
		return err
	}
	if !result {
		return fmt.Errorf("Invalid snippt hash. Hash:%s", hash)
	}

	snip, err := snippet.Get(hash)
	if err != nil {
		return err
	}

	base := `%s
Title: %s
FileType: %s
CreatedAt: %v
UpdatedAt: %v

%s`
	out := fmt.Sprintf(
		base,
		snip.Hash,
		snip.Title,
		snip.FileType,
		snip.CreatedAt,
		snip.UpdatedAt,
		snip.Content,
	)

	fmt.Println(out)
	return nil
}
