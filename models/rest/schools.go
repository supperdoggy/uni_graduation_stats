package rest

type ListSchools struct {
	Name  string `bson:"schoolName" json:"name"`
	Count int    `bson:"count"`
}

type ListSchoolsResponse struct {
	Schools []ListSchools `json:"schools,omitempty"`
	Count   int           `json:"count,omitempty""`
	Error   string        `json:"error,omitempty"`
}

type ListSchoolsTopCompaniesRequest struct {
	School string `json:"school"`
}

type ListSchoolsTopCompanies struct {
	Name  string `bson:"_id" json:"name"`
	Count int    `bson:"count" json:"count"`
}

type ListSchoolsTopCompaniesResponse struct {
	Companies []ListSchoolsTopCompanies `json:"companies,omitempty"`
	Count     int                       `json:"count,omitempty"`
	Error     string                    `json:"error,omitempty"`
}

type ListJobsBySchoolRequest struct {
	School string `json:"school"`
}

type ListJobsBySchool struct {
	Title string `bson:"title" json:"title"`
	Count int    `bson:"count" json:"count"`
}

type ListJobsBySchoolResponse struct {
	Jobs  []ListJobsBySchool `json:"jobs,omitempty"`
	Count int                `json:"count,omitempty"`
	Error string             `json:"error,omitempty"`
}

type CorrelationDegreeAndTitleRequest struct {
	School string `json:"school"`
}

type CorreelationExperience struct {
	Company   string `bson:"company" json:"company"`
	Title     string `bson:"title" json:"title"`
	StartDate string `bson:"startDate" json:"startDate"`
	EndDate   string `bson:"endDate" json:"endDate"`
}

type CorrelationDegreeAndTitle struct {
	Degree     string                   `bson:"degreeName" json:"degree"`
	Experience []CorreelationExperience `bson:"experience" json:"experience"`
}

type CorrelationDegreeAndTitleResponse struct {
	Correlations []CorrelationDegreeAndTitle `json:"correlations,omitempty"`
	Count        int                         `json:"count,omitempty"`
	Error        string                      `json:"error,omitempty"`
}

type SchoolDegreesRequest struct {
	School string `json:"school"`
}

type SchoolDegrees struct {
	Degree     string `bson:"degreeName" json:"degree"`
	SchoolName string `bson:"schoolName" json:"schoolName"`
	Count      int    `bson:"count" json:"count"`
}

type SchoolDegreesResponse struct {
	Degrees []SchoolDegrees `json:"degrees,omitempty"`
	Count   int             `json:"count,omitempty"`
	Error   string          `json:"error,omitempty"`
}
