package meter

import (
	"essence/prefix"
	"fmt"
	"testing"
)

func TestMeter(t *testing.T) {
	m := New(prefix.Centi, 100)
	fmt.Println(m)
	fmt.Println(m.Convert(prefix.No))
}
