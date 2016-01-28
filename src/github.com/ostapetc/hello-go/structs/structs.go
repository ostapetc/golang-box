package structs

type Person struct {
	name string
	age  int
	sex  string
}

func BuildStaticPerson() Person {
	return Person{"Artem", 30, "man"}
}
