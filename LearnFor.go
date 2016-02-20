package main 

import "fmt"

func main() {
	for i :=0; i<5; i++ {
		fmt.Println("Nilai i = ", i)  // Type 1
	}

	var j int = 0;
	for j<5 {
		fmt.Println("Nilai j = ", j)  // Type 2
		j++
	}

	for pos, char := range "CABAPBO" { 
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}
	for {
		if j++; j <= (10) {
			fmt.Println("Nilai j = ", j)
		}else{
			break
		}
	}
}