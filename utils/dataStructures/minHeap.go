package dataStructures

import "ProcessManger/types"
type MinHeap struct {
	items []types.Process
}

func (h *MinHeap) Insert(key types.Process) {
	h.items = append(h.items, key)
	h.HeapifyUp(len(h.items) - 1)
}

func (h *MinHeap) HeapifyUp(index int) {
	for compareProcess(h.items[parent(index)], h.items[index]) {
		h.Swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) HeapifyDown(index int, minHeap bool) {
	lastIndex := len(h.items) - 1
	smallest := index
	for {
		leftChildIndex := leftChild(index)
		rightChildIndex := rightChild(index)
		if leftChildIndex <= lastIndex && ((minHeap && compareProcess(h.items[leftChildIndex], h.items[smallest])) || (!minHeap && compareProcess(h.items[leftChildIndex], h.items[smallest]))) {
			smallest = leftChildIndex
		}
		if rightChildIndex <= lastIndex && ((minHeap && compareProcess(h.items[rightChildIndex], h.items[smallest])) || (!minHeap && compareProcess(h.items[rightChildIndex], h.items[smallest]))) {
			smallest = rightChildIndex
		}
		if smallest != index {
			h.Swap(index, smallest)
			index = smallest
		} else {
			break
		}
	}
}

func (h *MinHeap) Remove(index int) types.Process {
	toRemove := h.items[index]
	l := len(h.items) - 1
	if l == -1 {
		return types.Process{}
	}
	h.items[index] = h.items[l]
	h.items = h.items[:l]
	h.HeapifyDown(index, true)
	return toRemove
}

func (h *MinHeap) Extract() types.Process {
	return h.Remove(0)
}

func (h *MinHeap) Swap(index1 int, index2 int) {
	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}
