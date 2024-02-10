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
	DoctorRecScore int16
	ManageDiagosis string
	Feeling        string
}

// Get patient feedback through interactive session
func PatientInteraction(patientCaseDetails PatientCaseDetails) PatientFeedback {
	var patientFeedback PatientFeedback
	fmt.Println("Hi " + patientCaseDetails.PatientName.FirstName +
		", on a scale of 1-10, would you recommend Dr " +
		patientCaseDetails.DoctorName.LastName +
		" to a friend or family member? 1 = Would not recommend, 10 = Would strongly recommend")

	// Read a single line of input
	_, err := fmt.Scan(&patientFeedback.DoctorRecScore)
	if err != nil {
		log.Fatal("Error scanning user input", err)
	}
	fmt.Println("Thank you. You were diagnosed with " + patientCaseDetails.Dignosis +
		". Did Dr " + patientCaseDetails.DoctorName.LastName +
		" explain how to manage this diagnosis in a way you could understand?")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	patientFeedback.ManageDiagosis = scanner.Text()
	if strings.ToLower(patientFeedback.ManageDiagosis) == "no" {
		fmt.Println("Oh, sorry for that. We will inform Dr " + patientCaseDetails.DoctorName.LastName + " to contact you.\n")
	}
	fmt.Println("We appreciate the feedback, one last question: how do you feel about being diagnosed with " +
		patientCaseDetails.Dignosis + "?")

	patientFeedback.Feeling = scanMultilineInput()

	fmt.Println("Thanks again! Here’s what we heard:")
	fmt.Println("On a scale of 1-10, would you recommend Dr " + patientCaseDetails.DoctorName.LastName +
		" to a friend or family member? \nYour answer: " + strconv.Itoa(int(patientFeedback.DoctorRecScore)))
	fmt.Println("You were diagnosed with " + patientCaseDetails.Dignosis +
		". Did Dr " + patientCaseDetails.DoctorName.LastName +
		" explain how to manage this diagnosis in a way you could understand?\nYour answer: " + patientFeedback.ManageDiagosis)
	fmt.Println("How do you feel about being diagnosed with " + patientCaseDetails.Dignosis + "?\nYour answer: " + patientFeedback.Feeling)

	return patientFeedback
}

// Get multiple lines of input from patient
func scanMultilineInput() string {
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)
	var feedback string
	// Read input line by line
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break // Exit loop if an empty line is entered
		}
		feedback += "\n" + text
	}
	return feedback
}
