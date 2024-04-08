package asteroids

import "fmt"

type Asteroid struct {
	XPos, YPos, XSpeed, YSpeed float64
}

// don't overthink questions, there might be a better algorithmic solution to this problem
// but I am not smart enough to figure it out, so after a lot of time searching, I have chosen
// just use the brute force approach

func FindSafePath(data []Asteroid, gridWidth, gridHeight int) string {
	timeDiff := float64(60 * 60)
	grid := generateGrid(gridWidth, gridHeight)
	for _, a := range data {
		x, y := a.XPos+(timeDiff*a.XSpeed), a.YPos+(timeDiff*a.YSpeed)
		if a.XSpeed == 0 && a.YSpeed == 0 {
			if int(x) > gridWidth-1 || x < 0 || int(y) > gridHeight-1 || y < 0 {
				continue
			}
			grid[int(y)][int(x)] = true
			continue
		}
		for i := 0; i < 60; x, y, i = (x + a.XSpeed), (y + a.YSpeed), i+1 {
			if int(x) > gridWidth-1 || x < 0 {
				continue
			}
			if int(y) > gridHeight-1 || y < 0 {
				continue
			}
			grid[int(y)][int(x)] = true

		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !grid[i][j] {
				return fmt.Sprintf("%d:%d", j, i)
			}
		}
	}
	return ""
}

func generateGrid(width, height int) [][]bool {
	grid := make([][]bool, height)
	for idx := range grid {
		grid[idx] = make([]bool, width)
	}
	return grid
}
