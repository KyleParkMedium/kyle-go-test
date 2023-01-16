package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	//msg *Request
	cmd := &exec.Cmd{
		// Path: "./bash/execChaincode.sh",
		// Path: "/Users/park/test/sto-token-api/bash/execChaincode.sh",
		Path: "/Users/park/code/purefabric/execChaincode.sh",
		// Path: "./execChaincode.sh",

		// Args:   req.Args,
		// Path:   "./bash/a.sh",

		// for

		Args:   []string{"", "-F", "A"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	// check logic

}
