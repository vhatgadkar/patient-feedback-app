package main

type Diagnosis struct {
	id          string
	meta        Meta
	status      string
	code        Coding
	appointment Ref
}

// Extract Diagnosis info from emr record
func GetDiagnosisInfo(emrRecords *EmrRecords) Diagnosis {
	var diagnosis Diagnosis
	for _, entry := range emrRecords.Entries {
		if entry.Resource.ResourceType == "Diagnosis" {
			diagnosis.id = entry.Resource.Id
			diagnosis.meta = entry.Resource.Meta
			diagnosis.status = entry.Resource.Status
			diagnosis.code = entry.Resource.Code
			diagnosis.appointment = entry.Resource.Appointment
		}

	}
	return diagnosis
}
