package convert

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Ex struct {
	Name string
	Age  int
}

func Convert() {
	ex := Ex{Name: "Kyle", Age: 30}

	fmt.Println(ex)

	data, err := json.Marshal(ex)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(data))
	fmt.Println(data)

	data2 := new(interface{})
	// data2 := Ex{}
	err = json.Unmarshal(data, &data2) // Convert to a map
	fmt.Println(data2)
	fmt.Println(reflect.TypeOf(data2))

	result := make(map[string]interface{})
	err = json.Unmarshal(data, &result) // Convert to a map

	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))

}

func Input(arg interface{}) {

	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(arg)
}
