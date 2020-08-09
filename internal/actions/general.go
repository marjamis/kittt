package actions

import (
	"fmt"
	"log"

	"github.com/marjamis/kittt/internal/utils"
	"github.com/marjamis/kittt/pkg/kube"

	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
)

type Data struct {
	Name string
	// TODO fix the below to be more generic?
	Function   func(*kubernetes.Clientset) (interface{}, error)
	Categories []string
	Platforms  []string
}

//TODO rename
type ActionTypes map[string][]*Data

var (
	Clientset *kubernetes.Clientset
)

func init() {
	cs, err := kube.GenerateClientSet()
	if err != nil {
		panic(err)
	}

	Clientset = cs
}

func (ActionTypes ActionTypes) addToCategories(data *Data) {
	ActionTypes["all"] = append(ActionTypes["all"], data)
	for _, category := range data.Categories {
		if category != "all" {
			ActionTypes[category] = append(ActionTypes[category], data)
		}
	}
}

//TODO work out naming scheme for variables and functions
func actionExecute(item *Data) {
	//TODO this should be done once somewhere and then used
	results, err := item.Function(Clientset)
	if err != nil {
		panic(err.Error())
	}

	//TODO is this is a bit too tightly coupled, same in a few other places, but need to have a think about this
	switch viper.GetString("output") {
	case "file":
		err = utils.FormatFile(results, item.Name)
		if err != nil {
			panic(err.Error())
		}
	case "stdout":
		err = utils.FormatStdout(results, item.Name)
		if err != nil {
			panic(err.Error())
		}
	}
}

func (ActionTypes ActionTypes) RunThroughAll() {
	ActionTypes.RunThroughCategory("all")
}

func (ActionTypes ActionTypes) RunThroughCategory(category string) {
	if _, ok := ActionTypes[category]; !ok {
		var keys []string
		for k := range ActionTypes {
			keys = append(keys, k)
		}
		log.Panicf("This isn't a valid category please choose all (for all objects) or select from one of these categories: %v", keys)
	}

	for _, item := range ActionTypes[category] {
		fmt.Printf("Category: %s - Values: %s\n", category, item.Name)
		actionExecute(item)
	}
}
