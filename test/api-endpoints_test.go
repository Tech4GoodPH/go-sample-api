package test

import (
	"testing"

	C "github.com/Tech4GoodPH/go-sample-api.git/cmd"
)

func TestSample(t *testing.T) {
	value := 1
	if 1 != value {
		t.Errorf("Expected value of 1, but got %v ", value)
	}
}

func TestHomeMessage(t *testing.T) {
	ExpectedMessage := "Tech 4 Good !"
	if C.HomeMessage() != ExpectedMessage {
		t.Errorf("Expected value of '%v', but got '%v' ", ExpectedMessage, C.HomeMessage())
	}
}
