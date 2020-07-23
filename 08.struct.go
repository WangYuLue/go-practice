package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 引用传递
	sp := &s
	sp.age = 51
	fmt.Println(s.age)
	fmt.Println(sp.age)

	// 值传递
	sv := s
	sv.age = 100

	fmt.Println(s.age)
	fmt.Println(sv.age)
}
