package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func FormatStdout(input interface{}, name string) error {
	str, err := yaml.Marshal(input)
	if err != nil {
		return err
	}

	fmt.Printf("### %s\n%s\n###\n", name, str)
	return nil
}

func FormatFile(input interface{}, filename string) error {
	str, err := yaml.Marshal(input)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./"+filename, []byte(str), 0644)
	if err != nil {
		return err
	}

	return nil
}
