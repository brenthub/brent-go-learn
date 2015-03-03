package martini
import (
    "github.com/codegangsta/martini"
)

func Start(){

    var m *martini.ClassicMartini = martini.Classic()

    m.Get("/",func() string{ //this code can run
        return "Hello,world!"
    })

    m.RunOnAddr(":3001")
}