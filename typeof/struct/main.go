package main

import (
	dev "struct/dev"
)

type Partition struct {
	value   int
	address string
}

// type Partition interface {
// 	value int
// }

// type A interface {

// }

type wallet struct {
	name  string
	array []Partition
}

type w2 struct {
	name string
	mp   map[string]Partition
}

type A struct {
	Amount int
}

type AA struct {
	Name            string
	PartitionTokens map[string][]A `json:"tokens"`
}

func main() {

	// qqq := &dev.A{5}
	// fmt.Println(qqq)
	// x := map[string][]A{}
	// b := make([]A, 2)

	// b[0].Amount = 50
	// b[1].Amount = 100

	// // tq := map[string]A{}

	// x["QQ"] = b
	// x["YY"] = b

	// c := AA{"hi", x}

	// data, err := json.Marshal(c) // Convert to a json string
	// fmt.Println(data)

	// d := reflect.TypeOf(c).Kind()
	// fmt.Println(d)
	// fmt.Println(c)

	// e := c.PartitionTokens

	// fmt.Println(e)
	// f := reflect.TypeOf(e).Kind()
	// fmt.Println(f)

	// g, err := json.Marshal(e)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(g)

	// // qwer := A{}

	// // // // docType 확인
	// // resultMap := map[string]AA.A{}

	// // err = json.Unmarshal(g, &resultMap)

	// // // err = json.Unmarshal(g, &e)
	// // fmt.Println(resultMap)

	// // for a, b := range c {
	// // 	fmt.Println(a)
	// // 	fmt.Println(a)
	// // }

	// // a := AA{"hi"}

	// // x := map[string][]Partition{}
	// // // b := make([]Partition, 1)
	// // // b[0].value = 10
	// // // b[0].address = "A"

	// // b := make([]Partition, 1)
	// // b[0].value = 10
	// // b[0].address = "A"
	// // x["Q"] = append(b)

	// // c := make([]Partition, 1)
	// // c[0].value = 50
	// // c[0].address = "B"
	// // x["W"] = append(c)

	// // fmt.Println(x["Q"])

	// // // q := make(map[string]Partition)
	// // // q["Q"] = Partition{10, "W"}
	// // // fmt.Println(q)

	// qwe := map[string]Partition{}
	// qwe["A"] = Partition{5, "AA"}

	// // abc := map["AA"]Partition{value: 5, address: "AA"}
	// w := w2{"DEV", qwe}

	// fmt.Println(w)

	// // // e := make(map[string]Partition)
	// // // e["Z"] = Partition{100, "XX"}

	// // // w.mp = e
	// // // fmt.Println(w)
	// // // // b := make([]Partition, 1)
	// // // // b[0].value = 10
	// // // // b[0].address = "A"

	// // // // // c := make([]Partition, 1)
	// // // // // c[0].value = 50
	// // // // // c[0].address = "B"
	// // // // c := Partition{50, "B"}

	// // // // a := wallet{"park", b}

	// // // // a.array = append(a.array, c)

	// // // // fmt.Println(b)
	// // // // fmt.Println(a)

	// // // // // fmt.Println(a.array["B"])

	// dev.AA()
	dev.Basic()
}
