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
}


var RedisDb map[string]interface{} = {} // Redis's database


func ServerNew(ntype, laddr string) *Server {
    return &Server{nil, ntype, laddr}
}


func connGoFunc(conn *Connection) {
    conn.Process()
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

        conn := ConnectNew(client) // create new connection

        go connGoFunc(conn) // process connection
    }

    return nil
}
