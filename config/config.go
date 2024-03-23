package config

import (
	"errors"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func ProcessWorkflowYAML(filePath string) (*Workflow, error) {
	var workflow Workflow
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("yamlFile.Get err " + err.Error())
	}
	yamlFile, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("yamlFile.Get err " + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &workflow)
	if err != nil {
		return nil, errors.New("Unmarshal: " + err.Error())
	}
	return &workflow, nil
}
