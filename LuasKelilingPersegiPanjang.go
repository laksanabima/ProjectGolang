package main

import ("fmt"; "reflect";)

func main() {
	fmt.Println("---== Aplikasi Persegi Panjang ==---")
	fmt.Println("Masukkan Panjang")
	var p int = 12;
	var l int = 6;
	var kell int = 2*(p + l);
	var luas int = p * l;

	fmt.Println("Panjang : ", reflect.TypeOf(p), "\t =", p)
	fmt.Println("Lebar : ", reflect.TypeOf(l), "\t =", l)
	fmt.Println("Keliling Persegi Panjang : ", kell)
	fmt.Println("Luas Persegi Panjang : ", luas)
}