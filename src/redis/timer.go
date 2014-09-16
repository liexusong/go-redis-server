// Redis server implement by Golang
// author: Jim Howard (c) liexusong at qq dot com
// simple timer implement by time.Timer

package redis

import(
    "time"
)

// Types for timer

type TimerKickCb func(interface{})

type TimerKickNode struct {
    next *TimerKickNode
    freq int  // call frequency
    lrun int  // last run time
    cb TimerKickCb
    arg interface{}
}

type TimerKick struct {
    ticker *time.Ticker
    rqueue *TimerKickNode
}


func timerKickRunQueue(tk *TimerKick) {
    now := time.Now().Unix()

    // run the queue
    for node := tk.rqueue; node != nil; node = node.next {
        if int(now) - node.lrun >= node.freq {
            node.cb(node.arg)
            node.lrun = int(now)
        }
    }
}


func timerKickRutine(tk *TimerKick) {
    min := 99999999

    // find the min task
    for node := tk.rqueue; node != nil; node = node.next {
        if node.freq < min {
            min = node.freq
        }
    }

    tk.ticker = time.NewTicker(time.Duration(min) * time.Second)

    for {
        select {
        case <- tk.ticker.C:
            timerKickRunQueue(tk)
        }
    }
}


func TimerKickNew() *TimerKick {
    return &TimerKick{nil, nil}
}


func (tk *TimerKick) AddTimer(second int, cb TimerKickCb,arg interface{}) {
    node := &TimerKickNode{tk.rqueue, second, 0, cb, arg}
    tk.rqueue = node // add to queue
}


func (tk *TimerKick) Run() {
    go timerKickRutine(tk)
}

