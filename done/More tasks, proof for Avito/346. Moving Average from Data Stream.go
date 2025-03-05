package main

type BufferedQueue struct {
	buffer  []int
	head    int
	tail    int
	sum     int
	size    int
	maxSize int
}

func (this *BufferedQueue) Enqueue(val int) {
	if this.size < this.maxSize { // считаем, что maxSize > 0
		this.tail = (this.tail + 1) % this.maxSize
		this.buffer[this.tail] = val
		this.sum += val
		this.size++
		return
	}
	removed := this.buffer[this.head]
	this.buffer[this.head] = val
	this.tail = this.head
	this.head = (this.head + 1) % this.maxSize
	this.sum += val - removed
}

type MovingAverage struct {
	queue BufferedQueue
}

/** Initialize your data structure here. */
func ConstructorMA(size int) MovingAverage {
	if size == 0 {
		panic("size must be greater than 0")
	}
	buffer := make([]int, size)
	queue := BufferedQueue{buffer, 0, -1, 0, 0, size}
	return MovingAverage{queue}
}

func (this *MovingAverage) Next(val int) float64 {
	this.queue.Enqueue(val)
	return float64(this.queue.sum) / float64(this.queue.size)
}
