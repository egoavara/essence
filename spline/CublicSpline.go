package spline

import (
	"fmt"
	"github.com/pkg/errors"
	"sort"
	"strings"
)

type CubicSpline struct {
	x    []float32
	fx   []float32
	ddfx []float32
	fns  []func(x float32) float32
}

func NewCubicSpline(x []float32, fx []float32) (*CubicSpline, error) {
	if len(x) < 3 {
		return nil, errors.New("At least 3 points for CubicSpline")
	}
	if len(x) != len(fx) {
		return nil, errors.New("x.len != fx.len")
	}
	if !sort.IsSorted(F32Slice(x)) {
		return nil, errors.New("Unsorted data")
	}
	counts := len(x)
	fi1Mfi := make([]float32, counts-1) // f[i+1] - f[i]
	for i := range fi1Mfi {
		fi1Mfi[i] = fx[i+1] - fx[i]
	}
	xi1Mxi := make([]float32, counts-1) // x[i+1] - x[i]
	for i := range xi1Mxi {
		xi1Mxi[i] = x[i+1] - x[i]
	}
	// = xi1Mxi[i] * ddf[i] + 2*(xi1Mxi[i + 1] + xi1Mxi[i]) * ddf[i+1] + xi1Mxi[i + 1]ddf[i+2
	eq0 := make([]float32, counts-2) // 6 / xi1Mxi[i+1] * fi1Mfi[i+1] - 6 / xi1Mxi[i] * fi1Mfi[i]
	for i := range eq0 {
		eq0[i] = 6/xi1Mxi[i+1]*fi1Mfi[i+1] - 6/xi1Mxi[i]*fi1Mfi[i]
	}
	trm := NewTridiagonalMatrix(len(eq0))
	for i := range trm {
		if i == 0 {
			trm[i][1] = 2 * (xi1Mxi[i+1] + xi1Mxi[i])
			trm[i][2] = xi1Mxi[i+1]
		} else if i == len(trm)-1 {
			trm[i][0] = xi1Mxi[i]
			trm[i][1] = 2 * (xi1Mxi[i+1] + xi1Mxi[i])
		} else {
			trm[i][0] = xi1Mxi[i]
			trm[i][1] = 2 * (xi1Mxi[i+1] + xi1Mxi[i])
			trm[i][2] = xi1Mxi[i+1]
		}
	}

	ddf := append([]float32{0}, trm.Calculate(eq0...)...)
	ddf = append(ddf, 0)
	fns := make([]func(x float32) float32, counts-1)
	fns[0] = mkFunc(x[0], x[1], fx[0], fx[1], ddf[0], ddf[1])
	for i := 1; i < counts-2; i++ {
		fmt.Println(ddf, i)
		fns[i] = mkFunc(
			x[i],
			x[i+1],
			fx[i],
			fx[i+1],
			ddf[i],
			ddf[i+1],
		)
	}
	fns[counts-2] = mkFunc(x[counts-2], x[counts-1], fx[counts-2], fx[counts-1], ddf[counts-2], ddf[counts-1])
	return &CubicSpline{
		x:    x,
		fx:   fx,
		ddfx: ddf,
		fns:  fns,
	}, nil
}
func (s *CubicSpline) Get(x float32) float32 {
	var i = 0
	var v float32
	for i, v = range s.x {
		if x < v {
			break
		}
	}
	i = iclamp(i-1, 0, len(s.fns)-1)
	return s.fns[i](x)
}
func (s *CubicSpline) Range() (min, max float32) {
	return s.x[0], s.x[len(s.x)-1]
}

type TridiagonalMatrix [][3]float32

func (s TridiagonalMatrix) String() string {
	const space = "        "
	res := ""
	for i, v := range s {
		switch i {
		case 0:
			res += fmt.Sprintf("%s%7.4f %7.4f", strings.Repeat(space, i), v[1], v[2])
		case len(s) - 1:
			res += fmt.Sprintf("%s%7.4f %7.4f", strings.Repeat(space, i-1), v[0], v[1])
		default:
			res += fmt.Sprintf("%s%7.4f %7.4f %7.4f", strings.Repeat(space, i-1), v[0], v[1], v[2])
		}
		res += "\n"
	}
	return res
}
func (s TridiagonalMatrix) Calculate(results ...float32) []float32 {
	if len(results) != len(s) {
		return nil
	}
	last := len(s) - 1
	cs := make([]float32, len(s))
	ds := make([]float32, len(s))
	cs[0] = s[0][2] / s[0][1]
	ds[0] = results[0] / s[0][1]
	for i := 1; i < last; i++ {
		cs[i] = s[i][2] / (s[i][1] - s[i][0]*cs[i-1])
		ds[i] = (results[i] - s[i][0]*ds[i-1]) / (s[i][1] - s[i][0]*cs[i-1])
	}
	//cs[last] = 0
	ds[last] = (results[last] - s[last][0]*ds[last-1]) / (s[last][1] - s[last][0]*cs[last-1])
	//
	for i := last - 1; i >= 0; i-- {
		ds[i] -= cs[i] * ds[i+1]
	}
	return ds
}
func NewTridiagonalMatrix(length int) TridiagonalMatrix {
	return make([][3]float32, length)
}

func mkFunc(x0, x1, f0, f1, dff0, dff1 float32) func(x float32) float32 {
	return func(x float32) float32 {
		xMx0 := x - x0
		x1Mx := x1 - x
		dx := x1 - x0
		//
		a := dff0/6/dx*x1Mx*x1Mx*x1Mx + (f0/dx-dff0*dx/6)*x1Mx
		b := dff1/6/dx*xMx0*xMx0*xMx0 + (f1/dx-dff1*dx/6)*xMx0
		return a + b
	}
}
