package main

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type A struct {
	num big.Int
}

type B struct {
	num *big.Int
}

func AB() {
	q := big.NewInt(100)

	for i := 0; i < 5; i++ {
		q.Add(q, big.NewInt(100))
	}
	fmt.Println(q)
}

type TotalSupplyStruct struct {
	DocType string `json:"docType"`

	TotalSupply    int64    `json:"totalSupply"`
	ToTalSupplyBig *big.Int `json:"totalSupplyBig"`

	CreatedDate string `json:"createdDate"`
	UpdatedDate string `json:"updatedDate"`
}

func StructToMap(arg interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(arg) // Convert to a json string
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(data, &result) // Convert to a map
	return result, nil
}

func main() {

	totalSupply := TotalSupplyStruct{}

	totalSupplyMap, err := StructToMap(totalSupply)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(totalSupplyMap)
	// // AB()
	// // var a int64
	// // var b big.Int

	// // fmt.Println(a)
	// // fmt.Println(b)

	// // fmt.Println(reflect.TypeOf(a))
	// // fmt.Println(reflect.TypeOf(b))

	// // a = 5

	// // c := big.NewInt(5)

	// // fmt.Println(reflect.TypeOf(a))
	// // fmt.Println(reflect.TypeOf(b))
	// // fmt.Println(reflect.TypeOf(c))
	// // fmt.Println(reflect.TypeOf(*c))
	// // fmt.Println(reflect.TypeOf(&c))

	// // a := new(big.Int)
	// // fmt.Println(a)
	// // fmt.Println(reflect.TypeOf(a))

	// b := big.NewInt(1)
	// c := big.NewInt(2)

	// if b.Cmp(c) == 1 {
	// 	fmt.Println("b>c")
	// } else {
	// 	fmt.Println("b<c")
	// }
	// // b := big.NewInt(1)
	// // fmt.Println(b)
	// // fmt.Println(*b)
	// // fmt.Println(&b)
	// // fmt.Println(reflect.TypeOf(b))
	// // fmt.Println(reflect.TypeOf(&b))

	// // q := A{*big.NewInt(1)}
	// // fmt.Println(q)
	// // fmt.Println(reflect.TypeOf(q))

	// // w := B{big.NewInt(1)}
	// // fmt.Println(w)
	// // fmt.Println(reflect.TypeOf(w))

	// // zz := big.NewInt(11)
	// // e := B{zz}
	// // fmt.Println(e)
	// // fmt.Println(reflect.TypeOf(e))

	// // a.SetInt64(654)
	// // fmt.Println(a)

	// // a.SetInt64(6541)
	// // fmt.Println(a)

	// // a.SetBit(a, 3, 1)
	// // fmt.Println(a)

}
