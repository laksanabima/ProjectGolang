package main

import ("fmt";"os")

func main() {
	var pilihan int;
	var saldo int = 5000;
	for(true){
		fmt.Println("Menu:")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Menabung")
		fmt.Println("3. Mengambil")
		fmt.Println("4. Keluar")
		fmt.Println("Pilihan Menu : ")

		_, err := fmt.Scanf("%d", &pilihan)
	    if err != nil {
	    	//fmt.Println("Anda Telah Memilih menu : ", pilihan)
	    }

		switch (pilihan) {
			case 1:
				fmt.Println("Saldo Anda sebesar : Rp. ", saldo)
			break
			case 2:
				var err error;
				var menabung int;
				_, err = fmt.Scanln(&menabung)
				if err != nil {
				   fmt.Println("Anda Memilih Menabung")
				}
				print("Berapa Anda akan menabung : Rp. ")
				_, err = fmt.Scanln(&menabung)
				if err != nil {
				    fmt.Println("Error: ", err)
				}
				saldo = saldo + menabung;

				fmt.Println("Saldo anda sebesar : Rp. ", saldo)
			break
			case 3:
				var err error;
				var mengambil int;
				_, err = fmt.Scanln(&mengambil)
				if err != nil {
					fmt.Println("Anda Memilih Mengambil")
				}
				print("Berapa Anda akan mengambil: Rp. ")
				_, err = fmt.Scanln(&mengambil)
				if err != nil {
					fmt.Println(err)
				}
				saldo = saldo - mengambil
				if (saldo <= 999){
					fmt.Println("Maaf Saldo anda saat ini kurang dari Rp. 1000,-")
					fmt.Println("Anda tidak dapat melakukan penarikan uang.")
					saldo = saldo + mengambil
					fmt.Println("Salda anda sebesar : Rp. ", saldo)
				}else{
					fmt.Println("Saldo anda sebesar : Rp. ", saldo)			
				}
			break
			case 4:
		       	fmt.Println("Terimakasih telah menggunakan aplikasi ini")
		        os.Exit(0)
				break
			default:
				fmt.Println("Anda Memilih pilihan yang salah")
			break
		}
	}
	
}