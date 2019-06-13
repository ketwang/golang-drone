package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// eg: 0.0.1
	Version string = "version not defined, please run ./build.sh to build"

	// built at
	BuildTime string

	// git branch
	GitBranch string

	// git commit
	GitCommit string

	// Go version
	GoVersion string
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Show version info",
		Aliases: []string{"ver"},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("version:   ", Version)
			fmt.Println("git sha:   ", GitCommit)
			fmt.Println("git branch:	", GitBranch)
			fmt.Println("go version:", GoVersion)
			fmt.Println("built:     ", BuildTime)
			return nil
		},
	}
}
