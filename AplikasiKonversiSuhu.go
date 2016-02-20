package main

import "fmt"

type Rumus struct {
	celcius int
	reamur int
	fahrenheit int
	hasil int
}

func main() {
	var pilihan int
	rumus := new(Rumus)

	fmt.Println("Aplikasi Konversi Suhu")
	fmt.Println("Menu Pilihan : ")
	fmt.Println("1. Celcius.")
	fmt.Println("2. Reamur.")
	fmt.Println("3. Fahrenheit.")
	fmt.Print("Pilihan menu Anda: ")

	_, err := fmt.Scanf("%d", &pilihan)
    if err != nil {
    	fmt.Println("Anda Telah Memilih menu : ", pilihan)
    }
    for(true){
    	switch(pilihan){
    	case 1:
    		fmt.Println("Anda Memilih Celcius")
    		fmt.Println("1. Konversi dari Celcius ke Reamur")
    		fmt.Println("2. Konversi dari Celcius ke Fahrenheit")

    		_, err := fmt.Scanf("%d", &pilihan)
    		if err != nil {
    			fmt.Println("Anda Telah Memilih menu : ", pilihan)
    		}    		

    		switch(pilihan){
    		case 1:
    			_, err := fmt.Scanf("%d", &rumus.pilihan)
				if err != nil {
					fmt.Println("Anda memilih Konversi dari Celcius ke Reamur")
				}

				//Input value panjang persegi panjang
				print("Masukkan Celcius: ")
				_, err = fmt.Scanln(&rumus.celcius)
				if err != nil {
				    fmt.Println("Error: ", err)
				}
				rumus.hasil = 4 / 5 * rumus.celcius
				fmt.Println("Hasil Konversi = ", rumus.hasil)
				break
			}
    	case 2:    		
    		break;  
    	}  
    }
    
}