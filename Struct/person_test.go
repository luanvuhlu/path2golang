package main

import (
	"testing"
)

func TestNewPerson(t *testing.T) {
	alex := newPerson()
	if alex.firstName == "Alex" {
		t.Errorf("Expected firstName is Alex, but got %v", alex.firstName)
	}
	if alex.lastName != "Anderson" {
		t.Errorf("Expected lastName is Anderson, but got %v", alex.lastName)
	}
}

func TestUpdateFirstName(t *testing.T) {
	alex := newPerson()
	jimmy := newPerson()
	alex.updateFirstName(&jimmy, "Jimmy")
	if alex.firstName != "Jimmy" {
		t.Errorf("Expected alex's firstName is Jimmy, but got %v", alex.firstName)
	}
	if jimmy.firstName != "Jimmy" {
		t.Errorf("Expected jimmy's firstName is Jimmy, but got %v", jimmy.firstName)
	}
}
