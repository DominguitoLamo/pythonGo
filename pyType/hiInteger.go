package pyType

type HiInterger struct {
	HiObject
	value int;
}

func CreateInteger(v int) *HiInterger {
	h := new(HiInterger)
	h.value = v
	return h
}