package main

import "fmt"

func swap(x string, y string) {
	return y, x
}
func main() {
	a, b := swap("Kapoor", "vinay")
	fmt.Println(a, b)
}
