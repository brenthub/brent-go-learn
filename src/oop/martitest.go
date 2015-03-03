package oop
import (
    "fmt"
    "reflect"
    "github.com/codegangsta/inject"
)

type ClassicMartini struct {
    Router
    *Martini
}

type Router interface {
    // Get adds a route for a HTTP GET request to the specified matching pattern.
    Get(string, ...string)
}

type Martini struct {
    Injector
}

type Injector interface {
    Get(reflect.Type) reflect.Value
}

type AAA struct {

}

func (a AAA) Get(string, ...string){
    fmt.Println("aaaa")
}

func TestMartini(){
    var m *Martini = &Martini{inject.New()}
    //cn:=&ClassicMartini{m,&AAA{}}
    cn:=&ClassicMartini{&AAA{},m}
    cn.Get("aa","bbb")
}