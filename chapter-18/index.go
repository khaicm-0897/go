package main

import "fmt"

func main() {
	// var arr [5]float32
	// log.Print("%v", arr)
	// myEmptyArray2 := [2]string{}
	// fmt.Println(myEmptyArray2)
	// myEmptyArray3 := [2]string{}
	// fmt.Println(myEmptyArray3[1])
	testArray := [2]string{"Value1", "Value2"}
	UpdateArray1(&testArray)
	fmt.Println(testArray)
}

const NewValue = "changedValue"

func UpdateArray1(array *[2]string) {
	array[0] = NewValue
	fmt.Println(array)
}
