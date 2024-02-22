package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Patient App Main program
func main() {
	router := gin.Default()
	router.Use(corsMiddleware())
	router.GET("/patientCaseDetails", getPatientCaseDetails)
	router.POST("/patientFeedback", postPatientFeedback)
	router.Run("0.0.0.0:8080")
}
func getPatientCaseDetails(c *gin.Context) {
	// Read Json file with EMR data
	emrRecords := ReadEmrRecords()
	// Get Patient details like name, doctor's name and diagnosis
	patientInfo := GetPatientInfo(emrRecords)
	doctorInfo := GetDoctorInfo(emrRecords)
	diagnosisInfo := GetDiagnosisInfo(emrRecords)
	patientDetails := GetPatientCaseDetails(patientInfo, doctorInfo, diagnosisInfo)
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

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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
