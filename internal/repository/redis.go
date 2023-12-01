package psql

type Redis interface {
	GetAllSubjectTypes()
	GetAllSubjectNames()
	GetAllSemesters()
	GetAllInstitutes()

	GetUniqueSubjectTypes()
	GetUniqueSubjectNames()
	GetUniqueSemesters()
	GetUniqueInstitutes()
}
