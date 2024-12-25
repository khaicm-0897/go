package main

import (
	"fmt"
	"reflect"
	// "strconv"
)

func main() {
	var p *string
	var answer string = "4cccc2"

	// p = &answer
	fmt.Println(reflect.TypeOf(answer))
	// a, _ := strconv.ParseInt(p, 10, 64)

	fmt.Println(p)
	// fmt.Println(*p)
}
