package diagnostics

import "errors"

// use sliding window to slide from n to n-60
// data[n:60+n]
// calculate the initial sum
// data[n-1] from the sum then add data[n+60]
// excited about being able to actually use sliding window!!
// create a accumulator which account for all the time average crosses out defined limits

func FindInconsistences(data []int) (int, error) {

	if len(data) < 60 {
		return 0, errors.New("diagnostics.FindInconsistences: at least 60 seconds of data needed")
	}

	const (
		lowerBound = 1500
		upperBound = 1600
	)
	accum, sum := 0, 0

	for i := 0; i < 60; i++ {
		sum += data[i]
	}
	avg := sum / 60
	if avg < lowerBound || avg > upperBound {
		accum++
	}
	for i := 1; i < len(data)-60; i++ {
		sum = (sum - data[i-1]) + data[i+59]
		avg = sum / 60
		if avg < lowerBound || avg > upperBound {
			accum++
		}
	}
	return accum, nil

}
