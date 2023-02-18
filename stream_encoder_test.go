package learn_golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

/*
==	Streaming Encoder

-	Selain decoder, package json juga mendukung membuat Encoder yang bisa digunakan untuk menulis langsung JSON nya ke io.Writer
-	Dengan begitu, kita tidak perlu menyimpan JSON datanya terlebih dahulu ke dalam variable string atau []byte, kita bisa langsung tulis ke io.Writer

==	json.Encoder

-	Untuk membuat Encoder, kita bisa menggunakan function json.NewEncoder(writer)
-	Dan untuk menulis data sebagai JSON langsung ke writer, kita bisa gunakan function Encode(interface{})

*/

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Ahmad",
		MiddleName: "Mufied",
		LastName:   "Nugroho",
	}

	encoder.Encode(customer)
}
