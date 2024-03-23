package config

import (
	"errors"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func ProcessWorkflowYAMLFile(filePath string) (*Workflow, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("yamlFile.Get err " + err.Error())
	}
	yamlFile, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("yamlFile.Get err " + err.Error())
	}
	workflow, err := ProcessWorkflowYAML(string(yamlFile))
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return workflow, nil
}

func ProcessWorkflowYAML(yamlData string) (*Workflow, error) {
	var workflow Workflow
	err := yaml.Unmarshal([]byte(yamlData), &workflow)
	if err != nil {
		return nil, errors.New("Unmarshal: " + err.Error())
	}
	return &workflow, nil
}

func ProcessActionYAMLFile(filePath string) (*Action, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("yamlFile.Get err " + err.Error())
	}
	yamlFile, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("yamlFile.Get err " + err.Error())
	}
	workflow, err := ProcessActionYAML(string(yamlFile))
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return workflow, nil
}

func ProcessActionYAML(yamlData string) (*Action, error) {
	var action Action
	err := yaml.Unmarshal([]byte(yamlData), &action)
	if err != nil {
		return nil, errors.New("Unmarshal: " + err.Error())
	}
	return &action, nil
}
