package code

import (
	"chonlam.com/pythonGo/pyType"
)

type CodeObject struct {
	consts []*pyType.HiObject;
	names []*pyType.HiObject;
	varNames []*pyType.HiObject;
	freeVars []*pyType.HiObject;
	cellVars []*pyType.HiObject;
	fileName *pyType.HiString;
	coName *pyType.HiString;
	lineNum int;
}