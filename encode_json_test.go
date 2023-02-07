package golangjsonstudy

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ============================================================================================== //
// 1. Testing encoding di json
func LogJson(data interface{}) {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	LogJson("Octa")
	LogJson(1)
	LogJson(true)
	LogJson([]string{"Octa", "John", "Kennedy", "Fitzgerald"})
}

// ============================================================================================== //
