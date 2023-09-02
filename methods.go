package main

import "fmt"

type Creature struct {
	Name     string
	Greeting string
}

/*    func (c Creature) Greet() { // This is a method just like define a function a method on a struct type.
	fmt.Printf("%s says %s", c.Name, c.Greeting)
}
func main() {
	sammy := Creature{
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	//Creature.Greet(sammy)
	sammy.Greet() // calling method by sammy .
}    */

func (c Creature) Greet() Creature { // call method by dot
	fmt.Printf("%s says %s!\n", c.Name, c.Greeting)
	return c
}
func (c Creature) SayGoodbye(name string) { // call method by struct
	fmt.Println("Farewell", name, "!")
}
func main() {
	sammy := Creature{
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	sammy.Greet().SayGoodbye("gophers")
	Creature.SayGoodbye(Creature.Greet(sammy), "gophers")
}
