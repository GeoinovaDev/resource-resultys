package resource

type item struct {
	id int
	cb func()
}

func (it item) GetID() int {
	return it.id
}
