package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b = 2, 5
	var res1 = (a + b) * (a + b) / 2 // Arithmetic operations
	a++                              // Increment a by 1
	b += 10                          // Increment b by 10
	var res2 = a ^ b                 // Bitwise XOR
	var r = 3.5
	var res3 = math.Pi * r * r // Operations on floating-point type
	fmt.Printf("res1 : %v, res2 : %v, res3 : %v\n", res1, res2, res3)

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
