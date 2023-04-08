package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"runtime"
	"time"

	"github.com/ethereum/go-ethereum/rlp"
)

var memStatsBefore, memStatsAfter, memStatsBefore2, memStatsAfter2 runtime.MemStats

func main() {

	// 10,000 자리의 랜덤한 정수를 생성합니다.
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	bigIntVal := big.NewInt(0).Rand(
		random,
		big.NewInt(0).Exp(big.NewInt(10), big.NewInt(10000), nil),
	)

	fmt.Println(reflect.TypeOf(bigIntVal))

	runtime.ReadMemStats(&memStatsBefore) // 현재 메모리 사용량을 측정합니다.
	_, err := rlp.EncodeToBytes([]interface{}{bigIntVal})
	if err != nil {
		fmt.Println(err)
	}
	runtime.ReadMemStats(&memStatsAfter) // 현재 메모리 사용량을 다시 측정합니다.

	// // JSON 객체를 출력합니다.
	// fmt.Printf("Marshalled JSON: %s\n", string(jsonVal))
	// 메모리 사용량을 출력합니다.
	memoryUsage := int(memStatsAfter.TotalAlloc - memStatsBefore.TotalAlloc)
	fmt.Printf("Memory usage: %d bytes\n", memoryUsage)

	byte := []byte(bigIntVal.String())
	runtime.ReadMemStats(&memStatsBefore) // 현재 메모리 사용량을 측정합니다.
	_, err = rlp.EncodeToBytes([]interface{}{byte})

	if err != nil {
		fmt.Println(err)
	}
	runtime.ReadMemStats(&memStatsAfter) // 현재 메모리 사용량을 다시 측정합니다.

	// // JSON 객체를 출력합니다.
	// fmt.Printf("Marshalled JSON: %s\n", string(jsonVal))
	// 메모리 사용량을 출력합니다.
	memoryUsage = int(memStatsAfter.TotalAlloc - memStatsBefore.TotalAlloc)
	fmt.Printf("Memory usage: %d bytes\n", memoryUsage)

	// Big.Int
	// JSON 객체에 포함될 데이터를 맵으로 생성합니다.
	data := map[string]interface{}{
		"bigIntVal": bigIntVal,
	}

	runtime.ReadMemStats(&memStatsBefore) // 현재 메모리 사용량을 측정합니다.
	// JSON 객체를 생성합니다.
	_, err = json.Marshal(data)
	if err != nil {
		panic(err)
	}
	runtime.ReadMemStats(&memStatsAfter) // 현재 메모리 사용량을 다시 측정합니다.

	// // JSON 객체를 출력합니다.
	// fmt.Printf("Marshalled JSON: %s\n", string(jsonVal))
	// 메모리 사용량을 출력합니다.
	memoryUsage = int(memStatsAfter.TotalAlloc - memStatsBefore.TotalAlloc)
	fmt.Printf("Memory usage: %d bytes\n", memoryUsage)

	// String
	// JSON 객체에 포함될 데이터를 맵으로 생성합니다.
	data2 := map[string]interface{}{
		"bigIntVal": bigIntVal.String(),
	}

	runtime.ReadMemStats(&memStatsBefore2) // 현재 메모리 사용량을 측정합니다.
	// JSON 객체를 생성합니다.
	_, err = json.Marshal(data2)
	if err != nil {
		panic(err)
	}
	runtime.ReadMemStats(&memStatsAfter2) // 현재 메모리 사용량을 다시 측정합니다.

	// // JSON 객체를 출력합니다.
	// fmt.Printf("Marshalled JSON: %s\n", string(jsonVal2))
	// 메모리 사용량을 출력합니다.
	memoryUsage = int(memStatsAfter2.TotalAlloc - memStatsBefore2.TotalAlloc)
	fmt.Printf("Memory usage: %d bytes\n", memoryUsage)
}
