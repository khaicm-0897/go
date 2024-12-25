package main

import "fmt"

func main() {
	s := make([]int, 3)
	s[0] = 12
	s[2] = 3
	s = append(s, 5)

	fmt.Println(s, len(s), cap(s))

	customers := [4]string{"John Doe", "Helmuth Verein", "Dany Beril", "Oliver Lump"}
	// slice the array
	customersSlice := customers[0:1]
	fmt.Println(customersSlice)
	// modify original array
	customers[0] = "John Doe Modified"
	fmt.Println("After modification of original array")
	fmt.Println(customersSlice)

	// Slicing a string
	hotelName := "Go Dev Hotel"
	str := hotelName[0:6]
	fmt.Println(str)
	hotelName = "Java Dev Hotel"
	fmt.Println(str)

	names := [5]string{"join", "bob", "nik", "clar", "jax"}
	sli := names[1:3]

	fmt.Println(len(sli))
	fmt.Println(cap(sli))

	// slices in function params

	languages := []string{"Java", "PHP", "C"}
	fmt.Println("Capacity :", cap(languages))
	// Capacity : 3

	// call function
	addGo(languages)

	fmt.Println("Capacity :", cap(languages))
	// Capacity : 3
	fmt.Println(languages)
	// [Java PHP C]
	// what ! , where is Go ?????

}
func addGo(languages []string) {
	languages = append(languages, "Go")
	fmt.Println("in function, capacity", cap(languages))
}
