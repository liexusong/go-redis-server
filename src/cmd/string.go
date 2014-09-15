// Redis server implement by Golang
// author: Jim Howard (c) liexusong at qq dot com

package cmd

import (
    "redis"
)

func SetCmd(c *Connection) bool {
    GlobalCtx.lock.Lock()

    GlobalCtx.db[c.args[1]] = args[2]
    
    GlobalCtx.lock.Unlock()

    c.SendReply("+OK")
}

func GetCmd(c *Connection) bool {

}

