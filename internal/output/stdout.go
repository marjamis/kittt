package output

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func FormatStdout(input interface{}) error {
	str, err := yaml.Marshal(input)
	if err != nil {
		return err
	}

	fmt.Printf("### %s\n%s\n###\n", "type", str)
	return nil
}

func FormatFiles(input interface{}) error {

	return nil
}
