package entry

type Entry struct {
	Path         string
	VisitedCount int
	LastVisited  int
}

type Entries []*Entry

func (e Entries) Map(f func(*Entry) interface{}) Entries {
	result := make(Entries, len(e))
	for _, v := range e {
		result = append(result, v)
	}
	return result
}

func (e Entries) Filter(f func(*Entry) bool) Entries {
	result := make(Entries, 0)
	for _, v := range e {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
