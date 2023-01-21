package main

import "fmt"

func main() {

	userName()
	emailInput()
}

func userName() {
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Your name is:", name)

}

func emailInput() {
	var email string
	fmt.Println("Enter your email: ")
	fmt.Scanln(&email)
	fmt.Printf("your emai address is %v and your name ", email)
}
