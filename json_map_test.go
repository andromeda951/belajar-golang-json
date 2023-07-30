package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Mac Book Pro","price":2000000}`
	jsonByte := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonByte, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestJsonMapDecode2(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"","LastName":"","Age":0,"Maried":false,"Hobbies":["Gaming","Reading","Coding"],"Addresses":[{"Street":"Jalan a","Country":"Indonesia","PostalCode":12345},{"Street":"Jalan b","Country":"Indonesia","PostalCode":56789}]}`
	jsonByte := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonByte, &result)

	fmt.Println(result)
}
// map biasanya dipakai untuk decode json yang jumlah atributnya berubah-ubah dan semua data json akan menjadi map


// map jika di encode ke json menjadi object dan slice tetap menjadi array
func TestJsonMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Appel Mac Book Pro",
		"price": 20000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonMapEncode2(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Appel Mac Book Pro",
		"price": 20000000,
		"hobbies": []string{"gaming, reading"},
		"customer": Customer{
			FirstName: "budi",
		},
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}