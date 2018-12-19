package spline

import (
	"github.com/pkg/errors"
	"sort"
)

type StepSpline struct {
	x    []float32
	fx   []float32
}

func NewStepSpline(x []float32, fx []float32) (*StepSpline, error) {
	if len(x) < 1 {
		return nil, errors.New("At least 1 points for Step Spline")
	}
	if len(x) != len(fx) {
		return nil, errors.New("x.len != fx.len")
	}
	if !sort.IsSorted(F32Slice(x)) {
		return nil, errors.New("Unsorted data")
	}
	return &StepSpline{
		x:    x,
		fx:   fx,
	}, nil
}
func (s *StepSpline) Get(x float32) float32 {
	var i = 0
	var v float32
	for i, v = range s.x {
		if x < v {
			break
		}
	}
	i = iclamp(i-1, 0, len(s.fx) - 1)
	return s.fx[i]
}

func (s *StepSpline) Range() (min, max float32) {
	return s.x[0], s.x[len(s.x)-1]
}
