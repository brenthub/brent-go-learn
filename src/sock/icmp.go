package sock
import (
    //"net"
    "fmt"
    "os"
)
//Go语言标准库对此过程进行了抽象和封装。
//无论我们期望使用什么协议建立什么形式的连接，都只需要调用net.Dial()即可
func TestIcmp(){
//    host:="127.0.0.1:2633"
//    conn,err:=net.Dial("ip4:icmp",host)
//    checkError(err)
    var msg [512]byte
    msg[0]=8 //echo
    msg[1]=0 //code 0
    msg[2]=0 //checksum
    msg[3]=0 //checksum
    msg[4]=0 //identifier[0]
    msg[5]=13 //identifier[1]
    msg[6]=0 //sequence[0]
    msg[7]=37 //sequence[1]

}

func checkError(err error){
    if err!=nil{
        fmt.Fprint(os.Stderr,"Fatal error:%s",err.Error())
        os.Exit(1)
    }
}
