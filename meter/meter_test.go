package meter

import (
	"essence/prefix"
	"fmt"
	"testing"
)

func TestMeter(t *testing.T) {
	m := NewMeter(prefix.Centi, 100)
	fmt.Println(m)
	fmt.Println(m.Convert(prefix.No))
}
