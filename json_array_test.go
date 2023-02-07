package golangjsonstudy

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ============================================================================================== //
// 4. JSON dlm bntk array
func TestJSONArray(t *testing.T) {

	customer := Customer{
		FirstName:  "Octavianus",
		MiddleName: "Liemka",
		LastName:   "Kennedy",
		Hobbies:    []string{"Gaming", "Sleeping", "Eating"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}

func TestJSONArrayDecode(t *testing.T) {

	jsonString := `{"FirstName":"Octavianus","MiddleName":"Liemka","LastName":"Kennedy","Age":0,"Married":false,"Hobbies":["Gaming","Sleeping","Eating"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}
	fmt.Println(customer)

}

// ============================================================================================== //

// ============================================================================================== //
// Contoh Array in Array

func TestJSONArrayComplex(t *testing.T) {

	customer := Customer{
		FirstName:  "Octavianus",
		MiddleName: "Liemka",
		LastName:   "Kennedy",
		Hobbies:    []string{"Gaming", "Sleeping", "Eating"},
		Addresses: []Address{
			{
				Street:     "Mapple Street 10",
				Country:    "England",
				PostalCode: "345678",
			},
			{
				Street:     "Avocado Street 45",
				Country:    "Scotland",
				PostalCode: "98324",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}

func TestJSONArrayComplexDecode(t *testing.T) {

	jsonString := `{"FirstName":"Octavianus","MiddleName":"Liemka","LastName":"Kennedy","Age":0,"Married":false,"Hobbies":["Gaming","Sleeping","Eating"],"Addresses":[{"Street":"Mapple Street 10","Country":"England","PostalCode":"345678"},{"Street":"Avocado Street 45","Country":"Scotland","PostalCode":"98324"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}
	fmt.Println(customer)

}

func TestOnlyJSONArrayDecode(t *testing.T) {

	jsonString := `[{"Street":"Mapple Street 10","Country":"England","PostalCode":"345678"},{"Street":"Avocado Street 45","Country":"Scotland","PostalCode":"98324"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)

	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)

}

func TestOnlyJSONArray(t *testing.T) {

	addresses := []Address{
		{
			Street:     "Mapple Street 10",
			Country:    "England",
			PostalCode: "345678",
		},
		{
			Street:     "Avocado Street 45",
			Country:    "Scotland",
			PostalCode: "98324",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))

}

// ============================================================================================== //
