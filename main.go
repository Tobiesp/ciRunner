package main

import (
	"fmt"

	"tspdev.com/m/v2/config"
)

func main() {
	obj, err := config.ProcessWorkflowYAML("example.yaml")
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
	}

	fmt.Println(obj)

}
