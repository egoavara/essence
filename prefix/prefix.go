package prefix

type Prefix uint8

func (s Prefix) String() string {
	switch s {
	case Yotta:
		return "Yotta"
	case Zetta:
		return "Zetta"
	case Exa:
		return "Exa"
	case Peta:
		return "Peta"
	case Tera:
		return "Tera"
	case Giga:
		return "Giga"
	case Mega:
		return "Mega"
	case Kilo:
		return "Kilo"
	case Hecto:
		return "Hecto"
	case Deca:
		return "Deca"
	case No:
		return "No"
	case Deci:
		return "Deci"
	case Centi:
		return "Centi"
	case Milli:
		return "Milli"
	case Micro:
		return "Micro"
	case Nano:
		return "Nano"
	case Pico:
		return "Pico"
	case Femto:
		return "Femto"
	case Atto:
		return "Atto"
	case Zepto:
		return "Zepto"
	case Yocto:
		return "Yocto"
	}
	panic("unreachable")
}
func (s Prefix) Symbol() string {
	switch s {
	case Yotta:
		return "Y"
	case Zetta:
		return "Z"
	case Exa:
		return "E"
	case Peta:
		return "P"
	case Tera:
		return "T"
	case Giga:
		return "G"
	case Mega:
		return "M"
	case Kilo:
		return "k"
	case Hecto:
		return "h"
	case Deca:
		return "de"
	case No:
		return ""
	case Deci:
		return "d"
	case Centi:
		return "c"
	case Milli:
		return "m"
	case Micro:
		return "µ"
	case Nano:
		return "n"
	case Pico:
		return "p"
	case Femto:
		return "f"
	case Atto:
		return "a"
	case Zepto:
		return "z"
	case Yocto:
		return "y"
	}
	panic("unreachable")
}
func (s Prefix) Exponent() int {
	switch s {
	case Yotta:
		return 24
	case Zetta:
		return 21
	case Exa:
		return 18
	case Peta:
		return 15
	case Tera:
		return 12
	case Giga:
		return 9
	case Mega:
		return 6
	case Kilo:
		return 3
	case Hecto:
		return 2
	case Deca:
		return 1
	case No:
		return 0
	case Deci:
		return -1
	case Centi:
		return -2
	case Milli:
		return -3
	case Micro:
		return -6
	case Nano:
		return -9
	case Pico:
		return -12
	case Femto:
		return -15
	case Atto:
		return -18
	case Zepto:
		return -21
	case Yocto:
		return -24
	}
	panic("unreachable")
}

const (
	Yotta Prefix = iota
	Zetta Prefix = iota
	Exa   Prefix = iota
	Peta  Prefix = iota
	Tera  Prefix = iota
	Giga  Prefix = iota
	Mega  Prefix = iota
	Kilo  Prefix = iota
	Hecto Prefix = iota
	Deca  Prefix = iota
	No    Prefix = iota
	Deci  Prefix = iota
	Centi Prefix = iota
	Milli Prefix = iota
	Micro Prefix = iota
	Nano  Prefix = iota
	Pico  Prefix = iota
	Femto Prefix = iota
	Atto  Prefix = iota
	Zepto Prefix = iota
	Yocto Prefix = iota
)
