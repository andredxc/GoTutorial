package linked_list

type Lister interface {
	New()
	Add(obj interface{})
	Insert(ind int, obj interface{})
	Remove(obj interface{})
	Delete(ind int)
	At(int int) interface{}
	Size() int
	/*
		Front() interface{}
		Back() interface{}
	*/
}

type Noder interface {
	Prev() *Node
	Next() *Node
}

type List struct {
	size  int
	first *Node
	last  *Node
}

type Node struct {
	prev *Node
	next *Node
	data interface{}
}

func (l *List) Add(obj interface{}) {

	newNode := new(Node)
	newNode.data = obj

	if l.size == 0 {
		l.first = newNode
		l.last = l.first
	} else {
		newNode.prev = l.last
		l.last.next = newNode
		l.last = newNode
	}

	l.size++
}

func (l *List) Insert(ind int, obj interface{}) {

	// What to do with invalid ind parameter??? return error??? raise panic???
	// TODO: Spport negative indexes as counted backwards
	if ind >= l.size {
		ind = l.size - 1
	}
	if ind < 0 {
		ind = 0
	}

	curNode := l.findNodeAt(ind)

	if curNode != nil {
		l.addBefore(curNode, obj)
	} else {
		// List is empty
		l.Add(obj)
	}
}

func (l *List) addBefore(curNode *Node, obj interface{}) {
	newNode := new(Node)
	newNode.data = obj

	prevNode := curNode.prev
	if prevNode != nil {
		prevNode.next = newNode
	}

	newNode.prev = curNode.prev
	newNode.next = curNode.next

	curNode.prev = newNode

	if prevNode == nil {
		// First Node in the list
		l.first = newNode
	}
	// List's last never changes because we prepend to a not nil curNode
	l.size++
}

func (l *List) Remove(obj interface{}) {

	panic("Not implemented")
}

func (l *List) Delete(ind int) {

	if ind >= l.size {
		panic("Argument out of range")
	}

	curNode := l.findNodeAt(ind)

	if curNode == nil {
		panic("Internal error, index not found")
	}

	prevNode := curNode.prev
	nextNode := curNode.next

	if prevNode != nil {
		prevNode.next = curNode.next
	} else {
		l.first = nextNode
	}

	if nextNode != nil {
		nextNode.prev = curNode.prev
	} else {
		l.last = prevNode
	}

	l.size--
}

func (l *List) At(ind int) interface{} {

	curNode := l.findNodeAt(ind)

	if curNode == nil {
		panic("Argument out of range")
	}

	return curNode.data
}

func (l *List) Size() int {
	return l.size
}

func (l *List) findNodeAt(ind int) *Node {

	// TODO: Go from the end to the start when ind > l.size/2
	curNode := l.first
	curInd := 0
	for curNode != nil && (curInd < ind) {
		curNode = curNode.next
		curInd++
	}

	return curNode
}
