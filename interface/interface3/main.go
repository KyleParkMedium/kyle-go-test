package main

import "fmt"

type IGun interface {
	// setName(name string)
	// setPower(power int)
	// getName() string
	// getPower`() int
}

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

type Ak47 struct {
	Gun
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

func main() {
	ak47, _ := getGun("ak47")

	fmt.Println(ak47)
}
