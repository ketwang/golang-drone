package main

import (
	"github.com/spf13/cobra"
	"os"
	"util/version"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "util-agent",
	}

	rootCmd.AddCommand(version.NewCommand())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
