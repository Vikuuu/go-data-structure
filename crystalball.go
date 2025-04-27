package datstr

import "math"

func TwoCrystalBalls(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jumpAmount

	for ; i < len(breaks); i = i + jumpAmount {
		if breaks[i] {
			break
		}
	}
	for j := i - jumpAmount; j < i; j++ {
		if breaks[j] {
			return j
		}
	}

	return -1
}

// Time Complexity: O(sqrt(n))
