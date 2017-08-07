package main

import "fmt"

type Person struct {
	Name       string
	Age        int
	Occupation string
}

func (p *Person) Talk() {
	fmt.Printf("Hi my name is %v\n", p.Name)
}

func main() {
	p := &Person{"Omar", 29, "Developer"}
	p.Talk()
	fmt.Printf("%v is %v and works as a %v\n", p.Name, p.Age, p.Occupation)
}
