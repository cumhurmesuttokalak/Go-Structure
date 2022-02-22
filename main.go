package main

import (
	"fmt"
	p "pck/pck1"
)

func main() {
	p1 := p.People{Name: "Mesut"}
	p1.SetTckn(123456)
	p1.SetYas(34)
	fmt.Println(p1)
	fmt.Println(p1.GetTckn(), p1.GetYas())
}
