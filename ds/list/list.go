package list

type Node struct {
	next  *Node
	prev  *Node
	value interface{}
}

func (n *Node) Value() interface{} { return n.value }
func (n *Node) Next() *Node        { return n.next }
func (n *Node) Prev() *Node        { return n.prev }

type List struct {
	head *Node
	tail *Node
	size int
}

func New() *List            { return &List{} }
func (l *List) Head() *Node { return l.head }
func (l *List) Tail() *Node { return l.tail }
func (l *List) Size() int   { return l.size }

func (l *List) Traverse(doSmth func(node *Node) (stop bool), reverse bool) *Node {
	var node *Node
	if reverse {
		node = l.tail
	} else {
		node = l.head
	}

	for node != nil {
		if doSmth(node) {
			return node
		}

		if reverse {
			node = node.prev
		} else {
			node = node.next
		}
	}

	return nil
}

func (l *List) InsertAfter(node *Node, value interface{}) *Node {
	if node == nil {
		return nil
	}

	newNode := &Node{
		value: value,
		next:  node.next,
		prev:  node,
	}

	if node.next != nil {
		node.next.prev = newNode
	} else {
		l.tail = newNode
	}

	node.next = newNode
	l.size++
	return newNode
}

func (l *List) InsertBefore(node *Node, value interface{}) *Node {
	if node == nil {
		return nil
	}

	newNode := &Node{
		value: value,
		next:  node,
		prev:  node.prev,
	}

	if node.prev != nil {
		node.prev.next = newNode
	} else {
		l.head = newNode
	}

	node.prev = newNode
	l.size++
	return newNode
}

func (l *List) PushFront(value interface{}) *Node {
	if l.head != nil {
		return l.InsertBefore(l.head, value)
	}

	newNode := &Node{value: value}
	l.head = newNode
	l.tail = newNode
	l.size++

	return newNode
}

func (l *List) PushBack(value interface{}) *Node {
	if l.tail != nil {
		return l.InsertAfter(l.tail, value)
	}

	return l.PushFront(value)
}

func (l *List) Remove(node *Node) {
	if node == nil {
		return
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}

	l.size--
}

func (l *List) Union(list *List) *List {
	if list == nil || list.head == nil || list.tail == nil {
		return l
	}

	if l == nil || l.head == nil || l.tail == nil {
		return list
	}

	l.tail.next = list.head
	list.head.prev = l.tail
	l.tail = list.tail
	l.size += list.size

	return l
}
