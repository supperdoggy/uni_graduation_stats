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
	Companies []ListSchoolsTopCompanies `json:"school,omitempty"`
	Count     int                       `json:"count,omitempty"`
	Error     string                    `json:"error,omitempty"`
}
