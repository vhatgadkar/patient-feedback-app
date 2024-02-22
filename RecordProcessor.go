package main

type NameInfo struct {
	FirstName string
	LastName  string
}

type PatientCaseDetails struct {
	PatientName NameInfo
	DoctorName  NameInfo
	Dignosis    string
}

// Extract Patient details from emr record
func GetPatientCaseDetails(patientInfo Patient, doctorInfo Doctor, diagnosisInfo Diagnosis) PatientCaseDetails {
	var patientCaseDetails PatientCaseDetails

	patientCaseDetails.PatientName.LastName = patientInfo.lastName
	patientCaseDetails.PatientName.FirstName = patientInfo.firstName

	patientCaseDetails.DoctorName.FirstName = doctorInfo.firstName
	patientCaseDetails.DoctorName.LastName = doctorInfo.lastName

	patientCaseDetails.Dignosis = diagnosisInfo.code.CodeRecords[0].Name

	return patientCaseDetails
}
