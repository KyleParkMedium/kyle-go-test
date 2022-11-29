package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	dbPath, err := ioutil.TempDir("", "bookkeep")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dbPath)

}
