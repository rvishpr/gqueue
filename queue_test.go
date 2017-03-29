package gqueue

import (
	"testing"
)

func TestQueueAddDelete(t *testing.T) {
	q := NewQueue()

	if q.Empty() != true {
		t.Fatal("test failed")
	}

	q.Enqueue("test")
	if q.l.Len() != 1 {
		t.Fatal("test failed")
	}

	item := q.Peek()
	if item.(string) != "test"{
		t.Fatal("test failed")
	}

	item = q.Dequeue()
	if item.(string) != "test" {
		t.Fatal("test failed")
	}
	if q.l.Len() != 0 {
		t.Fatal("test failed")
	}
	if q.Empty() != true {
		t.Fatal("test failed")
	}
}

func TestQueueMultipleAddDelete(t *testing.T) {
	q := NewQueue()

	if q.Empty() != true {
		t.Fatal("test failed")
	}

	q.Enqueue("test1")
	if q.l.Len() != 1 {
		t.Fatal("test failed")
	}

	q.Enqueue("test2")
	if q.l.Len() != 2 {
		t.Fatal("test failed")
	}

	item := q.Peek()
	if item.(string) != "test1" {
		t.Fatal("test failed")
	}

	q.Enqueue("test3")
	if q.l.Len() != 3 {
		t.Fatal("test failed")
	}

	item = q.Peek()
	if item.(string) != "test1" {
		t.Fatal("test failed")
	}

	item = q.Dequeue()
	if item.(string) != "test1" {
		t.Fatal("test failed")
	}
	if q.l.Len() != 2 {
		t.Fatal("test failed")
	}

	item = q.Dequeue()
	if item.(string) != "test2" {
		t.Fatal("test failed")
	}
	if q.l.Len() != 1 {
		t.Fatal("test failed")
	}

	item = q.Dequeue()
	if item.(string) != "test3" {
		t.Fatal("test failed")
	}
	if q.l.Len() != 0 {
		t.Fatal("test failed")
	}
	if q.Empty() != true {
		t.Fatal("test failed")
	}
}

func assert(t testing.TB, cond bool) {
	if !cond {
		t.Fatal("test failed")
	}
}