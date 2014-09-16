// Redis server implement by Golang
// author: Jim Howard (c) liexusong at qq dot com
// simple timer generator implement by time.Ticker

package redis

import(
    "time"
)

// Types for timer

type TimerGenCb func(interface{})

type TimerGenNode struct {
    next *TimerGenNode
    freq int  // call frequency
    lrun int  // last run time
    cb TimerGenCb
    arg interface{}
}

type TimerGen struct {
    ticker *time.Ticker
    rqueue *TimerGenNode
}


func timerGenRunQueue(t *TimerGen) {
    now := time.Now().Unix()

    // run the queue
    for node := t.rqueue; node != nil; node = node.next {
        if int(now) - node.lrun >= node.freq {
            node.cb(node.arg)
            node.lrun = int(now)
        }
    }
}


// private function for timer
// using time.Ticker to generate signal
func timerGenRutine(t *TimerGen) {
    min := 99999999

    // find the min task
    for node := t.rqueue; node != nil; node = node.next {
        if node.freq < min {
            min = node.freq
        }
    }

    t.ticker = time.NewTicker(time.Duration(min) * time.Second)

    for {
        select {
        case <- t.ticker.C:
            timerGenRunQueue(t)
        }
    }
}


// create new timer generator
// return TimerGen object
func TimerGenNew() *TimerGen {
    return &TimerGen{nil, nil}
}


// add timer to timer generator
// return void
func (t *TimerGen) AddTimer(second int, cb TimerGenCb,arg interface{}) {
    node := &TimerGenNode{t.rqueue, second, 0, cb, arg}
    t.rqueue = node // add to queue
}


// run all timer
// async execute by goroutine
func (t *TimerGen) Run() {
    go timerGenRutine(t)
}

