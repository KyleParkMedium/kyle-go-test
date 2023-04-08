package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
)

type A struct {
	Name      string   `json:"name"`
	Age       *big.Int `json:"age"`
	Age2      *big.Int `json:"age2"`
	AgeString string   `json:"ageString"`
}

// Hex2Bytes returns the bytes represented by the hexadecimal string str.
func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}

type B struct {
	Num *big.Float `json:"num"`
	// Num string `json:"num"`
	// Num *big.Int `json:"num"`
}

func main() {

	// 	value := big.NewInt(0)
	// 	_, ok := value.SetString("15162319880118241914981489149891848914891894", 10)
	// 	if !ok {
	// 	}

	// 	encoded, _ := json.Marshal(value)
	// 	fmt.Println(encoded)
	// 	fmt.Println(reflect.TypeOf(encoded))

	// 	s := value.String()

	// 	encoded, _ = json.Marshal(s)
	// 	fmt.Println(encoded)
	// 	fmt.Println(reflect.TypeOf(encoded))

	// 	value2 := big.NewFloat(0)
	// 	_, ok = value2.SetString("15162319880118241914981489149891848914891894")
	// 	if !ok {
	// 	}

	// 	encoded, _ = json.Marshal(value2)
	// 	fmt.Println(encoded)
	// 	fmt.Println(reflect.TypeOf(encoded))

	// 	value3 := uint64(33)
	// 	encoded3, _ := json.Marshal(value3)
	// 	fmt.Println(encoded3)
	// 	fmt.Println(reflect.TypeOf(encoded3))

	// 	fmt.Println(value)
	// 	// fmt.Println(value2)
	// 	// qw := value1.String()

	// 	a := "111"

	// 	data := map[string]interface{}{
	// 		"num": a,
	// 	}

	// )	b, _ := json.Marshal(data)

	// 	// c := make(map[string]interface{})
	// 	// _ = json.Unmarshal(b, &c)

	// 	d := B{}
	// 	_ = json.Unmarshal(b, &d)
	// 	fmt.Println(d)
	// 	fmt.Println(reflect.TypeOf(d)

	// // for q, w := range c {
	// // 	fmt.Println(reflect.TypeOf(w))
	// // 	fmt.Println(q, w)
	// // }

	// // // Create a new big.Int value
	// // value1 := big.NewInt(1234561231232131789)
	// // _, ok := value1.SetString("15162319880118241914981489149891848914891894", 10)

	// // value1 := int64(33)
	// // value1 := "3.3"
	// value1 := big.NewInt(33)
	// // value1 := big.NewFloat(3.3)

	// // value1 := big.NewInt(0)
	// // _, ok := value1.SetString("15162319880118241914981489149891848914891894", 10)
	// // qw := value1.String()

	// // Serialize the big.Int value to JSON
	// encoded, _ := json.Marshal(value1)
	// fmt.Println(encoded)
	// fmt.Println(reflect.TypeOf(encoded))

	// // encoded, _ = json.Marshal(qw)
	// // fmt.Println(encoded)
	// // fmt.Println(reflect.TypeOf(encoded))

	// // // Convert a hex string to a []byte slice
	// // hexString := "deadbeef"
	// // bytes := Hex2Bytes(hexString)

	// // // Set a big.Int value from the []byte slice
	// // value := new(big.Int)
	// // value.SetBytes(bytes)
	// // fmt.Println(value)
	// // fmt.Println(reflect.TypeOf(value))

	// // account := new(A)
	// // fmt.Println(account)
	// // fmt.Println(reflect.TypeOf(account))

	set := big.NewInt(0)
	_, ok := set.SetString("15162319880118241914981489149891848914891894", 10)
	if !ok {
		fmt.Println("Failed to set big integer value")
		return
	}

	// s := set.String()

	// // // uintBig := big.NewInt(0)
	// // // a := uintBig.SetUint64(5555555616515616165)
	// // // if !ok {
	// // // 	fmt.Println("Failed to set big integer value")
	// // // 	return
	// // // }

	// // // set2 := big.NewFloat(11.1)

	// // // set2 := big.NewFloat(3.33)
	set2 := int64(11111)
	aa := set.String()

	data := map[string]interface{}{
		"name":      string("John"),
		"age":       set,
		"age2":      set2,
		"ageString": aa,
		// "age2": 11.1,
		// "age":  big.NewInt(111111111454455451),
		// "age":  int64(1111111111111111111),

	}

	// // // fmt.Println(reflect.TypeOf(data["age2"]))

	// // // fmt.Println(data)
	// // // bytes := []byte(data)
	jsbytes, _ := json.Marshal(data)

	// // c := make(map[string]interface{})
	// // _ = json.Unmarshal(jsbytes, &c)
	// // // fmt.Println(c)

	// // // // t := ""
	// // // // r := big.NewFloat(0)

	// // for q, w := range c {
	// // 	fmt.Println(reflect.TypeOf(w))
	// // 	fmt.Println(q, w)
	// // }
	// // // 	// if q == "ageString" {
	// // // 	// 	fmt.Println(w)
	// // // 	// 	t = w.(string)
	// // // 	// }
	// // // 	// if q == "age" {
	// // // 	// 	fmt.Println(w)
	// // // 	// 	fmt.Println(reflect.TypeOf(w))
	// // // 	// 	r.SetFloat64(w.(float64))
	// // // 	// 	fmt.Println(r)

	// // // 	// 	var s string = strconv.FormatFloat(w.(float64), 'E', -1, 32)
	// // // 	// 	fmt.Println(reflect.TypeOf(s))
	// // // 	// 	fmt.Println(s)

	// // // 	// 	// fmt.Println(=
	// // // 	// 	// r = w.(string)

	// // // 	// }

	// // // }

	// // // // fmt.Println(t)
	// // // // fmt.Println(r)

	// // // // z := big.NewInt(0)
	// // // // x, y := z.SetString(t, 10)
	// // // // fmt.Println(z)
	// // // // fmt.Println(x)
	// // // // fmt.Println(y)

	// // // // fmt.Println(reflect.TypeOf(z))
	// // // // fmt.Println(reflect.TypeOf(x))
	// // // // fmt.Println(reflect.TypeOf(y))

	// // // // // fmt.Println(bytes)
	// // // // fmt.Println(jsbytes)

	// // // // var b interface{}
	// // // // _ = json.Unmarshal(jsbytes, &b)
	// // // // fmt.Println(b)

	// // // // for q, w := range b {
	// // // // 	fmt.Println(q, w)
	// // // // }

	type Example struct {
		Num string `json:"num"`
	}

	ex := Example{}
	d := A{}
	_ = json.Unmarshal(jsbytes, &d)
	fmt.Println(d)
	// // // fmt.Println(d.Age)
	// // // // fmt.Println(d.Age2)
	// // // fmt.Println(d.AgeString)
	// // // fmt.Println(reflect.TypeOf(d.Age))
	// // // fmt.Println(reflect.TypeOf(d.AgeString))

	// // // z := big.NewInt(0)
	// // // _, ok = z.SetString(d.AgeString, 10)
	// // // if !ok {
	// // // 	fmt.Println(ok)

	// // }
	// // fmt.Println(z)
	// // fmt.Println(reflect.TypeOf(z))

	// // fmt.Println(reflect.ValueOf(z))
	// // fmt.Println(reflect.ValueOf(d.Age))
	// // if reflect.ValueOf(z) == reflect.ValueOf(d.Age) {
	// // 	fmt.Println("Equal")
	// // }

	// // fmt.Println(reflect.DeepEqual(z, d.Age))
	// // fmt.Println(reflect.DeepEqual(reflect.ValueOf(z), reflect.ValueOf(d.Age)))
	// // fmt.Println(d.Age)
	// // fmt.Println(reflect.TypeOf(d.Age))
	// // for q, w := range d {
	// // 	fmt.Println(reflect.TypeOf(q))
	// // 	fmt.Println(reflect.TypeOf(w))
	// // 	fmt.Println(q, w)
	// // }

	// // e := "A"
	// // r, _ := json.Marshal(e)
	// // w := []byte("A")
	// // fmt.Println(w)
	// // fmt.Println(r)

	// // // x
	// // var a map[string]interface{}
	// // // o
	// // q := make(map[string]interface{})
	// // fmt.Println(q)

	// // // b := "A"
	// // // a["key"] = b

	// // fmt.Println(a)
}
