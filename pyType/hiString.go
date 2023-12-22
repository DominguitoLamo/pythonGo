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

func (h HiString) Greater(x HiObject) HiObject {
	switch v:= x.(type) {
	case HiString:
		hSum := h.getAsciiSum()
		vSum := v.getAsciiSum()
		if (hSum > vSum) {
			return HiTrue()
		} else {
			return HiFalse()
		}
	default:
		panic("This is not string comparison!!")
	}
}

func (h HiString) Less(x HiObject) HiObject {
	switch v:= x.(type) {
	case HiString:
		hSum := h.getAsciiSum()
		vSum := v.getAsciiSum()
		if (hSum < vSum) {
			return HiTrue()
		} else {
			return HiFalse()
		}
	default:
		panic("This is not string comparison!!")
	}
}

func (h HiString) Equal(x HiObject) HiObject {
	switch v:= x.(type) {
	case HiString:
		hSum := h.getAsciiSum()
		vSum := v.getAsciiSum()
		if (hSum == vSum) {
			return HiTrue()
		} else {
			return HiFalse()
		}
	default:
		panic("This is not string comparison!!")
	}
}

func (h HiString) NotEqual(x HiObject) HiObject {
	switch v:= x.(type) {
	case HiString:
		hSum := h.getAsciiSum()
		vSum := v.getAsciiSum()
		if (hSum != vSum) {
			return HiTrue()
		} else {
			return HiFalse()
		}
	default:
		panic("This is not string comparison!!")
	}
}

func (h HiString) Ge(x HiObject) HiObject {
	switch v:= x.(type) {
	case HiString:
		hSum := h.getAsciiSum()
		vSum := v.getAsciiSum()
		if (hSum >= vSum) {
			return HiTrue()
		} else {
			return HiFalse()
		}
	default:
		panic("This is not string comparison!!")
	}
}

func (h HiString) Le(x HiObject) HiObject {
	switch v:= x.(type) {
	case HiString:
		hSum := h.getAsciiSum()
		vSum := v.getAsciiSum()
		if (hSum <= vSum) {
			return HiTrue()
		} else {
			return HiFalse()
		}
	default:
		panic("This is not string comparison!!")
	}
}

func (h HiString) getAsciiSum() int {
	total := 0
	for i := 0; i < h.Length; i++ {
		total += int(h.Value[i])
	}

	return total
}