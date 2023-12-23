package pyType

import (
	"fmt"
)

type HiNil struct {
}


func (h HiNil) Print() {
	fmt.Printf("nil")
}

func (h HiNil) Add(x HiObject) HiObject {
	panic("Nil isn't for calculation!")
}

func (h HiNil) Sub(x HiObject) HiObject {
	panic("String is not for calculation!")
}

func (h HiNil) Mul(x HiObject) HiObject {
	panic("String is not for calculation!")
}

func (h HiNil) Div(x HiObject) HiObject {
	panic("String is not for calculation!")
}

func (h HiNil) Mod(x HiObject) HiObject {
	panic("String is not for calculation!")
}

func (h HiNil) Greater(x HiObject) HiObject {
	panic("Comparison is not allowed!")
}

func (h HiNil) Less(x HiObject) HiObject {
	panic("Comparison is not allowed!")
}

func (h HiNil) Equal(x HiObject) HiObject {
	panic("Comparison is not allowed!")
}

func (h HiNil) NotEqual(x HiObject) HiObject {
	panic("Comparison is not allowed!")
}

func (h HiNil) Ge(x HiObject) HiObject {
	panic("Comparison is not allowed!")
}

func (h HiNil) Le(x HiObject) HiObject {
	panic("Comparison is not allowed!")
}
