package golangjsonstudy

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ============================================================================================== //
// 3. Decoding JSON, dari type JSON Object ([]byte) ke type yg kita mau spt string, int, atau struct
func TestDecodeJSON(t *testing.T) {

	jsonString := `{"FirstName":"John","MiddleName":"Fitzgerald","LastName":"Kennedy","Age":30,"Married":true}`
	jsonBytes := []byte(jsonString)

	// Ini HARUS pake pointer ya, kl gk data JSON-nya bakal di decode di dlm function Unmarshal-nya
	// Selain boros memory, kita juga gk bisa dpt hasil decodingnya
	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)

}

// ============================================================================================== //
