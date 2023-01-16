package path

import "os"

func Path() {
	path, _ := os.Getwd()
	println(path)
}
