package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// Structs for saving patient-feedback-raw-data.json data
type EmrRecords struct {
	ResType   string  `json:"resourceType"`
	Id        string  `json:"id"`
	Timestamp string  `json:"timestamp"`
	Entries   []Entry `json:"entry"`
}
type Entry struct {
	Resource Resource `json:"resource"`
}

type Resource struct {
	ResourceType string     `json:"resourceType"`
	Id           string     `json:"id"`
	Active       bool       `json:"active"`
	Names        []Name     `json:"name"`
	Contacts     []Contact  `json:"contact"`
	Gender       string     `json:"gender"`
	BirthDate    string     `json:"birthDate"`
	Address      []Location `json:"address"`
	Types        []Type     `json:"type"`
	Status       string     `json:"status"`
	Subject      Ref        `json:"subject"`
	Actor        Ref        `json:"actor"`
	Period       Period     `json:"period"`
	Appointment  Ref        `json:"appointment"`
	Code         Coding     `json:"code"`
	Meta         Meta       `json:"meta"`
}
type Name struct {
	Text   string   `json:"text"`
	Family string   `json:"family"`
	Given  []string `json:"given"`
}
type Contact struct {
	System string `json:"system"`
	Value  string `json:"value"`
	Use    string `json:"use"`
}
type Location struct {
	Use  string   `json:"use"`
	Line []string `json:"line"`
}
type Type struct {
	Text string `json:"text"`
}

type Ref struct {
	Reference string `json:"reference"`
}
type Period struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
type Coding struct {
	CodeRecords []CodeRec `json:"coding"`
}
type CodeRec struct {
	System string `json:"system"`
	Code   string `json:"code"`
	Name   string `json:"name"`
}
type Meta struct {
	LastUpdated string `json:"lastUpdated"`
}

// Read the sample JSON file and return EmrRecords
func ReadEmrRecords() EmrRecords {
	// Open sample Json file
	jsonFile, err := os.Open("patient-feedback-raw-data.json")
	if err != nil {
		log.Fatal("Error reading Json file", err)
	}
	// defer the closing of Json file, after parsing it will be closed
	defer jsonFile.Close()

	// Initialize the EmrRecords variable
	var emrRecords EmrRecords

	// Read Json file, unmarshall the Json
	byteValue, _ := io.ReadAll(jsonFile)

	// Unmarshall the byteArray that contains the Jsonfile's data
	err = json.Unmarshal(byteValue, &emrRecords)
	if err != nil {
		log.Fatal("Error unmarshaling json: ", err)
	}

	return emrRecords
}
