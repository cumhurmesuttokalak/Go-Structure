package pck

import (
	"fmt"
	"strconv"
)

type People struct {
	tckn     int
	Name     string
	yas      int
	Cinsiyet string
}

func (p People) GetTckn() int {
	return p.tckn
}
func (p *People) SetTckn(YeniTckn int) {
	if len(strconv.Itoa(YeniTckn)) != 11 {
		fmt.Println("Tckn 11 rakam olmalıdır")
	} else {
		p.tckn = YeniTckn
	}
}
func (p People) GetYas() int {
	return p.yas
}
func (p *People) SetYas(YeniYas int) {
	if YeniYas > 0 {
		p.yas = YeniYas
	} else {
		fmt.Println("Yas 0'dan kücük olamaz")
	}
}
