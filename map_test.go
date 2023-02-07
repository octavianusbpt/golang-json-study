package golangjsonstudy

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ============================================================================================== //
// 6. Biasanya data json itu dinamis, berubah2 atributnya, bisa nambah atau berkurang, mkny
// krg cocok kl pake struct, nah biar kita bisa nambah atau kurangin atribut, kita pake map
func TestMapDecode(t *testing.T) {

	jsonString := `{"id":"p001", "name":"Asus ROG 9000", "price":25000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])

}

func TestMapEncode(t *testing.T) {

	product := map[string]interface{}{
		"id":    "P001",
		"name":  "Asus ROG 9000",
		"price": 25000000,
	}
	bytes, err := json.Marshal(product)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}

// ============================================================================================== //
