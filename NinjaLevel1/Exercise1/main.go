package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x)
	fmt.Printf("%v\n", y)
	fmt.Println(z)

	fmt.Printf("%d\t%s\t%t", x, y, z)
}
