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

func ProcessGitterYAMLFile(filePath string) (*Gitter, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("yamlFile.Get err " + err.Error())
	}
	yamlFile, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("yamlFile.Get err " + err.Error())
	}
	workflow, err := ProcessGitterYAML(string(yamlFile))
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return workflow, nil
}

func ProcessGitterYAML(yamlData string) (*Gitter, error) {
	var gitter Gitter
	err := yaml.Unmarshal([]byte(yamlData), &gitter)
	if err != nil {
		return nil, errors.New("Unmarshal: " + err.Error())
	}
	return &gitter, nil
}

func ProcessDBConfigYAMLFile(filePath string) (*DBConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("yamlFile.Get err " + err.Error())
	}
	yamlFile, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("yamlFile.Get err " + err.Error())
	}
	workflow, err := ProcessDBConfigYAML(string(yamlFile))
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return workflow, nil
}

func ProcessDBConfigYAML(yamlData string) (*DBConfig, error) {
	var dbconfig DBConfig
	err := yaml.Unmarshal([]byte(yamlData), &dbconfig)
	if err != nil {
		return nil, errors.New("Unmarshal: " + err.Error())
	}
	return &dbconfig, nil
}
