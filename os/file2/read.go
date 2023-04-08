package file2

import (
	"encoding/gob"
	"fmt"
	"os"
)

func Read() {

	file, err := os.Open("/Users/park/test/kyle-go-test/os/wallet.gob")
	if err != nil {
		fmt.Println("file open error:", err)
		return
	}
	defer file.Close()

	// 파일에서 디코딩할 값을 가져옵니다.
	var people Mock
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&people)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	for _, person := range people.Data {

		fmt.Println(person)
	}
}
