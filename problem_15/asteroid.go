package asteroid

// use depth first search,
// make a memory matrix, which would be marked when we traverse a square in real matrix
// use functional recursion, travel each four sides, if its zero then return zero immediately
// custom type is a overkill for such a basic problem

func CalAverMass(grid [][]int) int {
	mem := make([][]bool, len(grid))
	for idx := range mem {
		mem[idx] = make([]bool, len(grid[idx]))
	}
	sum := 0
	numOfAsteroid := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			num := search(grid, mem, i, j)
			if num == 0 {
				continue
			}
			numOfAsteroid++
			sum += num
		}
	}
	return sum / numOfAsteroid
}

func search(grid [][]int, mem [][]bool, row, col int) int {
	if row >= len(grid) || row < 0 ||
		col >= len(grid[row]) || col < 0 ||
		mem[row][col] || grid[row][col] == 0 {
		return 0
	}
	mem[row][col] = true
	sum := grid[row][col]
	// travel up
	sum += search(grid, mem, row-1, col)
	// travel right
	sum += search(grid, mem, row, col+1)
	// travel down
	sum += search(grid, mem, row+1, col)
	// travel left
	sum += search(grid, mem, row, col-1)
	return sum
}
