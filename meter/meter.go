package meter

import (
	"fmt"
	"github.com/iamGreedy/essence/prefix"
	"math"
)

type Meter struct {
	prefix prefix.Prefix
	data float64
}
func (s Meter) String() string {
	return fmt.Sprintf("%s%sm", fmt.Sprint(s.data), s.prefix.Symbol())
}

func New(prefix prefix.Prefix, data float64) Meter {
	return Meter{prefix: prefix, data: data}
}
func (s Meter) Convert(pre prefix.Prefix) Meter {
	diff := s.prefix.Exponent() - pre.Exponent()
	return New(pre, s.data * math.Pow10(diff))
}
func (s Meter) F32() float32 {
	return float32(s.data)
}
func (s Meter) F64() float64 {
	return s.data
}


