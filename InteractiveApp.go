package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type PatientFeedback struct {
	DoctorRecScore int16
	ManageDiagosis string
	Feeling        string
}

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
	fmt.Println("Thank you. You were diagnosed with " +
		patientCaseDetails.Dignosis +
		". Did Dr " + patientCaseDetails.DoctorName.LastName +
		" explain how to manage this diagnosis in a way you could understand?")

	_, err = fmt.Scan(&patientFeedback.ManageDiagosis)
	if err != nil {
		log.Fatal("Error scanning user input", err)
	}
	if strings.ToLower(patientFeedback.ManageDiagosis) == "no" {
		fmt.Println("")
	}
	fmt.Println("We appreciate the feedback, one last question: how do you feel about being diagnosed with " +
		patientCaseDetails.Dignosis + "?")

	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)
	var feedback string
	// Read input line by line
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break // Exit loop if an empty line is entered
		}
		feedback += text
	}
	patientFeedback.Feeling = feedback
	fmt.Println("Thanks again! Here’s what we heard:\n" + feedback)

	return patientFeedback
}
