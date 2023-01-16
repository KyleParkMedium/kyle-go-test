package type1

import (
	"fmt"
	"reflect"
	"strconv"
)

func CheckType() {

	var a int64
	a = 11

	b := string(a)

	stringValue := strconv.FormatInt(a, 10)

	fmt.Println(reflect.TypeOf(b))
	fmt.Println(stringValue)
}
