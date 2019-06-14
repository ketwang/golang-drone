package ui

import "github.com/spf13/cobra"

var (
	UiCommand = &cobra.Command{
		Use: "stats",
		Args: cobra.ExactArgs(1),
		Long: "show xx stats",
		RunE: uiServe,
	}
)

func init()  {
	UiCommand.Flags().StringP("uuid", "u", "", "uuid string")
}

func uiServe(cmd *cobra.Command, args []string) error {
	return nil
}
