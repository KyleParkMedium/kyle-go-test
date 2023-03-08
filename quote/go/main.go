package main

import (
	"fmt"
	"strconv"
)

func main() {

	// a := "hi"
	// b := "\"hi\""

	// c := strconv.Quote(a)

	// d, err := strconv.ParseBool(a)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(d)

	// e, err := strconv.Atoi(a)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(e)

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// b10 := []byte("int (base 10):")
	// b10 = strconv.AppendInt(b10, -42, 10)
	// fmt.Println(string(b10))

	// b2 := []byte("int (base 2):")
	// b2 = strconv.AppendInt(b2, -42, 2)
	// fmt.Println(string(b2))

	// b16 := []byte("int (base 16):")
	// b16 = strconv.AppendInt(b16, -42, 16)
	// fmt.Println(string(b16))

	// b := []byte("quote:")
	// b = strconv.AppendQuote(b, `"Fran & Freddie's Diner"`)
	// fmt.Println(string(b))

	qu := `"a"`
	qb := `'a'`
	qc := "`aa`"
	fmt.Println(qu)
	fmt.Println(qb)
	fmt.Println(qc)

	b := []byte("quote:")
	b = strconv.AppendQuote(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))

}
