package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PatientFeedback struct {
	DoctorRecScore        int16
	ManageDiagosis        string
	ManageDiagosisComment string
	Feeling               string
}

// Get patient feedback through interactive session
func PatientInteraction(patientCaseDetails PatientCaseDetails) PatientFeedback {
	var patientFeedback PatientFeedback
	// Get recommend score for Doctor
	getRecommendScore(&patientFeedback, &patientCaseDetails)
	// Get explaination of manage diagnosis feedback
	getManageDiagnosisFeedback(&patientFeedback, &patientCaseDetails)
	// Get feeling about diagnosis
	getDiagnosisFeeling(&patientFeedback, &patientCaseDetails)
	// Write summary
	fmt.Println()
	writeSummary(&patientFeedback, &patientCaseDetails)

	return patientFeedback
}

// Get recommend score for Doctor, sending pointer because patientFeedback is modified
func getRecommendScore(patientFeedback *PatientFeedback, patientCaseDetails *PatientCaseDetails) {
	fmt.Println("Hi " + patientCaseDetails.PatientName.FirstName +
		", on a scale of 1-10, would you recommend Dr " +
		patientCaseDetails.DoctorName.LastName +
		" to a friend or family member? 1 = Would not recommend, 10 = Would strongly recommend")

	// Read a single line of input
	_, err := fmt.Scan(&patientFeedback.DoctorRecScore)
	if err != nil {
		log.Fatal("Error scanning user input", err)
	}
}

// Get explaination of manage diagnosis feedback
func getManageDiagnosisFeedback(patientFeedback *PatientFeedback, patientCaseDetails *PatientCaseDetails) {
	fmt.Println("\nThank you. You were diagnosed with " + patientCaseDetails.Dignosis +
		". Did Dr " + patientCaseDetails.DoctorName.LastName +
		" explain how to manage this diagnosis in a way you could understand?")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	patientFeedback.ManageDiagosis = scanner.Text()
	if strings.ToLower(patientFeedback.ManageDiagosis) == "no" {
		fmt.Println("Will you like to add a comment?")
		scanner.Scan()
		patientFeedback.ManageDiagosisComment = scanner.Text()
	}
}

// Get feeling about diagnosis
func getDiagnosisFeeling(patientFeedback *PatientFeedback, patientCaseDetails *PatientCaseDetails) {
	fmt.Println("\nWe appreciate the feedback, one last question: how do you feel about being diagnosed with " +
		patientCaseDetails.Dignosis + "?")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	patientFeedback.Feeling = scanner.Text()
}

// Write summary of patient feedback
func writeSummary(patientFeedback *PatientFeedback, patientCaseDetails *PatientCaseDetails) {
	fmt.Println("\nThanks again! Here’s what we heard:\n----------------------------------")

	fmt.Println("\nOn a scale of 1-10, would you recommend Dr " + patientCaseDetails.DoctorName.LastName +
		" to a friend or family member? \nYour answer: " + strconv.Itoa(int(patientFeedback.DoctorRecScore)))

	fmt.Println("\nYou were diagnosed with " + patientCaseDetails.Dignosis +
		". Did Dr " + patientCaseDetails.DoctorName.LastName +
		" explain how to manage this diagnosis in a way you could understand?\nYour answer: " +
		patientFeedback.ManageDiagosis)

	if patientFeedback.ManageDiagosisComment != "" {
		fmt.Println("Your comment: " + patientFeedback.ManageDiagosisComment)
	}

	fmt.Println("\nHow do you feel about being diagnosed with " + patientCaseDetails.Dignosis + "?\nYour answer: " + patientFeedback.Feeling + "\n")
}
