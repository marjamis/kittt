package actions

import (
	"fmt"
	"log"

	"github.com/marjamis/kittt/internal/output"
	"github.com/marjamis/kittt/pkg/k8sConnector"

	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
)

type Data struct {
	Name       string
	Function   func(*kubernetes.Clientset) (interface{}, error)
	Categories []string
}

type ActionCategories map[string][]*Data

func (actionCategories ActionCategories) addToCategories(data *Data) {
	actionCategories["all"] = append(actionCategories["all"], data)
	for _, category := range data.Categories {
		if category != "all" {
			actionCategories[category] = append(actionCategories[category], data)
		}
	}
}

func actionExecute(item *Data, clientset *kubernetes.Clientset) {
	results, err := item.Function(clientset)
	if err != nil {
		panic(err.Error())
	}

	//TODO this is a bit too tightly coupled, same in a few other places, but need to have a think about this
	switch viper.GetString("output") {
	case "file":
		fmt.Println("File")
		// err = output.FormatStdout(results, item.Name)
		// if err != nil {
		// 	panic(err.Error())
		// }
	case "stdout":
		err = output.FormatStdout(results)
		if err != nil {
			panic(err.Error())
		}
	}
}

func (actionCategories ActionCategories) RunThroughAll() {
	clientset, err := k8sConnector.GenerateClientSet()
	if err != nil {
		panic(err)
	}

	for _, item := range actionCategories["all"] {
		fmt.Printf("Category: %s - Values: %s\n", "all", item.Name)
		actionExecute(item, clientset)
	}
}

func (actionCategories ActionCategories) RunThroughCategory(category string) {
	if _, ok := actionCategories[category]; !ok {
		var keys []string
		for k := range actionCategories {
			keys = append(keys, k)
		}
		log.Panicf("This isn't a valid category please choose all (for all objects) or select from one of these categories: %v", keys)
	}

	clientset, err := k8sConnector.GenerateClientSet()
	if err != nil {
		panic(err)
	}
	for _, item := range actionCategories[category] {
		fmt.Printf("Category: %s - Values: %s\n", category, item.Name)
		actionExecute(item, clientset)
	}
}
