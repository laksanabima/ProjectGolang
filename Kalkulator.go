package main

import ("fmt")

func main() {
	var pilihan int;
	var bil1 int;
	var bil2 int;
	var hasil int;

	// for(true){
		fmt.Println("1. Tambah")
		fmt.Println("2. Kurang")
		fmt.Println("Masukkan Pilihan Anda : ")

		_, err := fmt.Scanf("%d", &pilihan)

        if err != nil {
        	fmt.Println(err)
        }

        switch(pilihan){
        case 1:
			_, err = fmt.Scanln(&bil1)
			if err != nil {
			    fmt.Println("Anda memilih Penjumlahan")
			}
			
			print("Bilangan Pertama: ")
			_, err = fmt.Scanln(&bil1)
			if err != nil {
			    fmt.Println("Error: ", err)
			}

			print("Bilangan Kedua: ")
			_, err = fmt.Scanln(&bil2)
			if err != nil {
			    fmt.Println("Error: ", err)
			}

        	hasil = bil1 + bil2;

        	fmt.Println("Hasil Penjumlahan = ", hasil)
        break
	    case 2:
	    	_, err = fmt.Scanln(&bil1)
			if err != nil {
			    fmt.Println("Anda memilih Pengurangan")
			}
	    	print("Bilangan Pertama: ")
			_, err = fmt.Scanln(&bil1)
			if err != nil {
			    fmt.Println("Error: ", err)
			}

			print("Bilangan Kedua: ")
			_, err = fmt.Scanln(&bil2)
			if err != nil {
			    fmt.Println("Error: ", err)
			}

        	hasil = bil1 - bil2;

        	fmt.Println("Hasil Pengurangan = ", hasil)
	    break
		default:
        	//fmt.Println("Anda memilih nomor : ",pilihan)
			fmt.Println("Anda Memilih Pilihan salah")
		break
        }      
	// }
	
}