package queue

type Queue struct {
	data []int
	rear int
	front int
	size int
	empty bool
}

func NewQueue(size int) *Queue{
	data := make([]int, size)
	return &Queue{data: data, rear: 0, front: 0, size: size, empty: true}
}

func (q *Queue) Enqueue(data int) bool{
	if q.IsFull() {
		return false
	}

	q.rear = (q.rear + 1) % q.size
	q.empty = false
	q.data[q.rear] = data
	return true
}

func (q *Queue) Dequeue() (int, bool){
	if q.IsEmpty() {
		return 0, false
	}

	data := q.data[q.front]
	q.front = (q.front + 1) % q.size
	if q.front == q.rear {
		q.empty = true
	}
	return data, true
}

func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.data[q.front], true
}

func (q *Queue) Get() ([]int, bool){
	if q.IsEmpty(){
		return nil, false
	}

	var data []int
	for i := q.rear; i != q.front; i = (i + 1) % q.size {
		data = append(data, q.data[i])
	}
	return data, true
}

func (q *Queue) Empty() {
	for i := range q.data {
		q.data[i] = 0
	}
}

func (q *Queue) IsEmpty() bool{
	return q.empty
}

func (q *Queue) IsFull() bool{
	return q.front == q.front && !q.empty
}