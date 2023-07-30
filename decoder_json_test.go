package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	_ = decoder.Decode(customer) 				// decoder.Decode(customer)

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
}

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")	// create CustomerOut.json (empty file)
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
	}

	_ = encoder.Encode(customer)				// write to CustomerOut.json file
}
