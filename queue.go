package gqueue

import (
    "sync"
    "container/list"
)

type Queue struct {
    lock sync.Mutex
    q []interface{}
    l *list.List
    count int
}

func NewQueue() (q *Queue) {
    q = &Queue{
        count:0,
    }

    q.l = list.New()

    return
}

func (q *Queue) Enqueue(item interface{}) {
    q.lock.Lock()
    defer q.lock.Unlock()

    q.l.PushBack(item)

    return
}

func (q *Queue) Dequeue() (item interface{}) {
    q.lock.Lock()
    defer q.lock.Unlock()

    if q.l.Len() == 0 {
        return
    }

    elem := q.l.Front()
    if elem == nil {
        return
    }

    item = q.l.Remove(elem)

    return
}

func (q *Queue) Peek() (item interface{}) {
    q.lock.Lock()
    defer q.lock.Unlock()

    if q.l.Len() == 0 {
        return
    }

    item = q.l.Front().Value
    return
}

func (q *Queue) Empty() (empty bool) {
    q.lock.Lock()
    defer q.lock.Unlock()

    empty = false
    if q.l.Len() == 0 {
        empty = true
    }
    return
}
