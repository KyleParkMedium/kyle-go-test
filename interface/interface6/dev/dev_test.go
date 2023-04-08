package dev

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestA(t *testing.T) {

	set := big.NewInt(0)
	_, ok := set.SetString("15162319880118241914981489149891848914891894", 10)
	if !ok {
		fmt.Println("Failed to set big integer value")
		return
	}
	// s := set.String()

	set2 := big.NewFloat(11.1)

	data := map[string]interface{}{
		// "name": string("John"),
		"age":  set,
		"age2": set2,
		// "age2": 11.1,
		// "age":  big.NewInt(111111111454455451),
		// "age":  int64(1111111111111111111),
		// "ageString": s,
	}

	fmt.Println(reflect.TypeOf(data["age2"]))

	jsbytes, _ := json.Marshal(data)

	c := make(map[string]interface{})
	_ = json.Unmarshal(jsbytes, &c)

}
