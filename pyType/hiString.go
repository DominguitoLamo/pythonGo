package pyType

import "fmt"

type HiString struct {
	Length int
	Value  string
}

func CreateHiString(s string) HiString {
	h := HiString{}
	h.Value = s
	h.Length = len(s)
	return h
}

func (h HiString) Print() {
	fmt.Printf("%s", h.Value)
}

func (h HiString) Add(x HiObject) HiObject {
	switch v := x.(type) {
	case HiInterger:
		result := HiObject(CreateHiString(fmt.Sprintf("%s%d", h.Value, v.value)))
		return result
	case HiString:
		result := HiObject(CreateHiString(fmt.Sprintf("%s%s", h.Value, v.Value)))
		return result
	default:
		panic("No suitable type!!")
	}
}