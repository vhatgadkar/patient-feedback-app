package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Test Emr JSON Patient record with struct defined in PatientInputReader file
func TestEmrStructPatientRecord(t *testing.T) {
	name := []Name{{
		Text:   "Tendo Tenderson",
		Family: "Tenderson",
		Given:  []string{"Tendo"},
	}}
	contact := []Contact{
		{
			System: "phone",
			Value:  "555-555-2021",
			Use:    "mobile",
		},
		{
			System: "email",
			Value:  "tendo@tendoco.com",
			Use:    "work",
		},
	}
	address := []Location{
		{
			Use: "home",
			Line: []string{
				"2222 Home Street",
			},
		},
	}
	resource := Resource{
		ResourceType: "Patient",
		Id:           "6739ec3e-93bd-11eb-a8b3-0242ac130003",
		Active:       true,
		Names:        name,
		Contacts:     contact,
		Gender:       "female",
		BirthDate:    "1955-01-06",
		Address:      address,
	}
	entries := []Entry{
		{
			Resource: resource,
		},
	}
	emrRecords := EmrRecords{
		ResType:   "Bundle",
		Id:        "0c3151bd-1cbf-4d64-b04d-cd9187a4c6e0",
		Timestamp: "2021-04-02T12:12:21Z",
		Entries:   entries,
	}

	jsonBytes, err := json.Marshal(emrRecords)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}
