package main

import "fmt"

// func main() {
// 	fmt.Println("just testing")
// 	deferex.LenAndUpper("Red")
// }

// struct 예제
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "pizza"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico.name)
}
