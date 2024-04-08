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

type Queue struct {
	data []Position
}

func (q *Queue) Head() Position {
	if len(q.data) == 0 {
		return Position{X: 0, Y: 0}
	}
	return q.data[len(q.data)-1]
}
func (q *Queue) Tail() Position {
	if len(q.data) == 0 {
		return Position{X: 0, Y: 0}
	}
	return q.data[0]
}

func (q *Queue) Enqueue(xPos, yPos int) {
	q.data = append(q.data, Position{X: xPos, Y: yPos})
}

func (q *Queue) Dequeue() Position {
	if len(q.data) == 0 {
		return Position{X: 0, Y: 0}
	}
	toReturn := q.data[0]
	q.data = q.data[1:]
	return toReturn
}

func (q *Queue) Exists(xPos, yPos int) bool {
	pos := Position{X: xPos, Y: yPos}
	for idx := range q.data {
		if q.data[idx] == pos {
			return true
		}
	}
	return false
}

func moveToPos(move byte) Position {
	switch move {
	case 'U':
		return Position{X: 0, Y: -1}
	case 'R':
		return Position{X: 1, Y: 0}
	case 'D':
		return Position{X: 0, Y: 1}
	case 'L':
		return Position{X: -1, Y: 0}
	}
	return Position{X: 0, Y: 0}
}

func CalculateScore(fruits *Queue, moves string, boardWidth, boardHeight int) int {
	snake := new(Queue)
	snake.Enqueue(0, 0)
	score := 0
	for _, move := range []byte(moves) {
		offSet := moveToPos(move)
		last := snake.Head()
		x, y := (last.X + offSet.X), (last.Y + offSet.Y)
		if x >= boardWidth || x < 0 ||
			y >= boardHeight || y < 0 ||
			snake.Exists(x, y) {
			return score
		}
		snake.Enqueue(x, y)
		score += 1
		if snake.Head() == fruits.Tail() {
			fruits.Dequeue()
			score += 100
			continue
		}
		snake.Dequeue()
	}
	return score
}
