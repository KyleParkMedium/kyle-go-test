package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type A struct {
	AA [32]byte
}

func main() {

	// myBytes := [32]byte{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}

	// // 바이트 배열을 16진수 문자열로 변환
	// myString := hex.EncodeToString(myBytes[:])

	// // a := [32]byte{}
	// // stringName := hex.EncodeToString(a[:])
	// // fmt.Println(stringName)

	// b := A{}
	// fmt.Println(b)
	// fmt.Println(b.AA)

	// // 16진수 문자열 출력
	// fmt.Println(myBytes)

	// fmt.Println(myString)

	// q := "123451231231231236789abcdef0112233445566778899aabbccddeeff001122334455667788"
	// w, err := hex.DecodeString(q)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(len(q))
	// fmt.Println(len(w))

	// // partition := [32]byte{}
	// stringPartition := hex.EncodeToString(partition[:])
	// fmt.Println(partition)
	// fmt.Println(stringPartition)

	// partition := "4ab21cd45c038f9b3d96c64ec46196774d762129989175922a6f01d71c3bb19e"
	partition := "0x4ab21cd45c038f9b3d96c64ec46196774d762129989175922a6f01d71c3bb19e"
	var check []byte
	if strings.HasPrefix(partition, "0x") {
		check, _ = hex.DecodeString(strings.TrimPrefix(partition, "0x"))
	} else {
		check, _ = hex.DecodeString(partition)
	}

	fmt.Println(check)
	fmt.Println(len(check))

}
