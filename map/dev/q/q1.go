package q

import (
	"fmt"
)

type C struct {
	DocType     string
	ID          string
	CreatedDate string
	UpdatedDate string
	Password    string
}

// func Q() {
// 	a := dev.C{ID: "hi"}

// 	fmt.Println(a)

// }

// args map[string]interface{}
func (h *C) Map() {

	fmt.Println("check")
}
