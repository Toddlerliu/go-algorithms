package queue

// (slice)
type Queue struct {
	Ele []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		Ele: make([]interface{}, 0),
	}
}

func (q Queue) Empty() bool {
	if q.Ele == nil || len(q.Ele) == 0 {
		return true
	}
	return false
}

func (q Queue) Size() int {
	return len(q.Ele)
}

// 入队列
func (q *Queue) Offer(item interface{}) {
	q.Ele = append(q.Ele, item)
}

// 出队并删除
func (q *Queue) Poll() interface{} {
	if q.Size() > 0 {
		v := q.Ele[0]
		q.Ele = q.Ele[1:q.Size()] // [ )
		return v
	}
	return nil
}

// 出队不删除
func (q *Queue) Peek() interface{} {
	if q.Size() > 0 {
		return q.Ele[0]
	}
	return nil
}

func (q *Queue) GetAll() []interface{} {
	return q.Ele
}
