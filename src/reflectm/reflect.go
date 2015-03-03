package reflectm
import (
    "fmt"
    "reflect"
   // "runtime/debug"
    //"log"
)

func Basic(){
    var x float32 = 3.0145
    fmt.Println(reflect.TypeOf(x))

    p:=reflect.ValueOf(&x)//要得到x的地址

    fmt.Println(p.Type(),p.Kind(),p.CanSet())

    v:=p.Elem()//要拿到elem，直接改不行

    fmt.Println(v.CanSet())//true

    v.SetFloat(7.1)

    fmt.Println(x)

    STest()

}

type T struct {
    A int
    B string
}

type A interface {

}

func STest(){
    t:=T{202,"HW236"}
    s := reflect.ValueOf(&t).Elem()
    typeofT :=s.Type()
    for i:=0;i<s.NumField();i++{
        f := s.Field(i)
        fmt.Printf("%d:%s %s=%v\n",i,typeofT.Field(i).Name,f.Type(),f.Interface())
    }

    var at = reflect.TypeOf((*A)(nil)).Elem()
    bo:=typeofT.Implements(at)
    fmt.Println(bo)
}
