package main

import (
	"github.com/spf13/cobra"
	"os"
	"util/cmd/agent/serve"
	"util/cmd/agent/ui"
	"util/version"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "util-agent",
	}

	rootCmd.AddCommand(version.NewCommand())

	rootCmd.AddCommand(ui.UiCommand)

	rootCmd.AddCommand(serve.ServeCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
