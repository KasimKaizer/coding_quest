package snake

// use queue data structure, when we move any position, enqueue the new position and dequeue from the queue
// when we encounter a fruit, we enqueue the new position but don't dequeue, this allows the snake can grow
// create a method on the queue to search, this method would take a position as its argument, and would return
// true / false based on if the position exists in the queue.
// this is no need to create a board or space per say. but would need to know size of the space tp calculate
// out of bounds

type Position struct {
	X, Y int
}

func CalculateScore(fruits []Position, moves string) int {
	panic("implement this")
}
