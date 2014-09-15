// Redis server implement by Golang
// author: Jim Howard (c) liexusong at qq dot com

package redis

import(
    "util"
    "net"
    "sync"
)

type Server struct {
    ln net.Listener
    ntype string
    laddr string
}


// Redis context
type Context struct {
    db map[string]interface{}
    expire map[string]int
    exit bool
    lock *sync.Mutex // lock database
}


// Global context
var GlobalCtx *Context
var ctxInit bool = false


func ServerNew(ntype, laddr string) *Server {
    return &Server{nil, ntype, laddr}
}


func connGoFunc(conn *Connection) {
    conn.Process()
}


func contextInit() {
    if ctxInit {
        return
    }

    GlobalCtx = &Context{make(map[string]interface{}), make(map[string]int),
        false, new(sync.Mutex)}
}


func (s *Server) Open() error {
    contextInit() // init global context

    ln, err = net.Listen(s.ntype, s.laddr)
    
    if err != nil {
        return err
    }

    s.ln = ln // save listener object

    for GlobalCtx.exit == false {
        client, err := ln.Accept() // accept new client
        if err != nil {
            continue
        }

        conn := ConnectNew(client) // create new connection

        go connGoFunc(conn) // process connection
    }

    return nil
}
