package entry

type ByPath []*Entry

func (e ByPath) Len() int {
	return len(e)
}

func (e ByPath) Less(i, j int) bool {
	return e[i].Path < e[j].Path
}

func (e ByPath) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type ByFrecency []*Entry

func (e ByFrecency) Len() int {
	return len(e)
}

func (e ByFrecency) Less(i, j int) bool {
	return Frecency(e[i]) < Frecency(e[j])
}

func (e ByFrecency) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
