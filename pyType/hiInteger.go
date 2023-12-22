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