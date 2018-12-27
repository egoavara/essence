package axis

type Axis uint8

const (
	X Axis = iota
	Y Axis = iota
	Z Axis = iota
)
const (
	// Alias
	Depth      = X
	Horizontal = Z
	Vertical   = Y
)
