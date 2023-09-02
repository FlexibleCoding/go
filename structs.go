/*
keyword type followed by the name of the struct. After this, use the
struct keyword followed by a pair of braces {} where you declare the
fields the struct will contain. Once you have defined the struct, you are
then able to declare variables that use this struct definition. This example
defines a struct and uses it:
When you run this code, you will see this output:
output
Sammy the Shark
We first define a Creature struct in this example, containing a Name
field of type string. Within the body of main, we create an instance of
*/
package main

import "fmt"

type Creature struct {
	Name string
	Age  int
	DOB  string
}

func main() {
	c := Creature{
		Name: "Sammy the Shark", Age: 25, DOB: "01/07/2000",
	}
	fmt.Println("Name:", c.Name)
	fmt.Println("Age:", c.Age)
	fmt.Println("DOB:", c.DOB)

}
