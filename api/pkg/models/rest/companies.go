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

type ListDegrees struct {
	Name  string `bson:"degreeName" json:"name"`
	Count int    `bson:"count" json:"count"`
}

type ListCompaniesTopSchools struct {
	Name    string        `bson:"_id" json:"name"`
	Count   int           `bson:"count" json:"count"`
	Degrees []ListDegrees `bson:"degrees" json:"degrees"`
}

type ListCompaniesTopSchoolsResponse struct {
	Schools []ListCompaniesTopSchools `json:"schools,omitempty"`
	Count   int                       `json:"count,omitempty"`
	Error   string                    `json:"error,omitempty"`
}

type TopHiredDegreesRequest struct {
	Company string `json:"company"`
	School  string `json:"school"`
}

type TopHiredDegrees struct {
	Name      string `bson:"degreeName" json:"name"`
	StartDate string `bson:"startDate" json:"startDate"`
	EndDate   string `bson:"endDate" json:"endDate"`
	Count     int    `bson:"count" json:"count"`
}

type TopHiredDegreesResponse struct {
	Degrees        []TopHiredDegrees `json:"degrees,omitempty"`
	Count          int               `json:"count,omitempty"`
	TotalEmployees int               `json:"totalEmployees,omitempty"`
	Error          string            `json:"error,omitempty"`
}
