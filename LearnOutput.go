package main

import "fmt"

func main() {
	var i int = 21
	var j bool = true;

	fmt.Printf("%v \n", i)        // menampilkan nilai i = 21
	fmt.Printf("%T \n", i)        // menampilkan tipe dta dari varibel i
	fmt.Printf("%% \n")           // menampilkan tanda %
	fmt.Printf("%t \n\n", j)      // menampilkan nilai boolean j

	fmt.Printf("%b \n", i)        // menampilkan nilai boolen j
	fmt.Printf("%c \n", '\u042F') // menampilkan unicode Russia
	fmt.Printf("%d \n", i)        // menampilkan nilai base 10 : 21
	fmt.Printf("%o \n", i)        // menampilkan nilai base 8 : 25
	fmt.Printf("%x \n", 15)       // menampilkan nilai base 16 : f
	fmt.Printf("%X \n", 15)       // menampilkan nilai base 16 : F

	var k float64 = 123.456;
	fmt.Printf("%f \n", k)        // menampilkan float : 123.456000
	fmt.Printf("%E \n", k)        // menampilkan float scientific : 1.2345670E+02
	fmt.Printf("%T \n", k)
}