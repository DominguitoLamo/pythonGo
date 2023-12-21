package pyType

type HiInterger struct {
	value int;
}

func CreateInteger(v int) *HiInterger {
	h := new(HiInterger)
	h.value = v
	return h
}