package cmd

import (
	"fmt"

	"github.com/marjamis/kittt/internal/actions"

	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Aliases: []string{
		"a",
	},
	Use:   "all",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running through getting all the files and running all the tests that this application will perform...")

		// TODO This should like be auto completed maybe on creation of array they have to register as well, similar to the thing to categories?
		actionTypes := []*actions.ActionTypes{
			&actions.GetCategories,
			&actions.TestCategories,
		}

		for _, actionType := range actionTypes {
			actionType.RunThroughAll()
		}
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
