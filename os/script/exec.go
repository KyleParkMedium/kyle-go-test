package script

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Cmd struct {
	// path of a program or executable
	Path string
	// arguments to invoke program or executable with
	Args []string
	// environment variables for the execution
	Env []string
	// current working directory of the execution
	Dir string
	// standard input for the execution
	Stdin io.Reader
	// standard output for the execution
	Stdout io.Writer
	// standard error for the execution
	Stderr io.Writer
	// Process is the underlying process, once started
	Process *os.Process
}

func Exec() {

	goExecPath, err := exec.LookPath("go")

	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("Go :", goExecPath)
	}

	cmdGoVer := &exec.Cmd{
		Path:   goExecPath,
		Args:   []string{goExecPath, "version"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	// fmt.Println(cmdGoVer.String())

	if err := cmdGoVer.Run(); err != nil {
		fmt.Println(err)
	}

	cmd := &exec.Cmd{
		Path:   "/Users/park/test/kyle-go-test/os/script/order.sh",
		Args:   []string{goExecPath, "version"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
