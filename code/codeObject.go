package code

import (
	"fmt"

	"chonlam.com/pythonGo/pyType"
)

type CodeObject struct {
	Consts []pyType.HiObject;
	Names []pyType.HiObject;
	VarNames []pyType.HiObject;
	FreeVars []pyType.HiObject;
	CellVars []pyType.HiObject;
	FileName pyType.HiString;
	CoName pyType.HiString;
	LineNum int;
	NoTable pyType.HiString;
	ArgCount int;
	NLocals int;
	StackSize int;
	Flags int;
	ByteCodes pyType.HiString
}

func (h CodeObject) Print() {
	fmt.Printf("")
}

func (h CodeObject) Add(x pyType.HiObject) pyType.HiObject {
	return nil
}
