package main

type Doctor struct {
	id        string
	firstName string
	lastName  string
}

// Extract Doctor info from emr record
func GetDoctorInfo(emrRecords *EmrRecords) Doctor {
	var doctor Doctor
	for _, entry := range emrRecords.Entries {
		if entry.Resource.ResourceType == "Doctor" {
			doctor.id = entry.Resource.Id
			doctor.firstName = entry.Resource.Names[0].Given[0]
			doctor.lastName = entry.Resource.Names[0].Family
		}

	}
	return doctor
}
