package main

import "fmt"

func main() {
	fmt.Println("=========inside main========================")
	fmt.Printf("variables: \n")
	//Decalare the variable
	var myStr string = "Hello Arpit"
	var myInt int = 100
	var myFloat float64 = 45.12
	//fmt.Println(myStr, myInt, myFloat)
	fmt.Println("=================================")
	fmt.Println("myString:", myStr)
	fmt.Println("myInterger:", myInt)
	fmt.Println("myFloat:", myFloat)

	// multiple declaration

	var (
		employeeId          int    = 101
		firstName, lastName string = "Arpit", "Kr Gaur"
	)
	fmt.Println("=================================")
	fmt.Println("Emp id:", employeeId)
	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)

	// short declaration

	name := "Arpit kumar gaur"
	age, company, isProgrammer := 27, "Times Internet", true

	fmt.Println("==================================")
	fmt.Println("Name is:", name)
	fmt.Println("Name is:", name)
	fmt.Println("Age is:", age, ", Comapny is :", company, ", isprogramer :", isProgrammer)
	//fmt.Println("Com is:",name)

}
