package main

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID    string
	Name  string
	Email string
	Age   int
}

func filteredUser(users []User) []User {
	var filteredUser []User

	for _, user := range users {
		if user.Age >= 18 {
			filteredUser = append(filteredUser, user)
		}
	}

	return filteredUser
}

func main() {
	users := []User{
		{
			ID:    uuid.NewString(),
			Name:  "Muhamad Angga",
			Email: "muhammadangga7890@gmail.com",
			Age:   23,
		}, {
			ID:    uuid.NewString(),
			Name:  "Jhon Doe",
			Email: "jhondoe@gmail.com",
			Age:   17,
		}, {
			ID:    uuid.NewString(),
			Name:  "Test User",
			Email: "testuser@gmail.com",
			Age:   15,
		}, {
			ID:    uuid.NewString(),
			Name:  "Test User 2",
			Email: "testuser2@gmail.com",
			Age:   20,
		},
	}

	users = filteredUser(users)

	for _, user := range users {
		fmt.Println("===============================================")
		fmt.Println("ID User : ", user.ID)
		fmt.Println("Name User : ", user.Name)
		fmt.Println("Email User : ", user.Email)
		fmt.Println("Age User : ", user.Age)
		fmt.Println("===============================================")
	}
}
