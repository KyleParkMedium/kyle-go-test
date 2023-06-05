package main

import (
	"fmt"
	"strconv"

	"github.com/the-medium-tech/mdl-sdk-go/contract"
)

func main() {
	// address.NewAddress
	// contract.BitcoinContract

	a, err := contract.NewContract("ethereum")
	if err != nil {
		fmt.Println(err)
	}

	args := []string{"a1", "b2", "c3"}
	fmt.Println(args)

	args_byte := a.StringsToBytes(args)
	fmt.Println(args_byte)

	hash := a.Hash(args_byte)
	fmt.Println(hash)

	v := "27"
	b, _ := strconv.Atoi(v)

	hexString := strconv.FormatInt(int64(b), 16)

	c, _ := strconv.ParseUint(hexString, 16, 32)
	fmt.Println(c)
	if err != nil {
		// return errors.New("failed to convert string to uint64")
	}
	fmt.Println(hexString)
}
