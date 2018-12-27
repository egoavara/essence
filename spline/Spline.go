package spline

type Spline interface {
	Get(x float32) float32
	Range() (min, max float32)
}
