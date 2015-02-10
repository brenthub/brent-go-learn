package oop
import "fmt"

type Integer int

type LessAdder interface {
    Less(b Integer) bool
}

func (a Integer) Less(b Integer) bool {
	return a < b
}

func TestInteger(){
    var a Integer =1

    var b LessAdder = &a

    fmt.Println(b.Less(2))

}

type Base struct {
    Name string
}

func (base *Base) Foo(){

}

func (base *Base) Bar(){
    fmt.Println("Base Bar")
}

type Foo struct {
    Base
    Age int
}

func (foo *Foo) Bar(){
    foo.Base.Bar()
    fmt.Println("foo Bar")
}



func Test(){
    var f *Foo =new(Foo)
    f.Name="1021"
    f.Age=12
    f.Bar()
    fmt.Println(f)
}