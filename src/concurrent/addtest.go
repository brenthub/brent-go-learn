package concurrent
import "fmt"


func Add(x,y int){
    z := x+y
    fmt.Println(z)
}

func Count(ch chan int){

    ch<-1 //向对应的channel中写入一个数据在这个channel被读取前，这个操作是阻塞的。
    //通知其它协程我已经完成了
    fmt.Println("counting")
}

func TestAdd(){
    for i:=0;i<10;i++{
        go Add(i,i)
    }
}

func TestCount(){
    //使用chan(channel)在两个或多个goroutine之间传递消息
    chs :=make([]chan int,10)
    for i:=0;i<10;i++{
        chs[i]=make(chan int)
        go Count(chs[i])
    }
    for _,ch:=range(chs){
        //读取chan
        <-ch
    }
}