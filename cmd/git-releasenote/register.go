package main

import (
	"git-releasenote/cmd/sub_cmd/changelog"
	"git-releasenote/cmd/sub_cmd/version"

	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(version.Command)
	rootCmd.AddCommand(changelog.Command)
	rootCmd.Execute()
}
