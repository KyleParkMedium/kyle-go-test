package main

import (
	"encoding/json"
	"fmt"

	"crypto2/ecdsa2"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"golang.org/x/text/encoding/unicode"
)

func dev() {
	data := []byte("Hello, world!")
	hash1 := sha3.NewLegacyKeccak256()
	hash1.Write(data)
	fmt.Printf("keccak256 hash of %s is %x\n", string(data), hash1.Sum(nil))

	// privateKey := os.Getenv("PRIVATE_KEY")

	privateKey := "d51f39c1e3b3c62e918decd154e5a89d81471eb1267c5049a20"

	if privateKey == "" {
		fmt.Println("PRIVATE_KEY environment variable is not set")
		return
	}

	privateKeyBytes, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to parse private key")
		return
	}
	fmt.Println(privateKeyBytes)

	// str := "안녕하세요"
	str := "안"
	bytes := []byte(str)
	a, _ := json.Marshal(str)
	fmt.Printf("%v\n", bytes)
	fmt.Println(string(bytes))
	fmt.Println(a)

	utf16be, _ := unicode.UTF16(unicode.BigEndian, unicode.UseBOM).NewEncoder().Bytes([]byte(str))
	fmt.Println(utf16be)
	fmt.Println(string(utf16be))

	hash := common.HexToHash("0x123456789abcdef123456789abcdef123456789abcdef123456789abcdef1234")
	fmt.Println(hash)
	fmt.Println(hash.Bytes())
	// fmt.Println(hash)
	// signature, _ := crypto.Sign(hash.Bytes(), privateKeyBytes)
	// pubKey, _ := crypto.Ecrecover(hash.Bytes(), signature)

	// fmt.Printf("Public key recovered from signature: %x\n", pubKey)

}

func main() {
	ecdsa2.EC()
	// hash2.Hash()
}
