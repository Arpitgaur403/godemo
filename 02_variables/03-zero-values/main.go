package main

import "fmt"

func main() {
	fmt.Println("Inside main fun")
	fmt.Println("===================")

	var (
		firstName, lastName string
		age                 int
		salary              float64
		isConfirmed         bool
	)

	fmt.Printf("firstName is : %s", firstName)
	fmt.Println()
	fmt.Printf("lastName is : %s", lastName)
	fmt.Println()
	fmt.Printf("age is : %d", age)
	fmt.Println()
	fmt.Printf("salary is : %f", salary)
	fmt.Println()
	fmt.Printf("isConfirmed is : %t", isConfirmed)
	fmt.Println()

}
