package golangjsonstudy

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// ============================================================================================== //
// 7. Biasanya kl mau baca data json itu bentukannya jrg bgt map atau array, mungkin jsonObject
// tapi paling sering langsung dari network, bentukan file, dll. Di golang kita bisa langsung
// baca data tersebut dgn package os ioreader(input output reader)
func TestStreamDecoder(t *testing.T) {

	// Balikannya ada hasil readernya dan error, buat sementara error handling skip dl
	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)

}

// ============================================================================================== //
