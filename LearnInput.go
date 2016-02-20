package main

import ("fmt"; "bufio"; "os"; "strings";)

func main(){
	fmt.Println("Masukkan nama Anda : ")
	scn := bufio.NewReader(os.Stdin);
	data, _ := scn.ReadString('\n');
	fmt.Printf("Selamat Datang %%", strings.TrimRight(data,"\n"))
}