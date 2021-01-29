package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	appVersion string
	buildTime  string
	goVersion  string
)

var Command = &cobra.Command{
	Use:   "version",
	Short: "show version info",
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("appVersion:", appVersion)
	fmt.Println("buildTime:", buildTime)
	fmt.Println("goVersion:", goVersion)
}
