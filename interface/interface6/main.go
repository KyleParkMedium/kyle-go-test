package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
)

type A struct {
	Name      string   `json:"name"`
	Age       *big.Int `json:"age"`
	AgeString string   `json:"ageString"`
}

func main() {

	set := big.NewInt(0)
	_, ok := set.SetString("15162319880118241914981489149891848914891894", 10)
	if !ok {
		fmt.Println("Failed to set big integer value")
		return
	}
	s := set.String()

	data := map[string]interface{}{
		"name": string("John"),
		"age":  set,
		// "age":  big.NewInt(111111111454455451),
		// "age":  int64(1111111111111111111),
		"ageString": s,
	}
	// fmt.Println(data)
	// bytes := []byte(data)
	jsbytes, _ := json.Marshal(data)

	c := make(map[string]interface{})
	_ = json.Unmarshal(jsbytes, &c)
	// fmt.Println(c)

	for q, w := range c {
		fmt.Println(reflect.TypeOf(w))
		fmt.Println(q, w)
	}

	// // fmt.Println(bytes)
	// fmt.Println(jsbytes)

	// var b interface{}
	// _ = json.Unmarshal(jsbytes, &b)
	// fmt.Println(b)

	// for q, w := range b {
	// 	fmt.Println(q, w)
	// }

	// d := A{}
	// _ = json.Unmarshal(jsbytes, &d)
	// fmt.Println(d)

	// fmt.Println(d.Age)
	// fmt.Println(reflect.TypeOf(d.Age))
	// for q, w := range d {
	// 	fmt.Println(reflect.TypeOf(q))
	// 	fmt.Println(reflect.TypeOf(w))
	// 	fmt.Println(q, w)
	// }

	// e := "A"
	// r, _ := json.Marshal(e)
	// w := []byte("A")
	// fmt.Println(w)
	// fmt.Println(r)

	// // x
	// var a map[string]interface{}
	// // o
	// q := make(map[string]interface{})
	// fmt.Println(q)

	// // b := "A"
	// // a["key"] = b

	// fmt.Println(a)
}
