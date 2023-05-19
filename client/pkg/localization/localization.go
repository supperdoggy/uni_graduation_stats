package localization

import "sync"

type Localization struct {
	data map[string]string
	mut  sync.Mutex
}

func NewLocalization() *Localization {
	return &Localization{
		data: map[string]string{
			"TopCompaniesBySchool_INPUT":    "Будь ласка введіть назву навчального закладу",
			"TopHiredDegrees_SCHOOL_INPUT":  "Будь ласка введіть назву навчального закладу",
			"TopHiredDegrees_COMPANY_INPUT": "Будь ласка введіть назву компанії",
			"TopSchoolsByCompany_INPUT":     "Будь ласка введіть назву компанії",
			"SchoolDegrees_INPUT":           "Будь ласка введіть назву навчального закладу",
		},
	}
}

func (l *Localization) Get(key string) string {
	l.mut.Lock()
	defer l.mut.Unlock()

	return l.data[key]
}
