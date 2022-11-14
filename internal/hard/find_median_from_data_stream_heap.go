package hard

import (
	"container/heap"
	"github.com/lelikptz/leetcode-go/internal/structure"
)

type MedianFinderHeap struct {
	low  *structure.IntMaxHeap
	high *structure.IntMinHeap
}

func NewMedianFinderHeap() MedianFinderHeap {
	low := &structure.IntMaxHeap{}
	heap.Init(low)

	high := &structure.IntMinHeap{}
	heap.Init(high)

	return MedianFinderHeap{low, high}
}

func (m *MedianFinderHeap) AddNum(num int) {
	if m.high.Len() == 0 {
		heap.Push(m.high, num)
	} else if num < (*m.high)[0] {
		heap.Push(m.low, num)
	} else {
		heap.Push(m.high, num)
	}

	diff := m.high.Len() - m.low.Len()
	if diff > 1 {
		heap.Push(m.low, heap.Pop(m.high))
	} else if diff < 0 {
		heap.Push(m.high, heap.Pop(m.low))
	}
}

func (m *MedianFinderHeap) FindMedian() float64 {
	if m.high.Len() == m.low.Len() {
		return float64((*m.low)[0]+(*m.high)[0]) / 2
	}

	return float64((*m.high)[0])
}
