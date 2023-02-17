package learn_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
==	JSON Array

-	Selain tipe dalam bentuk Object, biasanya dalam JSON, kita kadang menggunakan tipe data Array
-	Array di JSON mirip dengan Array di JavaScript, dia bisa berisikan tipe data primitif, atau tipe data kompleks (Object atau Array)
-	Di Go-Lang, JSON Array direpresentasikan dalam bentuk slice
-	Konversi dari JSON atau ke JSON dilakukan secara otomatis oleh package json menggunakan tipe data slice

*/

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Ahmad",
		MiddleName: "Mufied",
		LastName:   "Nugroho",
		Hobbies:    []string{"Coding", "Exercise", "Learning"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

/*

==	 Decode JSON Array

-	Selain menggunakan Array pada attribute di Object
-	Kita juga bisa melakukan encode atau decode langsung JSON Array nya
-	Encode dan Decode JSON Array bisa menggunakan tipe data Slice

*/

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Ahmad","MiddleName":"Mufied","LastName":"Nugroho","Age":0,"Married":false,"Hobbies":["Cding","Exercise","Learning"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Ahamd",
		Addresses: []Address{
			{
				Street:     "Jalan Belum Ada",
				Country:    "Indonesia",
				PostalCode: "9999",
			},
			{
				Street:     "Jalan Belum Dibangun",
				Country:    "Indonesia",
				PostalCode: "8888",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDeoced(t *testing.T) {
	jsonString := `{"FirstName":"Ahamd","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Belum Dibangun","Country":"Indonesia","PostalCode":"8888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)

}

func TestOnlyJSONArrayComplexDeoced(t *testing.T) {
	jsonString := `[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Belum Dibangun","Country":"Indonesia","PostalCode":"8888"}]`
	jsonBytes := []byte(jsonString)

	address := &[]Address{}
	err := json.Unmarshal(jsonBytes, address)
	if err != nil {
		panic(err)
	}

	fmt.Println(address)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Belum Ada",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
		{
			Street:     "Jalan Belum Dibangun",
			Country:    "Indonesia",
			PostalCode: "8888",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
