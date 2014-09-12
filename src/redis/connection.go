// Redis server implement by Golang
// author: Jim Howard (c) liexusong at qq dot com

package redis

import(
    "net"
)


type Connection struct {
    conn net.Conn
    buff []byte
}


func ConnectNew(conn net.Conn) {
    return &Connection{conn, make([]byte, 1024)}
}


func (c *Connection) Process() {
    nbytes, err := c.conn.Read(c.buff) // read buffer from connection
    if err != nil {
        c.conn.Close()
        return
    }

    //todo: parse command
}
