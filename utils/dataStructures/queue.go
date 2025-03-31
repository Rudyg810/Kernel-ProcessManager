package dataStructures

import (
	"ProcessManger/types"
)

type Queue struct {
	items []types.Process
}

func (q *Queue) Enqueue(item types.Process) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() types.Process {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Peek() types.Process {
	return q.items[0]
}

func (q *Queue) Contains(item types.Process) bool {
	for _, v := range q.items {
		if v.ID == item.ID {
			return true
		}
	}
	return false
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Clear() {
	q.items = []types.Process{}
}

func (q *Queue) Display() []types.Process {
	return q.items
}
