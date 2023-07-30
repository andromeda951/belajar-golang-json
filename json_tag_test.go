package belajar_golang_json

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

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "Mac Book Pro",
		ImageURL: "http://example.com/mac.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Mac Book Pro","image_url":"http://example.com/mac.png"}`
	// jsonString := `{"Id":"P0001","Name":j"Mac Book Pro","ImageURL":"http://example.com/mac.png"}`
	// jsonString := `{"Id":"P0001","NAME":"Mac Book Pro","IMAGE_URL":"http://example.com/mac.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)

	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.ImageURL)
}
