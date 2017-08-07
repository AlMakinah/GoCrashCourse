package main

import "fmt"

// START OMIT
type Human interface {
	Talk()
}

type Person struct {
	Name       string
	Age        int
	Occupation string
}

func (p *Person) Talk() {
	fmt.Printf("Hi my name is %v\n", p.Name)
}

func makeHumanTalk(h Human) {
	h.Talk()
}

func main() {
	p := &Person{"Omar", 29, "Developer"}
	makeHumanTalk(p)
}

// END OMIT
