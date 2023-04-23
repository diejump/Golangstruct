package main

//LV2
import "fmt"

type f interface {
	PrintType(s interface{})
	PrintValue(s interface{})
}

type m struct {
}

type DIY struct {
	name string
	age  int
}

func (r m) PrintType(s interface{}) {
	fmt.Printf("%T\n", s)
}

func (r m) PrintValue(s interface{}) {
	fmt.Printf("%v\n", s)
}
func main() {
	var s f = m{}
	diy := DIY{"xx", 18}
	s.PrintType(diy)
	s.PrintValue(diy)
}
