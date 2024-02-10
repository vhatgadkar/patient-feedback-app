package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
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
	jsonStr := string(jsonBytes)
	testStr := "{\"resourceType\":\"Bundle\",\"id\":\"0c3151bd-1cbf-4d64-b04d-cd9187a4c6e0\",\"timestamp\":\"2021-04-02T12:12:21Z\",\"entry\":[{\"resource\":{\"resourceType\":\"Patient\",\"id\":\"6739ec3e-93bd-11eb-a8b3-0242ac130003\",\"active\":true,\"name\":[{\"text\":\"Tendo Tenderson\",\"family\":\"Tenderson\",\"given\":[\"Tendo\"]}],\"contact\":[{\"system\":\"phone\",\"value\":\"555-555-2021\",\"use\":\"mobile\"},{\"system\":\"email\",\"value\":\"tendo@tendoco.com\",\"use\":\"work\"}],\"gender\":\"female\",\"birthDate\":\"1955-01-06\",\"address\":[{\"use\":\"home\",\"line\":[\"2222 Home Street\"]}],\"type\":null,\"status\":\"\",\"subject\":{\"reference\":\"\"},\"actor\":{\"reference\":\"\"},\"period\":{\"start\":\"\",\"end\":\"\"},\"appointment\":{\"reference\":\"\"},\"code\":{\"coding\":null},\"meta\":{\"lastUpdated\":\"\"}}}]}"
	assert.Equal(t, testStr, jsonStr)
}
