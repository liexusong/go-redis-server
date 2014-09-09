// SkipList implement by Golang
// author: Liexusong

package unit

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

var randGen := rand.New(rand.NewSource(time.Now().UnixNano()))


// private function for create SkipListNode
// return SkipListNode pointer
func nodeNew(key int, val interface{}, level int) *SkipListNode {
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
    return &SkipList{nodeNew(0, nil, MAX_LEVEL), 0}
}


// insert key:value pair into SkipList
// return void
func (l *SkipList) Insert(key int, val interface{}) {
    var p, q *SkipListNode
    var k int

    update := [MAX_LEVEL]*SkipListNode

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

    q = nodeNew(key, val, k)

    for i := 0; i < k; i++ {
        q.forward[i] = update[i].forward[i]
        update[i].forward[i] = q
    }
}

