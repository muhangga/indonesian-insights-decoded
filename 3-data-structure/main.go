package main

import (
	"fmt"
)

type User struct {
	ID    string
	Name  string
	Email string
	Age   int
}

func main() {

	user := User{
		ID:    "1",
		Name:  "Muhamad Angga",
		Email: "muhammadangga7890@gmail.com",
		Age:   23,
	}

	fmt.Println("ID User : ", user.ID)
	fmt.Println("Name User : ", user.Name)
	fmt.Println("Email User : ", user.Email)
	fmt.Println("Age User : ", user.Age)
}
