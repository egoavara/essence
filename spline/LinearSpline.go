package spline

import (
	"github.com/pkg/errors"
	"sort"
)

type LinearSpline struct {
	x    []float32
	fx   []float32
	dydx []float32
}

func NewLinearSpline(x []float32, fx []float32) (*LinearSpline, error) {
	if len(x) < 2 {
		return nil, errors.New("At least 2 points for Linear Spline")
	}
	if len(x) != len(fx) {
		return nil, errors.New("x.len != fx.len")
	}
	if !sort.IsSorted(F32Slice(x)) {
		return nil, errors.New("Unsorted data")
	}
	//
	dydx := make([]float32, len(x)-1)
	for i := range dydx {
		dydx[i] = (fx[i+1] - fx[i]) / (x[i+1] - x[i])
	}
	return &LinearSpline{
		x:    x,
		fx:   fx,
		dydx: dydx,
	}, nil
}
func (s *LinearSpline) Get(x float32) float32 {
	var i = 0
	var v float32
	for i, v = range s.x {
		if x < v {
			break
		}
	}
	i = iclamp(i-1, 0, len(s.dydx)-1)
	return s.dydx[i]*(x-s.x[i]) + s.fx[i]
}

func (s *LinearSpline) Range() (min, max float32) {
	return s.x[0], s.x[len(s.x)-1]
}
