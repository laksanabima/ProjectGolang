package main 

import "fmt"

type Rumus struct {
	luas int
	rumusSegitiga float32;
	keliling int;
}

func main() {
	var pilihan int;
	rumus := new(Rumus)

	fmt.Println("Aplikasi Penghitung Luas dan Keliling Bangun Datar")
	fmt.Println("Menu Pilihan : ")
	fmt.Println("1. Persegi Panjang.")
	fmt.Println("2. Segitiga.")
	fmt.Println("3. Lingkaran.")
	fmt.Print("Pilihan menu Anda: ")

	_, err := fmt.Scanf("%d", &pilihan)
    if err != nil {
    	fmt.Println("Anda Telah Memilih menu : ", pilihan)
    }

	//for (true) {
		switch(pilihan){
		case 1:
			//var luas, keliling int;
			_, err := fmt.Scanf("%d", &pilihan)
			if err != nil {
				fmt.Println("Anda memilih Bangun Datar Persegi Panjang")
			}

			//Input value panjang persegi panjang
			var panjang, lebar int;
			print("Masukkan panjang bangun: ")
			_, err = fmt.Scanln(&panjang)
			if err != nil {
			    fmt.Println("Error: ", err)
			}
			// Input value lebar
			print("Masukkan lebar bangun: ")
			_, err = fmt.Scanln(&lebar)
			if err != nil {
			    fmt.Println("Error: ", err)
			}
			//count luas & keliling
			
			rumus.luas = panjang * lebar
			rumus.keliling = 2 * (panjang + lebar)

			fmt.Println("Luas Bangun Persegi Panjang = ", rumus.luas)
			fmt.Println("Keliling Bangun Persegi Panjang = ", rumus.keliling)
			break
		case 2:	
			//var luas, keliling int;
			_, err := fmt.Scanf("%d", &pilihan)
			if err != nil {
				fmt.Println("Anda memilih Bangun Datar Segitiga")
			}

			var alas, sisi, tinggi int;
			print("Masukkan alas bangun : ")
			_, err = fmt.Scanln(&alas)
			if err != nil {
				fmt.Println(err)
			}

			print("Masukkan sisi bangun : ")
			_, err = fmt.Scanln(&sisi)
			if err != nil {
				fmt.Println(err)
			}

			print("Masukkan tinggi bangun : ")
			_, err = fmt.Scanln(&tinggi)
			if err != nil {
				fmt.Println(err)
			}
			
			rumus.luas = (alas / 2 ) * tinggi;
			rumus.keliling = alas + tinggi + sisi;

			fmt.Println("Luas Bangun Segitiga : ", rumus.luas)
			fmt.Println("Keliling Bangun Segitiga : ", rumus.keliling)
			break
		}
	//}
}