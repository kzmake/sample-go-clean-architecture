package command

import (
	"github.com/spf13/cobra"

	"gca/service/hoge/infrastructure/cli"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: cli.ListResources,
}

func init() {
	resourceCmd.AddCommand(listCmd)
}
