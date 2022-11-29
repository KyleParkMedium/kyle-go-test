package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
)

func main() {

	// a := [][]byte{{"aa", "bb"}}
	// fmt.Println(a)

	// type BlockHeader struct {
	// 	Number       uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// 	PreviousHash []byte `protobuf:"bytes,2,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
	// 	DataHash     []byte `protobuf:"bytes,3,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	// }

	// type BlockData struct {
	// 	Data [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// }

	// type Block struct {
	// 	Header *BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// 	Data   *BlockData   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// }

	// txids := []string{"tx1", "tx2", "tx3"}

	// a := []byte("a")
	// b := []byte("b")
	// txids2 := Block{BlockHeader{1, a, b}}

	// fmt.Println(txids)
	// fmt.Println(txids2)

	a := ConstructRandomBytes(32)
	fmt.Println(a)

	pubSimulationResults := [][]byte{}

	type Test struct {
		Name int
		Age  int
	}
	q := Test{1, 2}

	ab, _ := json.Marshal(q)
	fmt.Println(q)
	fmt.Println(ab)

	qq := new(Test)
	abc := json.Unmarshal(ab, qq)
	fmt.Println(abc)
	fmt.Println(qq)

	pubSimulationResults = append(pubSimulationResults, ab)
	pubSimulationResults = append(pubSimulationResults, ab)
	pubSimulationResults = append(pubSimulationResults, ab)

	fmt.Println(pubSimulationResults)

	fmt.Println(len(pubSimulationResults))
}

// ConstructRandomBytes constructs random bytes of given size
func ConstructRandomBytes(size int) []byte {
	value := make([]byte, size)
	_, err := rand.Read(value)
	if err != nil {
		fmt.Println(err)
	}
	return value
}
