package main

type Patient struct {
	id        string
	active    bool
	firstName string
	lastName  string
	contacts  []Contact
	gender    string
	birthDate string
	address   []Location
}

// Extract Patient info from emr record
func GetPatientInfo(emrRecords *EmrRecords) Patient {
	var patient Patient
	for _, entry := range emrRecords.Entries {
		if entry.Resource.ResourceType == "Patient" {
			patient.id = entry.Resource.Id
			patient.active = entry.Resource.Active
			patient.lastName = entry.Resource.Names[0].Family
			patient.firstName = entry.Resource.Names[0].Given[0]
			patient.contacts = entry.Resource.Contacts
			patient.gender = entry.Resource.Gender
			patient.birthDate = entry.Resource.BirthDate
			patient.address = entry.Resource.Address
		}

	}
	return patient
}
