package ipc
import "encoding/json"

type Request struct {
    Method string "method"
    Params string "params"
}

type Response struct {
    Code string "code"
    Body string "body"
}

type Server interface {
    Name() string
    Handle(method,params string) *Response
}

type IpcServer struct {
    Server
}

func NewIpcServer(server Server) *IpcServer{
    return &IpcServer{server}
}

func (server *IpcServer)Connect() chan string{
    session :=make(chan string,0)

    go func(c chan string){
        for{
            request:=<-c

            if request == "CLOSE"{
                break
            }

            var req Request
            err:=json.Unmarshal([]byte(request),&req)

        }
    }
}