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
func GetPatientCaseDetails(emrRecords EmrRecords) PatientCaseDetails {
	var patientCaseDetails PatientCaseDetails
	for i := 0; i < len(emrRecords.Entries); i++ {
		if emrRecords.Entries[i].Resource.ResourceType == "Patient" {
			patientCaseDetails.PatientName.LastName = emrRecords.Entries[i].Resource.Names[0].Family
			patientCaseDetails.PatientName.FirstName = emrRecords.Entries[i].Resource.Names[0].Given[0]
		} else if emrRecords.Entries[i].Resource.ResourceType == "Doctor" {
			patientCaseDetails.DoctorName.FirstName = emrRecords.Entries[i].Resource.Names[0].Given[0]
			patientCaseDetails.DoctorName.LastName = emrRecords.Entries[i].Resource.Names[0].Family
		} else if emrRecords.Entries[i].Resource.ResourceType == "Diagnosis" {
			patientCaseDetails.Dignosis = emrRecords.Entries[i].Resource.Code.CodeRecords[0].Name
		}

	}
	return patientCaseDetails
}
