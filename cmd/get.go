package cmd

import (
	"fmt"

	"github.com/marjamis/kittt/internal/actions"

	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Performs Kubernetes API calls to return data about the cluster",
		Long:  `To aid in troubleshooting the get call will go through a list, either an existing category or all options and provide the output to be used for troubleshooting.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := getValidate(); err != nil {
				fmt.Println("TODO STUB for an error")
			}
		},
	}

	getAllCmd = &cobra.Command{
		Use:   "all",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			actions.GetCategories.RunThroughAll()
		},
	}

	getCategoryCmd = &cobra.Command{
		Use:   "category",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			//TODO isSet() before trying to pass to the function (part of validate?)
			actions.GetCategories.RunThroughCategory(args[0])
		},
	}
)

func init() {
	getCmd.AddCommand(getAllCmd)
	getCmd.AddCommand(getCategoryCmd)
	rootCmd.AddCommand(getCmd)
}

func getValidate() error {
	return nil
}
