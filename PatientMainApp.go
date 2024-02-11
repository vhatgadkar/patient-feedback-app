package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Patient App Main program
func main() {
	router := gin.Default()
	router.GET("/patientCaseDetails", getPatientCaseDetails)
	router.POST("/patientFeedback", postPatientFeedback)
	router.Run("localhost:8080")
}
func getPatientCaseDetails(c *gin.Context) {
	// Read Json file with EMR data
	emrRecords := ReadEmrRecords()
	// Get Patient details like name, doctor's name and diagnosis
	patientDetails := GetPatientCaseDetails(emrRecords)
	c.IndentedJSON(http.StatusOK, patientDetails)
}
func postPatientFeedback(c *gin.Context) {
	var patientFeedback PatientFeedback
	// Get patient feedback
	err := c.BindJSON(&patientFeedback)
	if err != nil {
		fmt.Println("Couldn't get feedback from client")
	}
	// Save patient's feedback to a Json file well indented
	file, err := json.MarshalIndent(patientFeedback, "", " ")
	if err != nil {
		fmt.Println("Couldn't save feedback to file")
	}

	_ = os.WriteFile("saved-patient-feedback.json", file, 0644)
}
