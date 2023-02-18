package learn_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
==	JSON Tag

-	Secara default atribut yang terdapat di Struct dan JSON akan di mapping sesuai dengan nama atribut  yang sama (case sensitive)
-	Kadang ada style yang berbeda antara penamaan atribute di Struct dan di JSON, misal di JSON kita ingin menggunakan snake_case, tapi di Struct, kita ingin menggunakan PascalCase
-	Untungnya, package json mendukun Tag Reflection
-	Kita bisa menambahkan tag reflection dengan nama json, lalu diikuti dengan atribut yang kita inginkan ketika konversi dari atau ke JSON

*/

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Apple Mac Book Pro",
		ImageURL: "https://apple.com",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Mac Book Pro","image_url":"https://apple.com"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
