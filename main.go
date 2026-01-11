package main

import (
	"env-checker/cmd"
	"fmt"
	"github.com/lpernett/godotenv"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "envcheck",
	Short: "validate environmental variables against a schema",
}

func init() {
	rootCmd.AddCommand(
		cmd.CreateCmd(),
		cmd.ValidateCmd(),
		cmd.VersionCmd(),
	)
}

func main() {
	godotenv.Load()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
