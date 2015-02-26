package cg
import (
    "cgss/ipc"
    "sync")

type Message struct {
    From string "from"
    To string "to"
    Content string "content"
}

type CenterServer struct {
    servers map[string] ipc.Server
    player []*Player
    rooms []*Room
    mutex sync.RWMutex
}

func NewCenterServer() *CenterServer{
    servers := make(map[string] ipc.Server)
    players := make([] *Player,0)
    return &CenterServer{servers:servers,players:servers}
}

func (server *CenterServer) AddPlayer(params string) error{
    player :=
}

