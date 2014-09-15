// Redis server implement by Golang
// author: Jim Howard (c) liexusong at qq dot com

package redis

import(
    "net"
    "cmd"
)

const (
    READ_ARG_SIZE = iota
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
    rbuf := make([]byte, 0, 64)
    args := -1
    last, curr := 0, 0

    // read arg numbers
    for {
        nbytes, err := c.conn.Read(rbuf[last:])
        if err != nil {
            return false
        }

        last += nbytes

        if last <= 0 {
            continue
        }

        if rbuf[0] != '*' { // invaild redis protocol
            return false
        }

        for cnt, i := 0, 1; i < last; i++ {
            if rbuf[i] >= '0' && rbuf[i] <= '9' {
                cnt = cnt * 10 + (rbuf[i] - '0')

            } else { // finish arg numbers string buffer
                args = cnt
                curr = i
                break
            }
        }

        // enough CRLF string
        if last >= curr + 2 {
            if rbuf[curr] == '\r' && rbuf[curr + 1] == '\n' {
                curr += 2 // point to real data offset
                break
            } else {
                return false
            }
        }
    }

    // begin read args
    nbuf := make([]byte, 0, 64)
    if last > curr + 1 {
        copy(nbuf[0:], rbuf[curr:])
        last = last - curr
    }

    c.args = make([]string, 0, args)
    state := READ_ARG_SIZE
    nsize := -1

    for {
        switch state {
        case READ_ARG_SIZE:
            nbytes, err := c.conn.Read(nbuf[last:])
            if err != nil {
                return false
            }

            last += nbytes

            if last <= 0 {
                continue
            }

            if nbuf[0] != '$' { // invaild redis protocol
                return false
            }

            for cnt, i := 0, 1; i < last; i++ {
                if nbuf[i] >= '0' && nbuf[i] <= '9' {
                    cnt = cnt * 10 + (nbuf[i] - '0')
                } else {
                    nsize = cnt
                    curr = i
                    break
                }
            }

            // enough CRLF string
            if last >= curr + 2 {
                if rbuf[curr] == '\r' && rbuf[curr + 1] == '\n' {
                    curr += 2 // point to real data offset
                    state = READ_ARG_DATA
                } else {
                    return false
                }
            }

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
