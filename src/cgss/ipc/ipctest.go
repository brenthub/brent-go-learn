package ipc
import (
    "fmt")

type EchoServer struct {

}

func (server *EchoServer) Handle(method,param string) *Response{
    return &Response{method,method+param}
}

func (server *EchoServer) Name() string{
    return "EchoServer"
}

func TestIpc(){
    server := NewIpcServer(&EchoServer{})
    client1 :=NewIpcClient(server)
    client2 :=NewIpcClient(server)

    resp1,err1 :=client1.Call("From Client1","bb")
    resp2,err2 :=client2.Call("From Client2","aa")

    if err1==nil{
        fmt.Println(resp1)
    }
    if err2==nil{
        fmt.Println(resp2)
    }
    client2.Close()
    client1.Close()
}