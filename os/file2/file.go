package file2

import (
	crand "crypto/rand"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type Req struct {
	TokenWalletId string `json:"tokenWalletId"`
	Role          string `json:"role"`
	AccountNumber string `json:"accountNumber"`
}

type Mock struct {
	Data []Req
}

type Mock2 struct {
	Data map[string]string
}

func File() {

	file3, err := os.Create("distribute.json")
	if err != nil {
		panic(err)
	}
	encoder3 := json.NewEncoder(file3)

	// 랜덤 시드를 초기화합니다.
	rand.Seed(time.Now().UnixNano())

	mock := Mock{}
	mock.Data = make([]Req, 100)

	mock2 := Mock2{}
	mock2.Data = make(map[string]string, 100)

	for j := 2; j <= 101; j++ {

		m := fmt.Sprintf("wallet%v.gob", j-1)
		x := fmt.Sprintf("wallet%v.json", j-1)
		// 파일 열기
		file, err := os.Create(m)
		file2, err := os.Create(x)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// JSON 인코더 생성
		encoder := gob.NewEncoder(file)
		encoder2 := json.NewEncoder(file2)

		for i := 0; i < 100; i++ {

			req := Req{}
			req.TokenWalletId = GenerateUniqueID(j)
			req.Role = "ipo"
			// 랜덤 계좌번호를 생성하고 출력합니다.
			req.AccountNumber = generateAccountNumber()

			mock.Data[i] = req
			mock2.Data[req.TokenWalletId] = "5000"
		}

		// 구조체 파일에 쓰기
		err = encoder.Encode(mock)
		err = encoder2.Encode(mock)
		if err != nil {
			panic(err)
		}
	}

	err = encoder3.Encode(mock2)
	if err != nil {
		panic(err)
	}
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

func generateAccountNumber() string {
	// 계좌번호는 10자리 숫자로 구성됩니다.
	accountNumber := ""
	for i := 0; i < 10; i++ {
		// 0부터 9까지의 랜덤한 숫자를 생성합니다.
		digit := rand.Intn(10)
		// 숫자를 문자열로 변환하여 계좌번호에 추가합니다.
		accountNumber += fmt.Sprintf("%d", digit)
	}
	return accountNumber
}
