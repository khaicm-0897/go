package main

import (
	"fmt"
)

// import (
// 	"fmt"
// 	"reflect"
// 	// "strconv"
// )

// type Command struct {
// 	Name  string
// 	Flags []Flag
// }

// type Flag struct {
// 	Name      string
// 	Shorthand string
// }

// func getCommand() *Command {
// 	return &Command{Name: "Hello", Flags: []Flag{{"version", "v"}}}
// }

// func change(cmd *Command) {
// 	for _, flg := range cmd.Flags {
// 		flg.Name = "Changed"
// 	}
/*
	another way to change data
	for index, _ := range cmd.Flags {
		cmd.Flags[index].Name = "Changed"
	}
*/

// }

// use pointer change data
/*
type Command struct {
    Name  string
    Flags []*Flag
}

type Flag struct {
    Name      string
    Shorthand string
}

func getCommand() *Command {
    return &Command{Name: "Hello", Flags: []*Flag{{"version", "v"}}}
}

func change(cmd *Command) {
    for _, flg := range cmd.Flags {
        flg.Name = "Changed"
    }

}

func main() {
    cmd := getCommand()
    change(cmd)
    fmt.Println(cmd.Flags[0])

}
*/

// func main() {
// 	cmd := getCommand()
// 	change(cmd)
// 	fmt.Println(cmd.Flags[0])

// 	var p *string
// 	var a *int
// 	var answer string = "4cccc2"
// 	var name int = 50
// 	a = &name

// 	// p = &answer
// 	fmt.Println(reflect.TypeOf(answer))
// 	fmt.Println(reflect.TypeOf(name))
// 	// a, _ := strconv.ParseInt(p, 10, 64)

// 	fmt.Println(p)
// 	fmt.Println(*a)
// 	// fmt.Println(*p)
// }

// slice

// func main() {
// 	countries := []string{"Vietnam", "USA", "Japan", "Korea"}
// 	fmt.Println(countries)
// 	fmt.Println("--aaaaaa--")
// 	addCountries(countries)
// 	addCountriesWithPointer(&countries)
// 	upper(countries)
// 	fmt.Println("--bbbbb--")
// 	fmt.Println(countries)
// }

// func addCountries(countries []string) {
// 	countries = append(countries, "China")
// 	fmt.Println(countries)
// }

// func addCountriesWithPointer(countries *[]string) {
// 	*countries = append(*countries, "China")
// 	fmt.Println(countries)
// }

// func upper(countries []string) {
// 	for i, e := range countries {
// 		// fmt.Println(k)
// 		fmt.Println(i)
// 		countries[i] = strings.ToUpper(e)
// 	}
// }

// go pointer with slice and receiver with struct

// func main() {
// 	cart := Cart{
// 		ID: 			 "123",
// 		CreateDate: time.Now(),
// 	}

// 	cartPtr := &cart
// 	cartPtr.Iteems = []Item{
// 		{SKU: "123", Quantity: 2},
// 		{SKU: "456", Quantity: 3},
// 	}
// 	log.Println(cartPtr)
// }

// type Cart struct {
// 	ID: string
// 	CreateDate: time.Time
// 	Items: []*Item
// }

// type Item struct {
// 	SKU: string
// 	Quantity: int
// }

// Method with pointer receiver

// pointer/methods-pointer-receivers/main.go

type Cat struct {
	Color string
	Age   uint8
	Name  string
}

func (cat *Cat) Meow() {
	fmt.Println("Meooooow")
}

func (cat *Cat) Rename(newName string) {
	cat.Name = newName
}

func (cat Cat) RenameV2(newName string) {
	cat.Name = newName
	fmt.Println("Inside RenameV2:", cat.Name)
	fmt.Println(cat.Name)
}

func main() {
	cat := Cat{Color: "blue", Age: 8, Name: "Milow"}
	cat.Rename("Bob")
	fmt.Println(cat.Name)
	// Bob

	cat.RenameV2("Ben")
	fmt.Println(cat.Name)
	// Bob
}
