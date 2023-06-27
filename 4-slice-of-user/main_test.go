package main

import (
	"testing"

	"github.com/google/uuid"
)

func TestFilteredUsers(t *testing.T) {
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

	if len(users) != 2 {
		t.Errorf("Expected 2, got %d", len(users))
	}
}
