package main

type Appointment struct {
	id      string
	status  string
	types   []Type
	subject Ref
	actor   Ref
	period  Period
}

// Extract appointment info from emr record
func GetAppointmentInfo(emrRecords *EmrRecords) Appointment {
	var appointment Appointment
	for _, entry := range emrRecords.Entries {
		if entry.Resource.ResourceType == "Appointment" {
			appointment.id = entry.Resource.Id
			appointment.status = entry.Resource.Status
			appointment.types = entry.Resource.Types
			appointment.subject = entry.Resource.Subject
			appointment.actor = entry.Resource.Actor
			appointment.period = entry.Resource.Period
		}

	}
	return appointment
}
