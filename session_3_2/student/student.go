package student

type Student struct {
	Name string
	id   int
}

func (s Student) Id() int {
	return s.id
}
