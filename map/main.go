package main

import (
	"fmt"
	"map/dev"
)

type C struct {
	DocType     string
	ID          string
	CreatedDate string
	UpdatedDate string
	Password    string
}

func A(d map[string]interface{}) {
	fmt.Println(d)
	// map[string]interface{}
}

func main() {
	// // b := C{
	// // 	DocType:     "",
	// // 	ID:          "admin",
	// // 	CreatedDate: "adminpw",
	// // 	UpdatedDate: "",
	// // 	Password:    "",
	// // }

	// d := make(map[string]interface{})
	// d["qq"] = "QQ"
	// d["ww"] = "WW"
	// d["ee"] = C{
	// 	DocType:     "",
	// 	ID:          "admin",
	// 	CreatedDate: "adminpw",
	// 	UpdatedDate: "",
	// 	Password:    "",
	// }

	// e := make(map[string]string)
	// e["qq"] = "QQ"
	// e["ww"] = "WW"
	// // d := {id:"a", password:"b"}

	// fmt.Println(d)
	// fmt.Println(e)
	// A(d)

	// // abc := C{
	// // 	DocType:     "",
	// // 	ID:          "admin",
	// // 	CreatedDate: "adminpw",
	// // 	UpdatedDate: "",
	// // 	Password:    "",
	// // }
	// // dev.Parse(abc)

	// // abc := C{
	// // 	DocType:     "",
	// // 	ID:          "admin",
	// // 	CreatedDate: "adminpw",
	// // 	UpdatedDate: "",
	// // 	Password:    "",
	// // }
	// ee := dev.Parse()

	// fmt.Println(ee)

	// q.Q()
	// dev.W()
	dev.E()
}
