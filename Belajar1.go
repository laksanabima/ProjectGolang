package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Bima Laksana Putra")

	var x float32 = 0.5
	var y int = int(x)
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.2f\n", decimalNumber)

	fmt.Println("this is = ", y)
	fmt.Println("this is float = ", x )
}