package main

import "fmt"

func main() {
	var myInt8 int8 = 99
	/*
	  When you don't declare any type explicitly, the type inferred is `int`
	  (The default type for integers)
	*/
	var myInt = 1204

	var myUint uint = 501

	var myHexNumber = 0xFF  // Use prefix '0x' or '0X' for declaring hexadecimal numbers
	var myOctalNumber = 034 // Use prefix '0' for declaring octal numbers

	var myFloat32 float32 = 4.5
	var myFloat = 9.12 // // Type inferred as `float64` (the default type for floating-point numbers)

	fmt.Printf("%d, %d, %d, %#x, %#o %f %f\n", myInt8, myInt, myUint, myHexNumber, myOctalNumber, myFloat32, myFloat)

	fmt.Println()
	fmt.Println("===============================")
	var myByte byte = 'T'
	var myRune rune = '♥'

	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)

	fmt.Println()
	fmt.Println("===============================")

	var website = "\thttps://github.com/Arpitgaur403/godemo\t\n"

	// Raw String (Can span multiple lines. Escape characters are not interpreted)
	var siteDescription = `\t\t deome golang development.\t\n`

	fmt.Println(website, siteDescription)

}
