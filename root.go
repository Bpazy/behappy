package really

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// buildVer represents 'really' build version
	buildVer string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "really",
		Short: "TODO",
		Long: `TODO
`,
		Run: func(cmd *cobra.Command, args []string) {
			Run()
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "版本号",
		Long:  `查看 really 的版本号`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(buildVer)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

func Run() {

}
