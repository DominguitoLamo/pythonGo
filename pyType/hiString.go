package pyType

type HiString struct {
	HiObject
	length int;
	value string;
}

func CreateHiString(s string) *HiString {
	h := new(HiString)
	h.value = s
	h.length = len(s)
	return h
}

