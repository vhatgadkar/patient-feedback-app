package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Patient App Main program
func main() {
	// Read Json file with EMR data
	emrRecords := ReadEmrRecords()
	// Get Patient details like name, doctor's name and diagnosis
	patientDetails := GetPatientCaseDetails(emrRecords)
	// Get patient feedback
	patientFeedback := PatientInteraction(patientDetails)

	// Save patient's feedback to a Json file well indented
	file, err := json.MarshalIndent(patientFeedback, "", " ")
	if err != nil {
		fmt.Println("Couldn't save feedback to file")
	}

	_ = os.WriteFile("saved-patient-feedback.json", file, 0644)

}
