package reflectm
import (
    "fmt"
    "reflect"
    "runtime/debug"
    "log"
)

func Basic(){
    var x float64 = 3.0145
    fmt.Println(reflect.TypeOf(&x))

    v:=reflect.ValueOf(x)

    fmt.Println(v.Type(),v.Float(),v.Kind())

    v.SetFloat(7.1)

    fmt.Println()



}