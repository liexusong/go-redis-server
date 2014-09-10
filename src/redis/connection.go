// Redis server implement by Golang
// author: Jim Howard (c) liexusong at qq dot com

package redis

import(
    "net"
)


type Connection struct {
    conn net.Conn
}


func RedisConnectionNew(conn net.Conn) {
    return &Connection{conn}
}


func (c *Connection) MainLoop() {
    for {
        c.handler(c)
    }
}
