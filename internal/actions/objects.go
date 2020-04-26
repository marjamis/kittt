package actions

import (
	"context"

	"github.com/marjamis/kittt/internal/output"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	// Used for auth
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type Data struct {
	Name       string
	Function   func(*kubernetes.Clientset)
	Categories []string
}

var (
	GetCategories = map[string][]*Data{}

	awsAuthConfigMap = &Data{
		Name: "aws-auth",
		Function: func(clientset *kubernetes.Clientset) {
			namespace := "kube-system"
			name := "aws-auth"
			results, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
			if err != nil {
				panic(err.Error())
			}

			output.FormatStdout(results)
			if err != nil {
				panic(err.Error())
			}
		},
		Categories: []string{
			"all",
			"dns",
		},
	}

	awsNodePods = &Data{
		Name: "aws-node",
		Function: func(clientset *kubernetes.Clientset) {
			namespace := "kube-system"
			name := "aws-auth"
			results, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
			if err != nil {
				panic(err.Error())
			}

			output.FormatStdout(results)
			if err != nil {
				panic(err.Error())
			}
		},
		Categories: []string{
			"all",
		},
	}
)

func init() {
	addToCategories(awsAuthConfigMap)
	addToCategories(awsNodePods)
}

func addToCategories(data *Data) {
	//TODO dedup the all and make sure no duplication if all is set as well
	for _, category := range data.Categories {
		GetCategories[category] = append(GetCategories[category], data)
	}
}
