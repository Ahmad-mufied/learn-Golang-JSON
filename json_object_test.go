package learn_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
==	JSON Object

-	Pada materi sebelumnya kita melakukan encode data seperti string, number, boolean, dan tipe data primitif lainnya
-	Walaupun memang bisa dilakukan, karena sesuai dengan kontrak interface{}, namun tidak sesuai dengan kontrak JSON
-	Jika mengikuti kontrak json.org, data JSON bentuknya adalah Object dan Array

*/

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJSON(t *testing.T) {
	customer := Customer{
		FirstName:  "Ahmad",
		MiddleName: "Mufied",
		LastName:   "Nugroho",
		Age:        22,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
