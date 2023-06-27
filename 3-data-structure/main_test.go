package main

import (
	"testing"
)

func TestDataStructure(t *testing.T) {

	user := User{
		ID:    "1",
		Name:  "Muhamad Angga",
		Email: "muhammadangga7890@gmail.com",
		Age:   23,
	}

	if user.ID != "1" {
		t.Errorf("Expected 1, got %s", user.ID)
	}

}
