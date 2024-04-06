package travel

import "math"

type Coordinate struct {
	X int
	Y int
	Z int
}

func CalculateDistance(data []*Coordinate) int {
	totalDist := 0
	for i := 0; i < len(data)-1; i++ {
		cur := data[i]
		next := data[i+1]
		totalDist += CalculatePythagoras((next.X - cur.X), (next.Y - cur.Y), (next.Z - cur.Z))
	}
	return totalDist
}

func CalculatePythagoras(x, y, z int) int {
	return int(math.Sqrt(float64((x * x) + (y * y) + (z * z))))
}
