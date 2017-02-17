// Closure
package main

import (
	"fmt"
)

func main() {
	n1 := getNumber(1, 50)
	n2 := getNumber(2, 0)
	n3 := getNumber(3, 3)

	fmt.Println("N1----------")
	fmt.Println(n1())
	fmt.Println(n1())
	fmt.Println(n1())
	fmt.Println("N2----------")
	fmt.Println(n2())
	fmt.Println(n2())
	fmt.Println(n2())
	fmt.Println("N3----------")
	fmt.Println(n3())
	fmt.Println(n3())
	fmt.Println(n3())
}

func getNumber(n int, init int) func() int {
	var i int = init
	return func() int {
		i += n
		return i
	}
}
