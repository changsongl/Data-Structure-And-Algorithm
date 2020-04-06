package queue

type Queue struct {
	rear *Node
}

type Node struct {
	next *Node
	data int
}

func (q *Queue) Enqueue(data int) {
	n := q.rear
	q.rear = &Node{next: n, data: data}
}

func (q *Queue) Dequeue() (int, bool) {
	if q.rear == nil {
		return 0, false
	}

	if q.rear.next == nil {
		data := q.rear.data
		q.rear = nil
		return data, true
	}

	p := q.rear
	for p.next != nil {
		p = p.next
	}

	data := p.next.data
	p.next = nil
	return data, true
}

func (q *Queue) Peek() (int, bool) {
	if q.rear == nil {
		return 0, false
	}
	return q.rear.data, true
}

func (q *Queue) Get() []int {
	var dataSlice []int
	node := q.rear

	for node != nil {
		dataSlice = append(dataSlice, node.data)
		node = node.next
	}
	return dataSlice
}

func (q *Queue) IsEmpty() bool {
	return q.rear == nil
}

func (q *Queue) Empty() {
	q.rear = nil
}
