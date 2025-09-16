package action

import human "github.com/lastlife77/wb-tech/l1/1.1/Human"

type Action struct {
	human.Human
}

func New(name string, age int, sex bool) *Action {
	return &Action{
		Human: *human.New(name, age, sex),
	}
}
