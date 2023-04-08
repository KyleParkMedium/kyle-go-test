package hash2

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/sha3"
)

func Hash() {
	// 입력할 데이터
	data := []byte("Hello, world!")

	// SHA-3-256 해시 함수 생성
	hash_func := sha3.New256()
	// 데이터 업데이트
	hash_func.Write(data)
	// 해시 값 출력
	hash_value := fmt.Sprintf("%x", hash_func.Sum(nil))
	fmt.Println(hash_value)

	// SHA-256 해시 함수 생성
	hash_func = sha256.New()
	// 데이터 업데이트
	hash_func.Write(data)
	// 해시 값 출력
	hash_value = fmt.Sprintf("%x", hash_func.Sum(nil))
	fmt.Println(hash_value)

	// Keccak-256 해시 함수 생성
	hash_func = sha3.NewLegacyKeccak256()
	// 데이터 업데이트
	hash_func.Write(data)
	// 해시 값 출력
	hash_value = fmt.Sprintf("%x", hash_func.Sum(nil))
	fmt.Println(hash_value)
	fmt.Println(hash_func.Sum(nil))
	fmt.Println(len(hash_func.Sum(nil)))

	a := "안녕하세요"
	b, _ := json.Marshal(a)
	fmt.Println(b)
	fmt.Println(len(b))

}
