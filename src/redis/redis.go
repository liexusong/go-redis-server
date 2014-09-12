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


// Redis context
type Context struct {
    db map[string]interface{}
    exit bool
    ch chan int
}


var RedisContext *Context


func ServerNew(ntype, laddr string) *Server {
    return &Server{nil, ntype, laddr}
}


func connGoFunc(conn *Connection) {
    conn.Process()
}


func contextInit() {
    RedisContext = Context{make(map[string]interface), false, make(chan int)}
}


func (s *Server) Open() error {
    s.ln, err = net.Listen(s.ntype, s.laddr)
    
    if err != nil {
        return err
    }

    for RedisContext.exit == false {
        client, err := ln.Accept() // accept new client
        if err != nil {
            continue
        }

        conn := ConnectNew(client) // create new connection

        go connGoFunc(conn) // process connection
    }

    return nil
}
