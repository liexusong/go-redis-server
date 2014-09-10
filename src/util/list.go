// package <util>
// List implement by Golang
// author: Jim Howard (c) liexusong at qq dot com

package util

type ListNode struct {
    val interface{}
    prev *ListNode
    next *ListNode
}

type List struct {
    size int
    head *ListNode
    tail *ListNode
}


type ListError struct {
    msg string
}


func (e *ListError) Error() string {
    return e.msg
}


// create new list node
// return list node pointer
func nodeNew(val interface{}) *ListNode {
    return &ListNode{val, nil, nil}
}


// create new list object
// return list object pointer
func ListNew() *List {
    return &List{0, nil, nil}
}


// add value to list head
// return void
func (l *List) AddHead(val interface{}) {
    node := nodeNew(val)

    node.next = l.head // new node's next point to list head

    if l.head != nil {
        l.head.prev = node
    }

    l.head = node // list's head point to new node

    if l.tail == nil {
        l.tail = node
    }

    l.size++
}


// add value to list tail
// return void
func (l *List) AddTail(val interface{}) {
    node := nodeNew(val)

    node.prev = l.tail // new node's prev point to list tail

    if l.tail != nil {
        l.tail.next = node
    }

    l.tail = node // list's tail point to new node

    if l.head == nil {
        l.head = node
    }

    l.size++
}


// get list head's value and delete head node
// return value by success or error by failed
func (l *List) DelHead() (interface{}, error) {
    if l.size <= 0 {
        return nil, &ListError{"List was empty"}
    }

    ret := l.head // get list's head node

    l.head = l.head.next // list's head point to next node
    if (l.head != nil) {
        l.head.prev = nil
    } else {
        l.tail = nil
    }

    l.size--

    return ret.val, nil
}


// get list head's value
// return value by success or error by failed
func (l *List) GetHead() (interface{}, error) {
    if l.size <= 0 {
        return nil, &ListError{"List was empty"}
    }

    return l.head.val, nil
}


// get list tail's value and delete tail node
// return value by success or error by failed
func (l *List) DelTail() (interface{}, error) {
    if l.size <= 0 {
        return nil, &ListError{"List was empty"}
    }

    ret := l.tail // get list's tail node

    l.tail = l.tail.prev // list's tail point to prev node
    if (l.tail != nil) {
        l.tail.next = nil
    } else {
        l.head = nil
    }

    l.size--

    return ret.val, nil
}


// get list tail's value
// return value by success or error by failed
func (l *List) GetTail() (interface{}, error) {
    if l.size <= 0 {
        return nil, &ListError{"List was empty"}
    }

    return l.tail.val, nil
}


// get list's size
func (l *List) Size() int {
    return l.size
}

