package golangjsonstudy

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

// ============================================================================================== //
// 5. Biasanya kl di json formatnya pake snake case. Contoh image_url
// Kl di golang pakenya pascal case. Contoh: ImageURL
// Biar bisa saling ngerti disediain tag
func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "1234",
		Name:     "Computer",
		ImageURL: "http://image.com/computer",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {

	jsonString := `{"id":"1234","name":"Computer","image_url":"http://image.com/computer"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	err := json.Unmarshal(jsonBytes, product)

	if err != nil {
		panic(err)
	}
	fmt.Println(product)

}

// ============================================================================================== //
