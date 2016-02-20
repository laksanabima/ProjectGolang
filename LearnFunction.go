 package main

 import (
         "fmt"
 )

 func main() {

   	var err error
	var username string

	print("Username: ")
	_, err = fmt.Scanln(&username)
	if err != nil {
	    fmt.Println("Error: ", err)
	}

	var password string
	print("Password: ")
	_, err = fmt.Scanln(&password)
	if err != nil {
	    fmt.Println("Error: ", err)
	}

	var status string
	print("Status: ")
	_, err = fmt.Scanln(&status)
	if err != nil {
	    fmt.Println("Error: ", err)
	}

	fmt.Println(username, password, status)
	return

 }