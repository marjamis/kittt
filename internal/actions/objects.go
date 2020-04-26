package actions

import (
	"context"

	"github.com/marjamis/kittt/internal/output"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"

	// Used for auth
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func AwsAuthConfigMap(clientset *kubernetes.Clientset) {
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
}

func AwsNodePods(clientset *kubernetes.Clientset) {
	namespace := "kube-system"
	name := map[string]string{
		"k8s-app": "aws-node",
	}
	results, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels.Set(name).AsSelector().String(),
	})
	if err != nil {
		panic(err.Error())
	}

	output.FormatStdout(results)
	if err != nil {
		panic(err.Error())
	}
}
