package main

// interfaces/first-example/main.go
//...

// type Human struct {
// 	Firstname string
// 	Lastname  string
// 	Age       int
// 	Country   string
// }

// type Cat struct {
// 	Name string
// }

// type Dog struct {
// 	Name string
// }

// func (c Cat) ReceiveAffection(from Human) {
// 	fmt.Printf("The cat named %s has received affection from Human named %s\n", c.Name, from.Firstname)
// }

// func (c Cat) GiveAffection(to Human) {
// 	fmt.Printf("The cat named %s has given affection to Human named %s\n", c.Name, to.Firstname)
// }

// func (c Cat) Bark() {
// 	fmt.Printf("Meo Meo")
// }

// type DomesticAnimal interface {
// 	ReceiveAffection(from Human)
// 	GiveAffection(to Human)
// }

// type Bark interface {
// 	Bark()
// }

// func Pet(animal DomesticAnimal, human Human) {
// 	animal.GiveAffection(human)
// 	animal.ReceiveAffection(human)
// }

// func PetBark(animal Bark) {
// 	animal.Bark()
// }

// func main() {

// 	// Create the Human
// 	var john Human
// 	john.Firstname = "John"

// 	// Create a Cat
// 	var c Cat
// 	c.Name = "Maru"

// 	// then a dog
// 	var d Dog
// 	d.Name = "Medor"

// 	Pet(c, john)
// 	Pet(d, john)
// 	PetBark(c)
// }

type Snake struct {
	Name string
}

type Human struct {
	Firstname string
	Lastname  string
	Age       int
	Country   string
}

// we do not implement the methods ReceiveAffection and GiveAffection intentionally
//...

type DomesticAnimal interface {
	ReceiveAffection(from Human)
	GiveAffection(to Human)
}

func Pet(animal DomesticAnimal, human Human) {
	animal.GiveAffection(human)
	animal.ReceiveAffection(human)
}

func main() {

	var john Human
	john.Firstname = "John"

	var snake Snake
	snake.Name = "Joe"

	Pet(snake, john)
}
