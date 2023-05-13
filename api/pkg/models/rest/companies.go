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
