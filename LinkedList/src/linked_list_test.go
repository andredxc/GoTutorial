package linked_list

import (
	"math/rand"
	"testing"
)

func TestFirstAdd(t *testing.T) {

	var got int
	var ok bool
	val := 5

	testCase := func(l *List, val int) {
		if l.size != 1 {
			t.Errorf("Size is not 1 after the first Add")
		}
		if l.first != l.last {
			t.Errorf("List's first and last are not the same after the first Add")
		}
		if got, ok = l.first.data.(int); !ok {
			t.Errorf("Could not retrieve val=%d added", val)
		}
		if got != val {
			t.Errorf("Data not added, want=%d, got=%d", val, got)
		}
	}

	l := new(List)
	l.Add(val)
	testCase(l, val)
	// Lower case variables can be accessed because we are in the same linked_list package
	l = nil
	l = new(List)
	l.Insert(0, val)
	testCase(l, val)
}

// Insert element at position 0 when the list already has 1 or more elements
func TestInsertAsFirstNode(t *testing.T) {

}

// Insert element at the last position when the list already has 1 or more elements
func TestInsertAsLastNode(t *testing.T) {

}

// Insert element in the body of the list
func TestMiddleInsert(t *testing.T) {

}

func TestRandomInserts(t *testing.T) {

	var ok bool

	listSize := 10
	insertions := 10
	// Setup
	l := new(List)
	for i := 0; i < listSize; i++ {
		l.Add(rand.Int())
	}

	for i := 0; i < insertions; i++ {
		ind := rand.Int() % l.size
		val := rand.Int()
		l.Insert(ind, val)

		got := l.Size()
		want := listSize + i + 1
		if got != want {
			t.Errorf("Size, got: %d, want: %d", got, want)
		}

		if got, ok = l.At(ind).(int); !ok {
			t.Errorf("Type, could not retrieve int value for index=%d", ind)
		}

		want = val
		if got != want {
			t.Errorf("Elements, got: %d, want: %d", got, want)
		}
	}
}
