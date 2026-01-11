package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var (
	version = "1.0.0"
	commit  = "dev"
	date    = time.Now().Format("2006-01-02 15:04:05")
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("envcheck version %s\n", version)
		fmt.Printf("commit: %s\n", commit)
		fmt.Printf("built: %s\n", date)
	},
}

func VersionCmd() *cobra.Command {
	return versionCmd
}