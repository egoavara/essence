package spline

import "math"

type F32Slice []float32

func (s F32Slice) Len() int { return len(s) }
func (s F32Slice) Less(i, j int) bool {
	return s[i] < s[j] || math.IsNaN(float64(s[i])) && !math.IsNaN(float64(s[j]))
}
func (s F32Slice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func iclamp(i, min, max int) int {
	if i < min {
		return min
	}
	if i > max {
		return max
	}
	return i
}
