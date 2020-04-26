package cmd

import (
	"github.com/marjamis/kittt/internal/actions"
	"github.com/marjamis/kittt/pkg/k8sConnector"

	"github.com/spf13/cobra"

	// Used for auth
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var allObjects = []string{
	"aws-auth",
	"aws-nodes",
}

// awsAuthCmCmd represents the awsAuthCm command
var getObjects = &cobra.Command{
	Use:   "getObjects",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientset, err := k8sConnector.GenerateClientSet()
		if err != nil {
			panic(err)
		}

		actions.AwsAuthConfigMap(clientset)
		actions.AwsNodePods(clientset)

	},
}

func init() {
	rootCmd.AddCommand(getObjects)
}
