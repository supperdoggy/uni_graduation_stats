package db

type Experience struct {
	Title            string `json:"title"`
	Company          string `json:"company"`
	StartDate        string `json:"startDate"`
	EndDate          string `json:"endDate"`
	EndDateIsPresent bool   `json:"endDateIsPresent"`
	Location         string `json:"location"`
	Description      string `json:"description"`
}

type Education struct {
	SchoolName   string `json:"schoolName"`
	DegreeName   string `json:"degreeName"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	FieldOfStudy string `json:"fieldOfStudy"`
}

type UserProfile struct {
	FullName    string `json:"fullName"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Person struct {
	ID          string       `json:"_id"`
	UserProfile UserProfile  `json:"userProfile"`
	Experiences []Experience `json:"experiences"`
	Education   []Education  `json:"education"`
	Skills      []string     `json:"skills"`
}
