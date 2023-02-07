package golangjsonstudy

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

// ============================================================================================== //
// 2. Biasanya JSON itu datanya direpresentasikan sbg JSON oebjct / array. Skrg kita mau coba buat
// JSON Object
func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		MiddleName: "Fitzgerald",
		LastName:   "Kennedy",
		Age:        30,
		Married:    true,
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

// ============================================================================================== //
