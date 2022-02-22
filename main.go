package main

import (
	"fmt"
	p "pck/pck1"
)

func main() {
	p1 := new(p.People)
	p1.Cinsiyet = "Kadın"
	p1.Name = "Ayşe"
	p1.SetTckn(12341232132)
	p1.SetYas(24)
	p1.SetYas(25)
	fmt.Println(*p1)
}
