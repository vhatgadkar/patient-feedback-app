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
func GetPatientCaseDetails(emrRecords *EmrRecords) PatientCaseDetails {
	var patientCaseDetails PatientCaseDetails
	for _, entry := range emrRecords.Entries {
		if entry.Resource.ResourceType == "Patient" {
			patientCaseDetails.PatientName.LastName = entry.Resource.Names[0].Family
			patientCaseDetails.PatientName.FirstName = entry.Resource.Names[0].Given[0]
		} else if entry.Resource.ResourceType == "Doctor" {
			patientCaseDetails.DoctorName.FirstName = entry.Resource.Names[0].Given[0]
			patientCaseDetails.DoctorName.LastName = entry.Resource.Names[0].Family
		} else if entry.Resource.ResourceType == "Diagnosis" {
			patientCaseDetails.Dignosis = entry.Resource.Code.CodeRecords[0].Name
		}

	}
	return patientCaseDetails
}
