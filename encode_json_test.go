package learn_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
==	Pengenalan JSON

-	JSON singkatan dari JavaScript Object Notation, merupakan struktur format data yang bentuknya seperti Object di JavaScript
-	JSON merupakan struktur format data yang paling banyak digunakan saat kita membuat RESTful API
-	Dan pate kelas ini kita akan menggunakan JSON juga
-	https://www.json.org/json-en.html

==	Package json

-	Go-Lang sudah menyediakan package json, dimana kita bisa menggunakan package ini untuk melakukan konversi data ke JSON (encode) atau sebaliknya (decode)
-	https://pkg.go.dev/encoding/json

==	Encode JSON

-	Go-Lang telah menyediakan function untuk melakukan konversi data ke JSON, yaitu menggunakan function json.Marshal(interface{})
-	Karena parameter nya adalah interface{}, maka kita bisa masukan tipe data apapun ke dalam function Marshal

*/

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("Ahmad")
	logJson(1)
	logJson(true)
	logJson([]string{"Ahmad", "Mufied", "Nugroho"})
}
