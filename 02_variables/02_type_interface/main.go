package main

import "fmt"

func main() {
	fmt.Println("Inside main function")
	fmt.Println("=============================")
	// Type inference
	var name = "Arpit gaur" // Type declaration is optional here.
	fmt.Printf("Variable 'name' is of type %T\n", name)

	// Multiple variable declarations with inferred types

	var firstName, lastName, age, salary = "Arpit", "gaur", 27, 50000.00
	fmt.Println("=============================")
	//fmt.Printf("firstname: %T", "llastname: %T", "Age: %T", "salary: %T \n", firstName, lastName, age, salary)
	fmt.Println("=============================")
	fmt.Printf(firstName, lastName, age, salary)

}
