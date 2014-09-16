// Redis server implement by Golang
// author: Jim Howard (c) liexusong at qq dot com

package redis

import(
    "time"
    "util"
)

type TimerKick struct {
    timer *time.Timer
    rqueue *util.list
}

type TimerKickCb func(interface{})

type TimerKickNode struct {
    freq int  // call frequency
    lrun int  // last run time
    cb TimerKickCb
    arg interface{}
}


func timerKickRunQueue(tk *TimerKick) {
    nowts := time.Now().Unix()

    // run the queue
    for node := tk.rqueue.head; node != nil; node = node.next {
        ent := (*TimerKickNode)(node.val)
        if int(nowts) - ent.lrun >= freq {
            ent.cb(ent.arg)
            ent.lrun = nowts
        }
    }
}


func timerKickRutine(tk *TimerKick) {
    min := 99999999

    for node := tk.rqueue.head; node != nil; node = node.next {
        ent := (*TimerKickNode)(node.val)
        if ent.freq < min {
            min = ent.freq
        }
    }

    tk.timer = time.NewTimer(min * time.Second)

    for {
        select {
        case <- tk.timer.C
            timerKickRunQueue(tk)
        }
    }
}


func TimerKickNew() *TimerKick {
    return &TimerKick{nil, util.ListNew()}
}


func (tk *TimerKick) AddTimer(second int, cb TimerKickCb, arg interface{}) {
    node := TimerKickNode{second, 0, cb, arg}
    tk.rqueue.AddTail(node)
}


func (tk *TimerKick) Run() {
    go timerKickRutine(tk)
}
