package main

import (
	"basic"
	"fmt"
    "sorter"
    "oop"
    "strconv"
    //"concurrent"
    //"cgss/ipc"
    "web"
    //ph "photoweb" //当存在相同的子包时取别名
    "reflectm"
)

func main() {
	basic.Test()
	basic.MapTest()
	fmt.Println(basic.Add(2, 3, 4, 5))
	fmt.Println(basic.OpenFile("aa"))
    sorter.Sorter()

    oop.TestInteger()
    oop.Test()


    a,err:=strconv.ParseInt("23322222222222",10,0)
    fmt.Println(a,err)

    oop.MusicTest()

    //main返回时并不等待其它goroutine结束
    //concurrent.TestAdd()

    //concurrent.TestCount()

    //ipc.TestIpc()

    basic.RegTest()

    basic.HeapTest()

    oop.TestDuck()

    web.ExampleURL()

    //ph.Start()

    reflectm.Basic()

}

