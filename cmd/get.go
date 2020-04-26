package cmd

import (
	"fmt"
	"log"

	"github.com/marjamis/kittt/internal/actions"
	"github.com/marjamis/kittt/pkg/k8sConnector"

	"github.com/spf13/cobra"
	// Used for auth

	_ "k8s.io/client-go/plugin/pkg/client/auth"
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
			clientset, err := k8sConnector.GenerateClientSet()
			if err != nil {
				panic(err)
			}
			for _, item := range actions.GetCategories["all"] {
				fmt.Printf("Category: %s - Values: %s\n", "all", item.Name)
				item.Function(clientset)
			}
		},
	}

	getCategoryCmd = &cobra.Command{
		Use:   "category",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			category := args[0]
			if _, ok := actions.GetCategories[category]; !ok {
				var keys []string
				for k := range actions.GetCategories {
					keys = append(keys, k)
				}
				log.Panicf("This isn't a valid category please choose all (for all objects) or select from one of these categories: %v", keys)
			}

			clientset, err := k8sConnector.GenerateClientSet()
			if err != nil {
				panic(err)
			}
			for _, item := range actions.GetCategories[category] {
				fmt.Printf("Category: %s - Values: %s\n", category, item.Name)
				item.Function(clientset)
			}
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
