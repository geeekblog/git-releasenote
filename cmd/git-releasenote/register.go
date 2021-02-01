package main

import (
	"fmt"
	"git-releasenote/cmd/sub_cmd/changelog"
	"git-releasenote/cmd/sub_cmd/version"
	"os"

	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(version.Command)
	rootCmd.AddCommand(changelog.Command)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
