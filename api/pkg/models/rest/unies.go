package rest

type ListUniversitiesSchools struct {
	Name  string `bson:"schoolName" json:"name"`
	Count int    `bson:"count"`
}

type ListUniversitiesResponse struct {
	Schools []ListUniversitiesSchools `json:"schools,omitempty"`
	Count   int                       `json:"count,omitempty""`
	Error   string                    `json:"error,omitempty"`
}
