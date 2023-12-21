package pyType

type HiObject interface {
	Print();
	Add(x *HiObject) *HiObject;
}