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

type ListCompaniesTopUniversitiesRequest struct {
	Company string `json:"company"`
}

type ListCompaniesTopUniversities struct {
	Name  string `bson:"_id" json:"name"`
	Count int    `bson:"count" json:"count"`
}

type ListCompaniesTopUniversitiesResponse struct {
	Universities []ListCompaniesTopUniversities `json:"universities,omitempty"`
	Count        int                            `json:"count,omitempty"`
	Error        string                         `json:"error,omitempty"`
}
