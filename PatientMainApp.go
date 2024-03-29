package main

import (
	"encoding/json"
	"log"
	"os"
)

// Patient App Main program
func main() {
	// Read Json file with EMR data
	emrRecords := ReadEmrRecords()
	// Get Patient details like name, doctor's name and diagnosis
	patientInfo := GetPatientInfo(emrRecords)
	doctorInfo := GetDoctorInfo(emrRecords)
	diagnosisInfo := GetDiagnosisInfo(emrRecords)
	patientDetails := GetPatientCaseDetails(patientInfo, doctorInfo, diagnosisInfo)
	// Get patient feedback
	patientFeedback := PatientInteraction(patientDetails)

	// Save patient's feedback to a Json file
	savePatientFeedback(patientFeedback)

}

// Save patient's feedback to a Json file well indented
func savePatientFeedback(patientFeedback PatientFeedback) {
	file, err := json.MarshalIndent(patientFeedback, "", " ")
	if err != nil {
		log.Fatal("Couldn't marshall feedback to file:", err)
	}

	err = os.WriteFile("saved-patient-feedback.json", file, 0644)
	if err != nil {
		log.Fatal("Couldn't save feedback to file:", err)
	}
}
