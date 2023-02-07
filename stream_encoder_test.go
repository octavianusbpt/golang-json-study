package golangjsonstudy

import (
	"encoding/json"
	"os"
	"testing"
)

// ============================================================================================== //
// 8. Buat encode dari variabel bawaan golang spt map, struct atau array langsung ke stream io.Writer
func TestStreamEncoder(t *testing.T) {

	writer, _ := os.Create("CustomerFile.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "John",
		MiddleName: "Fitzgerald",
		LastName:   "Kennedy",
		Age:        30,
		Married:    true,
	}

	encoder.Encode(customer)

}

// ============================================================================================== //
