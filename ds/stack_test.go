package ds

import "testing"

func TestNewStack(t *testing.T) {
	s := NewStack()
	if s.Head != nil || s.size != 0 {
		t.Errorf("Head, Expected: %v, Got: %v. size, Expected: %v, Got: %v.", nil, s.Head, 0, s.size)
		t.Fail()
	}
}

func TestPush(t *testing.T) {
	ts := NewStack()
	pl := ts.size
	var data uint8 = 1
	ts.Push(data)
	cl := ts.size
	if pl != cl-1 {
		t.Errorf("size, Expected: %v, Got: %v.", cl-1, pl)
		t.Fail()
	}
	if ts.Head.data != data {
		t.Errorf("Head, Expected: %v, Got: %v.", data, ts.Head.data)
		t.Fail()
	}
}

func dummyStack() Stack {
	return Stack{
		Head: &Node{
			data: 1,
			next: &Node{
				data: 2,
				next: nil,
			},
		},
		size: 2,
	}
}

func TestPopOnEmpty(t *testing.T) {
	ts := NewStack()
	if d := ts.Pop(); d != nil || ts.size != 0 {
		t.Errorf("data, Expected: %v, Got: %v. size, Expected: %v, Got: %v.", nil, d, 0, ts.size)
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	ts := dummyStack()
	if data := ts.Pop(); *data != 1 || ts.Head.data != 2 || ts.size != 1 {
		t.Errorf("data, Expected: %v, Got: %v. Head, Expected: %v, Got: %v. size, Expected: %v, Got: %v.", 1, data, 2, ts.Head.data, 1, ts.size)
		t.Fail()
	}
}

func TestSize(t *testing.T) {
	ts := dummyStack()
	if ts.Size() != ts.size {
		t.Errorf("size, Expected: %d, Got: %d.", ts.size, ts.Size())
		t.Fail()
	}
}
