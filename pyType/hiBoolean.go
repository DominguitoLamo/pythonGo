package pyType

func HiTrue() HiInterger {
	return HiInterger{value: 1}
}

func HiFalse() HiInterger {
	return HiInterger{value: 0}
}

func IsHiTrue(x HiObject) bool {
	switch v := x.(type) {
	case HiInterger:
		return v.value == 1
	default:
		return false
	}
}

func IsHiFalse(x HiObject) bool {
	switch v := x.(type) {
	case HiInterger:
		return v.value == 0
	default:
		return false
	}
}