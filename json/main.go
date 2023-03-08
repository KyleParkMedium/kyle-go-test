package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// JSON 데이터
	data := `
    {
        "partitionTokens": {
            "kyleToken10": [
                {
                    "amount": 0,
                    "createdDate": "2023-02-28 10:28:47",
                    "docType": "DOCTYPE_TOKEN",
                    "endDate": "2023-02-28",
                    "expiredDate": "",
                    "grade": "2",
                    "investmentPeriod": "180",
                    "isLocked": true,
                    "numberOfTokens": 0,
                    "publicOfferingAmount": 50000000,
                    "publisher": "kyle@medium.com",
                    "ror": "50",
                    "startDate": "2023-02-14",
                    "tokenHolder": "",
                    "tokenId": "kyleToken10",
                    "updatedDate": "2023-02-28 10:28:47"
                }
            ]
        }
    }
    `

	// JSON 파싱
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(data), &obj)
	if err != nil {
		panic(err)
	}

	fmt.Println(obj)

	// "partitionTokens.kyleToken10[0].investmentPeriod" 값 읽기
	kyleToken10, ok := obj["partitionTokens"].(map[string]interface{})["kyleToken10"].([]interface{})
	if !ok || len(kyleToken10) == 0 {
		panic("invalid data format")
	}
	investmentPeriod, ok := kyleToken10[0].(map[string]interface{})["investmentPeriod"].(string)
	if !ok {
		panic("invalid data format")
	}

	fmt.Println("investmentPeriod:", investmentPeriod)
}
