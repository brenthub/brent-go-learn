package oop
import "fmt"

type Bird interface {
    Call() *Voice

}

type Duck interface {
    Bird
    Walk()
}

type Voice struct {
    msg string
}

type RealDuck struct {
    Name string
}

func (duck *RealDuck) Call() *Voice{

    return &Voice{"ga ga"}
}

func (duck *RealDuck) Walk(){
    fmt.Println(duck.Name,"RealDuck on walking!")
}


type NiDuck struct {
    Wight int
}

func (duck *NiDuck) Call() *Voice{

    return &Voice{"not call"}
}

func (duck *NiDuck) Walk(){
    fmt.Println(duck.Wight,"NiDuck on walking!")
}

func TestDuck(){
    var r Duck =&RealDuck{"Tom"}
    fmt.Println(r.Call())

    r.Walk()

    //*RealDuck前面的*只能作用于struct当为interface时会报错
    var rs *RealDuck=&RealDuck{"tin"}
    rs.Walk()

    var n Duck =&NiDuck{12}
    fmt.Println(n.Call())
    n.Walk()
}