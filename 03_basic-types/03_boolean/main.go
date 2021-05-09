package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("===============================")

	var myBoolean bool = true
	var anotherBoolean = false // Inferred type

	var truth = 3 <= 5
	var falsehood = 10 != 10

	// Short Circuiting
	var resp1 = 10 > 20 && 5 == 5     // Second operand is not evaluated since first evaluates to false
	var resp2 = 2*2 == 4 || 10%3 == 0 // Second operand is not evaluated since first evaluates to true

	fmt.Println(myBoolean, anotherBoolean, truth, falsehood, resp1, resp2)

}
