package main

import (
	"errors"
	"fmt"
)

type Heap struct {
	heap []int
}

func (h *Heap) getLeft(nodeIndex int) (error, int) {
	if 2*nodeIndex < len(h.heap) {
		return nil, h.heap[2*nodeIndex]
	}
	return errors.New("no left element"), -1
}

func (h *Heap) getRight(nodeIndex int) (error, int) {
	if 2*nodeIndex+1 < len(h.heap) {
		return nil, h.heap[2*nodeIndex+1]
	}
	return errors.New("no right element"), -1
}

func (h *Heap) getRoot(nodeIndex int) (error, int) {
	if nodeIndex/2 >= 1 {
		return nil, h.heap[nodeIndex/2]
	}
	return errors.New("no left element"), -1
}

func (h *Heap) getItemAt(index int) (error, int) {
	if index >= 1 && index <= h.heap[0] {
		return nil, h.heap[index]
	}
	return errors.New("index out of bounds"), -1
}

func (h *Heap) getSize() int {
	return h.heap[0]
}

func (h *Heap) swap(i, j int) {
	t := h.heap[i]
	h.heap[i] = h.heap[j]
	h.heap[j] = t
}

func (h *Heap) getItems() []int {
	return h.heap[1:]
}

func (h *Heap) popMin() int {
	el := h.heap[1]
	h.heap[1] = h.heap[h.heap[0]]
	h.heap = h.heap[:h.heap[0]]
	h.heap[0] -= 1
	h.bubbleDown(1)
	return el
}

func (h *Heap) add(data int) {
	if len(h.heap) == 0 {
		h.heap = append(h.heap, 0)
	}
	h.heap[0] += 1
	h.heap = append(h.heap, data)
	h.heapify(h.heap[0])
}

func (h *Heap) heapify(i int) {
	if i == 1 {
		return
	}
	err, r := h.getRoot(i)
	c := h.heap[i]

	if err == nil {
		if r > c {
			h.swap(i/2, i)
			h.heapify(i / 2)
		}
	} else {
	}
}

func (h *Heap) bubbleDown(i int) {
	err, elementAsParent := h.getItemAt(i)
	leftErr, left := h.getLeft(i)
	rightErr, right := h.getRight(i)
	if err != nil {
		return
	}
	if (leftErr == nil && elementAsParent < left) && (rightErr == nil && elementAsParent < right) {
		if leftErr != nil {
		}
		if rightErr != nil {
		}
		return
	}
	if leftErr != nil && rightErr != nil {
		return
	}
	var smallest int
	if leftErr != nil {
		smallest = 2*i + 1
	} else if rightErr != nil {
		smallest = 2 * i
	} else {
		if left <= right {
			smallest = 2 * i
		} else {
			smallest = 2*i + 1
		}
	}
	h.swap(i, smallest)
	h.bubbleDown(smallest)
}

func main() {
	var h Heap
	h.add(5)
	h.add(6)
	h.add(1)
	h.add(2)
	h.add(0)
	h.add(8)
	h.add(7)
	h.add(6)
	h.add(3)
	for ; h.getSize() > 0; {
		fmt.Println(h.popMin())
	}

}
