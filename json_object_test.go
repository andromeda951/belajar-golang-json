package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode int
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Maried     bool
	Hobbies    []string
	Addresses  []Address
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
		Age:        27,
		Maried:     true,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
		Age:        27,
		Maried:     true,
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplext(t *testing.T) {
	customer := Customer{
		FirstName: "Eko",
		Hobbies:   []string{"Gaming", "Reading", "Coding"},
		Addresses: []Address{
			{
				Street:     "Jalan a",
				Country:    "Indonesia",
				PostalCode: 12345,
			},
			{
				Street:     "Jalan b",
				Country:    "Indonesia",
				PostalCode: 56789,
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestDecodeJsonArray(t *testing.T) {
	jsonArray := `[{"Street":"Jalan a","Country":"Indonesia","PostalCode":12345},{"Street":"Jalan b","Country":"Indonesia","PostalCode":56789}]`
	jsonBytes := []byte(jsonArray)

	addresses := &[]Address{}
	json.Unmarshal(jsonBytes, addresses)

	fmt.Println(addresses)
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"Kurniawan","LastName":"Khannedy","Age":27,"Maried":true,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"","LastName":"","Age":0,"Maried":false,"Hobbies":["Gaming","Reading","Coding"],"Addresses":[{"Street":"Jalan a","Country":"Indonesia","PostalCode":12345},{"Street":"Jalan b","Country":"Indonesia","PostalCode":56789}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Addresses)
}

func TestOnlyJsonArrayComplext(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan a",
			Country:    "Indonesia",
			PostalCode: 12345,
		},
		{
			Street:     "Jalan b",
			Country:    "Indonesia",
			PostalCode: 56789,
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))

}
