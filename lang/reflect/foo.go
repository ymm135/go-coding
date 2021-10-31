package reflect

import "fmt"


type Foo struct {
	a  int
	s  string
}

func (f Foo) Hello(s string) {
	fmt.Println("Foo Hello => ", s)
}
