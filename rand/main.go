package main

import (
	crand "crypto/rand"
	"encoding/json"
	"math"
	"math/big"
	"math/rand"
	"os"
)

/*
*
crypto/rand string을 이용해 유니크한 ID 생성
*/
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type Mock struct {
	TokenId              string           `json:"tokenId"`
	PublicOfferingAmount int64            `json:"publicOfferingAmount"`
	Recipients           map[string]int64 `json:"recipients"`
}

func CreateMock() {

	// 파일 열기
	file, err := os.Create("person.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// JSON 인코더 생성
	encoder := json.NewEncoder(file)

	mock := Mock{}
	mock.TokenId = "mediumToken"
	mock.PublicOfferingAmount = 0
	mock.Recipients = make(map[string]int64)

	for i := 0; i < 10000; i++ {
		// a := GenerateUniqueID(64)
		a := ""
		for j := 0; j < i; j++ {
			a += "_"
		}

		mock.Recipients[a] = int64((rand.Intn(100) + 1) * 5000)
		mock.PublicOfferingAmount += mock.Recipients[a]
	}

	// 구조체 파일에 쓰기
	err = encoder.Encode(mock)
	if err != nil {
		panic(err)
	}

	// return mock
}

func GenerateUniqueID(length int) string {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func main() {
	CreateMock()
	// a := GenerateUniqueID(64)
	// fmt.Println(a)
}
