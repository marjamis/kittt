package actions

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"

	// Used for auth
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var (
	// Creates the array of Get actions broken down by category
	GetCategories = ActionTypes{}

	//
	awsAuthConfigMap = &Data{
		Name: "aws-auth",
		Function: func(clientset *kubernetes.Clientset) (interface{}, error) {
			namespace := "kube-system"
			name := "aws-auth"
			results, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
			return results, err
		},
		Categories: []string{
			"auth",
		},
		Platforms: []string{
			"eks",
		},
	}

	awsNodePods = &Data{
		Name: "aws-node",
		Function: func(clientset *kubernetes.Clientset) (interface{}, error) {
			namespace := "kube-system"
			selector := map[string]string{
				"k8s-app": "aws-node",
			}
			results, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
				LabelSelector: labels.Set(selector).String(),
			})

			return results, err
		},
		Categories: []string{
			"cni",
		},
		Platforms: []string{
			"eks",
		},
	}
)

func init() {
	GetCategories.addToCategories(awsAuthConfigMap)
	GetCategories.addToCategories(awsNodePods)
}
