package align

type Align uint8

const (
	No       Align = iota
	Zero     Align = iota
	Negative Align = iota
	Positive Align = iota
)
const (
	// Aliases
	Center = Zero
	Left   = Negative
	Back   = Negative
	Bottom = Negative
	Right   = Positive
	Forward = Positive
	Top     = Positive
)
