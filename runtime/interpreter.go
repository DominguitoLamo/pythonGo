package runtime

import (
	"fmt"

	"chonlam.com/pythonGo/code"
	"chonlam.com/pythonGo/pyType"
)

type Interpreter struct {
	stack []pyType.HiObject;
	consts []pyType.HiObject;
}

func (inter *Interpreter) Run(codes code.CodeObject) {
	pc := 0
	codeLength := codes.ByteCodes.Length

	inter.stack = make([]pyType.HiObject, 0)
	inter.consts = codes.Consts

	for (pc < codeLength) {
		opCode := int(codes.ByteCodes.Value[pc])
		pc++
		hasArgument := (opCode & 0xFF) >= code.HAVE_ARGUMENT

		opArg := 0
		if (hasArgument) {
			byte1 := int(codes.ByteCodes.Value[pc] & 0xFF)
			pc++
			opArg = int(codes.ByteCodes.Value[pc] & 0xFF) << 8 | byte1
			pc++
		}

		switch (opCode) {
		case code.LOAD_CONST:
			inter.stack = append(inter.stack, inter.consts[opArg])
		case code.PRINT_ITEM:
			v := inter.popStack()
			v.Print()
		case code.PRINT_NEWLINE:
			fmt.Printf("\n")
		case code.BINARY_ADD:
			l := inter.popStack()
			r := inter.popStack()
			inter.stack = append(inter.stack, l.Add(r))
		case code.RETURN_VALUE:
			inter.popStack()
		default:
			fmt.Printf("Error: Unrecognized byte code %d\n", opCode)
		}
	}
}

func (inter *Interpreter) popStack() pyType.HiObject {
	length := len(inter.stack)
	result := inter.stack[length - 1]
	inter.stack = inter.stack[:length - 1]
	return result
}