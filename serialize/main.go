package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/rlp"
)

type Ex struct {
	// Num *big.Int `json:"num"`
	A1 int64
	A2 int64
}

type Ex2 struct {
	By []byte `json:"by"`
}

func main() {

	// a.Num = big.NewInt(100)

	// a := big.NewInt(100)

	a := []byte("hello")

	fmt.Println(a)

	q, _ := json.Marshal(a)

	fmt.Println(q)

	// w, _ := rlp.EncodeToBytes([]interface{}{a})
	w, _ := rlp.EncodeToBytes(a)

	fmt.Println(w)
	fmt.Println(reflect.TypeOf(w))

	var r interface{}
	if err := rlp.DecodeBytes(w, &r); err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
