// Redis server implement by Golang
// author: Jim Howard (c) liexusong at qq dot com

package redis

import(
    "net"
    "cmd"
)

const (
    READ_ARG_NUMS = iota
    READ_ARG_SIZE
    READ_ARG_DATA
)


type Connection struct {
    conn net.Conn
    args []string
}


var cmdMaps[string]func([]string) bool = {
    "set": cmd.SetCmd,
    "get": cmd.GetCmd
}


// create new connection
// return Connection pointer
func ConnectNew(conn net.Conn) *Connection {
    return &Connection{conn, nil}
}


// get command args
// return true when success or false when failed
func (c *Connection) GetArgs() bool {
    state := READ_ARG_NUMS

    for {
        switch state {
        case READ_ARG_NUMS:
        case READ_ARG_SIZE:
        case READ_ARG_DATA:
        }
    }
}


// process connection
// return void
func (c *Connection) Process() {
    ok := c.GetArgs()
    if !ok {
        // todo: log error message
        return
    }

    if len(c.args) <= 0 {
        // todo: log error message
        return
    }

    cmdName := c.args[0] // command name

    callback, ok := cmdMaps[cmdName] // command callback
    if !ok {
        // todo: log error message
        return
    }

    ok := callback(c.args[1:])
    if !ok {
        // todo: log error message
    }
}
