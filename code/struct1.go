package main

import "fmt"

func main() {
	// START OMIT
	var p = struct {
		Name string
		Age  int
	}{"Example", 36}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Println("Name:", p.Name)
	// END OMIT
}
