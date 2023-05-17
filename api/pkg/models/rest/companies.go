package rest

type ListCompanies struct {
	Name  string `bson:"companyName" json:"name"`
	Count int    `bson:"count" json:"count"`
}

type ListCompaniesResponse struct {
	Companies []ListCompanies `json:"companies,omitempty"`
	Count     int             `json:"count,omitempty"`
	Error     string          `json:"error,omitempty"`
}

type ListCompaniesTopSchoolsRequest struct {
	Company string `json:"company"`
}

type ListCompaniesTopSchools struct {
	Name  string `bson:"_id" json:"name"`
	Count int    `bson:"count" json:"count"`
}

type ListCompaniesTopSchoolsResponse struct {
	Schools []ListCompaniesTopSchools `json:"schools,omitempty"`
	Count   int                       `json:"count,omitempty"`
	Error   string                    `json:"error,omitempty"`
}
