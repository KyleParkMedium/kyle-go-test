package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// a := byte("hello")

	// name := "hello"
	// a := []byte("hello")
	// fmt.Println(a)
	// b, _ := json.Marshal(name)
	// fmt.Println(b)

	// params := []string{"get", "name"}
	// queryArgs := make([][]byte, len(params))
	// for i, arg := range params {
	// 	queryArgs[i] = []byte(arg)
	// }

	// fmt.Println(queryArgs)

	a := make(map[string]string)
	a["A"] = "B"
	b, _ := json.Marshal(a)
	fmt.Println(b)

	c := []byte(b)
	fmt.Println(c)
	// b, _ := json.Marshal(a)
	// fmt.Println(b)

	// c, _ := json.Marshal("hello")
	// fmt.Println(c)

	// // var q string

	// d, _ := json.Marshal(100)
	// fmt.Println(d)

	// w := make([]byte, 0)
	// fmt.Println(w)

	// var q interface{}
	// _ = json.Unmarshal(w, &q)
	// fmt.Println(q)

}
