package pyType

import (
	"fmt"
)

type HiInterger struct {
	value int
}

func CreateInteger(v int) HiInterger {
	h := HiInterger{}
	h.value = v
	return h
}

func (h HiInterger) Print() {
	fmt.Printf("%d", h.value)
}

func (h HiInterger) Add(x HiObject) HiObject {
	switch v := x.(type) {
	case HiInterger:
		result := HiObject(CreateInteger(h.value + v.value))
		return result
	case HiString:
		result := HiObject(CreateHiString(fmt.Sprintf("%d%s", h.value, v.Value)))
		return result
	default:
		panic("No suitable type!!")
	}
}

func (h HiInterger) Sub(x HiObject) HiObject {
	switch v := x.(type) {
	case HiInterger:
		result := HiObject(CreateInteger(h.value - v.value))
		return result
	case HiString:
		result := HiObject(CreateHiString(fmt.Sprintf("%d%s", h.value, v.Value)))
		return result
	default:
		panic("No suitable type!!")
	}
}

func (h HiInterger) Mul(x HiObject) HiObject {
	switch v := x.(type) {
	case HiInterger:
		result := HiObject(CreateInteger(h.value * v.value))
		return result
	case HiString:
		result := HiObject(CreateHiString(fmt.Sprintf("%d%s", h.value, v.Value)))
		return result
	default:
		panic("No suitable type!!")
	}
}

func (h HiInterger) Div(x HiObject) HiObject {
	switch v := x.(type) {
	case HiInterger:
		result := HiObject(CreateInteger(h.value / v.value))
		return result
	case HiString:
		result := HiObject(CreateHiString(fmt.Sprintf("%d%s", h.value, v.Value)))
		return result
	default:
		panic("No suitable type!!")
	}
}

func (h HiInterger) Mod(x HiObject) HiObject {
	switch v := x.(type) {
	case HiInterger:
		result := HiObject(CreateInteger(h.value % v.value))
		return result
	case HiString:
		result := HiObject(CreateHiString(fmt.Sprintf("%d%s", h.value, v.Value)))
		return result
	default:
		panic("No suitable type!!")
	}
}

func (h HiInterger) Greater(x HiObject) HiObject {
	switch v:= x.(type) {
	case HiInterger:
		if (h.value > v.value) {
			return HiTrue()
		} else {
			return HiFalse()
		}
	default:
		panic("This is not numeric comparison!!")
	}
}

func (h HiInterger) Less(x HiObject) HiObject {
	switch v:= x.(type) {
	case HiInterger:
		if (h.value < v.value) {
			return HiTrue()
		} else {
			return HiFalse()
		}
	default:
		panic("This is not numeric comparison!!")
	}
}

func (h HiInterger) Equal(x HiObject) HiObject {
	switch v:= x.(type) {
	case HiInterger:
		if (h.value == v.value) {
			return HiTrue()
		} else {
			return HiFalse()
		}
	default:
		panic("This is not numeric comparison!!")
	}
}

func (h HiInterger) NotEqual(x HiObject) HiObject {
	switch v:= x.(type) {
	case HiInterger:
		if (h.value != v.value) {
			return HiTrue()
		} else {
			return HiFalse()
		}
	default:
		panic("This is not numeric comparison!!")
	}
}

func (h HiInterger) Ge(x HiObject) HiObject {
	switch v:= x.(type) {
	case HiInterger:
		if (h.value >= v.value) {
			return HiTrue()
		} else {
			return HiFalse()
		}
	default:
		panic("This is not numeric comparison!!")
	}
}

func (h HiInterger) Le(x HiObject) HiObject {
	switch v:= x.(type) {
	case HiInterger:
		if (h.value <= v.value) {
			return HiTrue()
		} else {
			return HiFalse()
		}
	default:
		panic("This is not numeric comparison!!")
	}
}