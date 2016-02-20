package main 

import ("fmt")

func main() {
	var karakter byte = '?'
	switch (karakter){
		case 'A' : fmt.Println("Karakter A")
		case 'B' : fmt.Println("Karakter B")
		case 'C' : fmt.Println("Karakter C")
	}
	
	// fmt.Println("Masukkan Nilai Siswa : ")
	// scn := bufio.NewReader(os.Stdin);
	// data, _ := scn.ReadString('\n')
	var nilai int32 = 91; 

	switch {
		case nilai > 90 : fmt.Println("Nilai Anda A")
		case nilai > 70 : fmt.Println("Nilai Anda B")
		case nilai > 60 : fmt.Println("Nilai Anda C")
	}

	for i := 0; i < 10; i++{
		switch {
			case i%2 == 0 : fmt.Println("Nilai 1 % 2 ",i)
			case i%5 == 0 : fmt.Println("Nilai 1 % 5 ",i)
		}
	}
}