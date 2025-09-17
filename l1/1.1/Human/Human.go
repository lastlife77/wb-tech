package human

type Human struct {
	Name string
	Age  int
	//True is male
	//False is female
	Sex bool
}

func New(name string, age int, sex bool) *Human {
	return &Human{
		Name: name,
		Age:  age,
		Sex:  sex,
	}
}

func (h *Human) GetOlder() {
	h.Age++
}

func (h *Human) IsMale() bool {
	return h.Sex
}

func (h *Human) IsFemale() bool {
	return !h.Sex
}
