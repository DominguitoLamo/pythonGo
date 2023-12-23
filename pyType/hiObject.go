package pyType

type HiObject interface {
	Print();

	Add(x HiObject) HiObject;
	Sub(x HiObject) HiObject;
	Mul(x HiObject) HiObject;
	Div(x HiObject) HiObject;
	Mod(x HiObject) HiObject;

	Greater(x HiObject) HiObject;
	Less(x HiObject) HiObject;
	Equal(x HiObject) HiObject;
	NotEqual(x HiObject) HiObject;
	Ge(x HiObject) HiObject;
	Le(x HiObject) HiObject;
}