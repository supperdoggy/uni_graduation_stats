package states

import "sync"

const (
	NONE = iota
	TopCompaniesBySchool_INPUT
	TopHiredDegrees_SCHOOL_INPUT
	TopHiredDegrees_COMPANY_INPUT
	TopSchoolsByCompany_INPUT
	SchoolDegrees_INPUT
)

type state struct {
	state  int
	values []string
}

type StateManager struct {
	data map[int64]state
	mut  sync.Mutex
}

func NewStateManager() *StateManager {
	return &StateManager{
		data: make(map[int64]state),
	}
}

func (s *StateManager) SetState(id int64, st int, values []string) {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.data[id] = state{
		state:  st,
		values: values,
	}
}

func (s *StateManager) SetStateValues(id int64, values []string) {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.data[id] = state{
		state:  s.data[id].state,
		values: values,
	}
}

func (s *StateManager) GetState(id int64) (int, []string) {
	s.mut.Lock()
	defer s.mut.Unlock()

	return s.data[id].state, s.data[id].values
}

func (s *StateManager) DeleteState(id int64) {
	s.mut.Lock()
	defer s.mut.Unlock()

	delete(s.data, id)
}
