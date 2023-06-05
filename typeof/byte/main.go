package main

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
)

type A struct {
	Name  int
	Value int
}

func main() {

	// hex.
	// a, _ := hex.DecodeString("1b")
	// fmt.Println(a)
	v64, _ := strconv.ParseUint("1b", 16, 32)
	fmt.Println(v64)
	// var data [1]byte
	// data[0] = 1111

	// fmt.Println(data)

	// a := []byte("0xaaa")
	// fmt.Println(a)
	// a := A{1, 2}

	// b, _ := hex.DecodeString("949fa6dcead10fdde842989bf72b9ed338b4ba496c6d757ccf44ef9ae83824e2")
	// fmt.Println(b)
	// fmt.Println(crypto.Keccak256(b))
	fmt.Println(crypto.Keccak256([]byte("949fa6dcead10fdde842989bf72b9ed338b4ba496c6d757ccf44ef9ae83824e2")))

	// 0x37617282f682b2b3bac5ef5ad9256694d08bc968dcf72a2387be583e61d78b3a
	fmt.Println(hex.DecodeString("37617282f682b2b3bac5ef5ad9256694d08bc968dcf72a2387be583e61d78b3a"))
	// b, err := json.Marshal(a)

	fmt.Println(hex.DecodeString("5f8c1af10af45c24cc9255bb7ba38dfec7f82765576e7c2c8ef354440a04c07c"))

	// fmt.Println(err)
	// fmt.Println(b)

	// aa := A{}

	// // var aa interface{}
	// // aa := make(interface{})

	// err = json.Unmarshal(b, &aa)
	// fmt.Println(aa)
}
