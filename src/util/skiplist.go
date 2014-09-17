// package <util>
// SkipList implement by Golang
// author: Jim Howard (c) liexusong at qq dot com

package util

import(
    "time"
    "math/rand"
)


const MAX_LEVEL int = 32

type SkipListNode struct {
    key int
    val interface{}
    forward []*SkipListNode
}

type SkipList struct {
    header *SkipListNode
    level int
}

type SkipListError struct {
    msg string
}


// random number gen
var randGen = rand.New(rand.NewSource(time.Now().UnixNano()))

// SkipList error object
// return error message
func (err *SkipListError) Error() string {
    return err.msg
}

// private function for create SkipListNode
// return SkipListNode pointer
func skipListNodeNew(key int, val interface{}, level int) *SkipListNode {
    return &SkipListNode{key, val, make([]*SkipListNode, level)}
}


// private function for create level by random
// return random level
func randLevel() int {
    level := 1

    for level < MAX_LEVEL && randGen.Int() % 2 == 1 {
        level++
    }

    return level
}


// public function for create new SkipList
// return SkipList object pointer
func SkipListNew() *SkipList {
    return &SkipList{skipListNodeNew(0, nil, MAX_LEVEL), 0}
}


// insert key:value pair into SkipList
// return void
func (l *SkipList) Insert(key int, val interface{}) {
    var p, q *SkipListNode
    var k int

    update := [MAX_LEVEL]*SkipListNode{}

    p = l.header
    k = l.level

    // find update node
    for i := k - 1; i >= 0; i-- {
        for {
            q = p.forward[i] // next forward node
            if q == nil || q.key >= key { // find the less node than key
                break
            }
            p = q
        }
        update[i] = p
    }

    // the node exists
    if q != nil && q.key == key {
        return
    }

    k = randLevel()

    if k > l.level {
        for i := l.level; i < k; i++ {
            update[i] = l.header
        }
        l.level = k
    }

    q = skipListNodeNew(key, val, k)

    for i := 0; i < k; i++ {
        q.forward[i] = update[i].forward[i]
        update[i].forward[i] = q
    }
}


// find the value of search key
// return the value
func (l *SkipList) Find(key int) (val interface{}, err error) {
    var q *SkipListNode

    p := l.header
    k := l.level

    for i := k - 1; i >= 0; i-- {
        for {
            q = p.forward[i]

            if q != nil && q.key <= key {
                if q.key == key {
                    return q.val, nil
                }

                p = q

                continue
            }
            break
        }
    }

    return nil, &SkipListError{"Not found value by the search key"}
}


// delete the value of search key
// return error by failed or nil by success
func (l *SkipList) Delete(key int) error {
    var q *SkipListNode

    update := [MAX_LEVEL]*SkipListNode{}

    p := l.header
    k := l.level

    for i := k - 1; i >= 0; i-- {
        for {
            q = p.forward[i]
            if q != nil && q.key < key {
                p = q
                continue
            }
            break
        }

        update[i] = p
    }

    if q != nil && q.key == key {
        for i := 0; i < l.level; i++ {
            if update[i].forward[i] == q {
                update[i].forward[i] = q.forward[i]
            }
        }

        q = nil // decr refcount

        for i := l.level - 1; i >= 0; i-- {
            if l.header.forward[i] == nil {
                l.level--
            }
        }

        return nil
    }

    return &SkipListError{"Not found value by the search key"}
}


