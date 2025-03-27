// Package sarif is a non-exhuastive representation of the SARIF format.
// It just exposes the types for fields that are actively used by the Grafana Vale template.
package sarif

import (
	"encoding/json"
	"os"
)

type File struct {
	Runs    []Run  `json:"runs"`
	Version string `json:"version"`
	Schema  string `json:"$schema"`
}

type Run struct {
	Results []Result `json:"results"`
	Tool    Tool     `json:"tool"`
}

type Result struct {
	Level     string     `json:"level"`
	Message   Message    `json:"message"`
	Locations []Location `json:"locations"`
	RuleID    string     `json:"ruleId"`
}

type Message struct {
	Text string `json:"text"`
}

type Location struct {
	PhysicalLocation PhysicalLocation `json:"physicalLocation"`
}

type PhysicalLocation struct {
	ArtifactLocation ArtifactLocation `json:"artifactLocation"`
	Region           Region           `json:"region"`
}

type ArtifactLocation struct {
	URI string `json:"uri"`
}

type Region struct {
	StartLine   int `json:"startLine"`
	StartColumn int `json:"startColumn"`
}

type Tool struct {
	Driver Driver `json:"driver"`
}

type Driver struct {
	Name           string `json:"name"`
	InformationURI string `json:"informationUri"`
	Version        string `json:"version"`
}

func NewFromFile(path string) (File, error) {
	var file File

	data, err := os.ReadFile(path)
	if err != nil {
		return file, err
	}

	if err := json.Unmarshal(data, &file); err != nil {
		return file, err
	}

	return file, nil
}
