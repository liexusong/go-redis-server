// Redis server implement by Golang
// author: Jim Howard (c) liexusong at qq dot com

package redis

import(
    "util"
    "net"
)

type Server struct {
    ln net.Listener
    ntype string
    laddr string
    db map[string] interface{}
}


func RedisServerNew(ntype, laddr string) *Server {
    return &Server{nil, ntype, laddr}
}


func redisConnectionLoop(conn *Connection) {
    conn.MainLoop()
}


func (s *Server) Open() error {
    s.ln, err = net.Listen(s.ntype, s.laddr)
    
    if err != nil {
        return err
    }

    for {
        client, err := ln.Accept()
        if err != nil {
            continue
        }

        conn := RedisConnectionNew(client)

        go redisConnectionLoop(conn)
    }

    return nil
}

