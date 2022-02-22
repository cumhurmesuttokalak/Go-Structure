package pck1

import "fmt"

type People struct {
	tckn int
	Name string
	yas  int
}

func (p People) GetTckn() int {
	return p.tckn
}
func (p *People) SetTckn(yeniNum int) {
	p.tckn = yeniNum
}
func (p People) GetYas() int {
	return p.yas
}
func (p *People) SetYas(yeniYas int) {
	if yeniYas > 0 {
		p.yas = yeniYas
	} else {
		fmt.Println("Yaş 0'dan kücük olamaz")
	}

}
