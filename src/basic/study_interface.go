package basic

import "fmt"

func InterfaceMain()  {
	var st *St = nil
	var it Inter = st
	fmt.Printf("%T,%p,%#V,%d\n", st, st, st, st)
	fmt.Printf("%T,%p,%#V,%d\n", it, it, it, it)
	fmt.Println(st == it, st == nil, it == nil, st, it,&it)
	if it != nil {
		//it.Ping()
		it.Pang()
	}
}


type Inter interface {
	Ping()
	Pang()
}

type St struct{}

func (St) Ping() {
	fmt.Println("ping")
}

func (*St) Pang() {
	fmt.Println("Pang")
}