# patient-feedback-app
This is a toy program to collect patient feedback after an appointment

Prompt

Please build a toy program to collect patient feedback after an appointment. Attached is some JSON with sample data about patients and their appointments which should drive the program. And below is a sample script with what questions each patient should be asked.

We know we’ve given you a relatively small amount of time, and we want this to be fun. You just need to build something clean that functions and meets the requirements noted below.

Sample Script

Program asks: “Hi [Patient First Name], on a scale of 1-10, would you recommend Dr [Doctor Last Name] to a friend or family member? 1 = Would not recommend, 10 = Would strongly recommend”

Assume patient responds with an integer between 1-10

Program asks: “Thank you. You were diagnosed with [Diagnosis]. Did Dr [Doctor Last Name] explain how to manage this diagnosis in a way you could understand?”

Record whatever the patient responds with, consider special treatment of Yes/No

Program asks: “We appreciate the feedback, one last question: how do you feel about being diagnosed with [Diagnosis]?”

Record whatever the patient responds with

Program ends with: “Thanks again! Here’s what we heard:”

Take some creative license with how patients are shown what the program recorded
