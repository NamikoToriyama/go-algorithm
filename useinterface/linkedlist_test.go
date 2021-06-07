package useinterface

import (
	"errors"
	"testing"
)

func checkNodeValue(got *Node, want *Node) error {
	for got.next != nil {
		if got.value != want.value {
			return errors.New("Does not same value")
		}
		got = got.next
		want = want.next
	}
	return nil
}

func TestRemoveNode(t *testing.T) {
	node := &Node{1, &Node{2, &Node{3, &Node{4, &Node{0, nil}}}}}
	want := &Node{2, &Node{3, &Node{4, &Node{0, nil}}}}
	if got := node.RemoveNode(0); checkNodeValue(got, want) != nil {
		t.Fatalf("want = %v, got = %v", want, got)
	}

	node = &Node{1, &Node{2, &Node{3, &Node{4, &Node{0, nil}}}}}
	want = &Node{1, &Node{3, &Node{4, &Node{0, nil}}}}
	if got := node.RemoveNode(1); checkNodeValue(got, want) != nil {
		t.Fatalf("want = %v, got = %v", want, got)
	}

	node = &Node{1, &Node{2, &Node{3, &Node{4, &Node{0, nil}}}}}
	want = node
	if got := node.RemoveNode(4); checkNodeValue(got, want) != nil {
		t.Fatalf("want = %v, got = %v", want, got)
	}
}

func TestInsertNode(t *testing.T) {
	node := &Node{1, &Node{2, &Node{3, &Node{4, &Node{0, nil}}}}}
	want := &Node{0, &Node{1, &Node{2, &Node{3, &Node{4, &Node{0, nil}}}}}}
	if got := node.InsertNode(0, 0); checkNodeValue(got, want) != nil {
		t.Fatalf("want = %v, got = %v", want, got)
	}

	node = &Node{1, &Node{2, &Node{3, &Node{4, &Node{0, nil}}}}}
	want = &Node{1, &Node{2, &Node{5, &Node{3, &Node{4, &Node{0, nil}}}}}}
	if got := node.InsertNode(5, 2); checkNodeValue(got, want) != nil {
		t.Fatalf("want = %v, got = %v", want, got)
	}

	node = &Node{1, &Node{2, &Node{3, &Node{4, &Node{0, nil}}}}}
	want = node
	if got := node.InsertNode(5, 5); checkNodeValue(got, want) != nil {
		t.Fatalf("want = %v, got = %v", want, got)
	}
}

func TestReverseNode(t *testing.T) {
	node := &Node{1, &Node{2, &Node{3, &Node{4, &Node{0, nil}}}}}
	want := &Node{4, &Node{3, &Node{2, &Node{1, &Node{0, nil}}}}}
	if got := node.ReverseList(); checkNodeValue(got, want) != nil {
		t.Fatalf("want = %v, got = %v", want, got)
	}

	node = &Node{2, &Node{0, nil}}
	want = node
	if got := node.ReverseList(); checkNodeValue(got, want) != nil {
		t.Fatalf("want = %v, got = %v", want, got)
	}

}

func TestCreateList(t *testing.T) {
	want := &Node{0, nil}
	if got := CreateList(); checkNodeValue(got, want) != nil {
		t.Fatalf("want = %v, got = %v", want, got)
	}
}
