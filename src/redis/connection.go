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


var cmdMaps[string]func(c *Connection) bool = {
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
    rbuf := make([]byte, 64)
    args := -1
    last, curr := 0, 0

    // read arg numbers
    for {
        nbytes, err := c.conn.Read(rbuf[last:])
        if err != nil {
            fmt.Println(err)
            return false
        }

        last += nbytes

        if last <= 0 {
            continue
        }

        if rbuf[0] != '*' { // invaild redis protocol
            fmt.Println("Invaild redis protocol")
            return false
        }

        for cnt, i := 0, 1; i < last; i++ {
            if rbuf[i] >= '0' && rbuf[i] <= '9' {
                cnt = cnt * 10 + int(rbuf[i] - '0')

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
                fmt.Println("Invaild redis protocol")
                return false // invaild redis protocol
            }
        }
    }

    // read args data

    if last > curr + 1 {
        copy(rbuf[0:], rbuf[curr:]) // reset read buffer
        last = last - curr
    }

    c.args = make([]string, args) // create args array
    state := READ_ARG_SIZE

    var psize int
    var pcurr int
    var pdata []byte
    var pinit bool

    for {
        switch state {
        case READ_ARG_SIZE:
            nbytes, err := c.conn.Read(rbuf[last:])
            if err != nil {
                fmt.Println(err)
                return false
            }

            last += nbytes

            if last <= 0 {
                continue
            }

            if rbuf[0] != '$' { // invaild redis protocol
                fmt.Println("Invaild redis protocol")
                return false
            }

            for cnt, i := 0, 1; i < last; i++ {
                if rbuf[i] >= '0' && rbuf[i] <= '9' {
                    cnt = cnt * 10 + int(rbuf[i] - '0')
                } else {
                    psize = cnt
                    curr = i
                    break
                }
            }

            // enough CRLF string
            if last >= curr + 2 {
                if rbuf[curr] == '\r' && rbuf[curr + 1] == '\n' {
                    curr += 2 // point to real data offset
                    state = READ_ARG_DATA
                    pinit = false
                } else {
                    fmt.Println("Invaild redis protocol")
                    return false
                }
            }

        case READ_ARG_DATA:
            if pinit == false {
                pdata = make([]byte, psize)
                pcurr = 0
                pinit = true
            }

            nbytes, err := c.conn.Read(pdata[pcurr:])
            if err != nil {
                fmt.Println(err)
                return false
            }

            pcurr += nbytes
            if pcurr >= psize {
                c.args.append(string(pdata)) // append to connection's args
                state = READ_ARG_SIZE
            }
        }
    }
}


// send data to client connection
// return true when success or false when failed
func (c *Connection) SendReply(msg string) bool {
    wbuf := ([]byte)(msg) // change string to byte slice
    last, total := 0, len(wbuf)

    for {
        nbytes, err := c.conn.Write(wbuf[last:])
        if err != nil {
            fmt.Println(err)
            return false
        }

        last += nbytes

        if last >= total {
            break
        }
    }

    return true
}


// process connection
// return void
func (c *Connection) Process() {
again:
    ok := c.GetArgs()
    if !ok {
        // todo: log error message
        return
    }

    if len(c.args) <= 0 {
        // todo: log error message
        return
    }

    cmdName := c.args[0]

    callback, ok := cmdMaps[cmdName] // command callback
    if !ok {
        fmt.Println("Not found command `%s'\n", cmdName)
        return
    }

    ok := callback(c)
    if !ok {
        // todo: log error message
    }

    goto again // process again
}

