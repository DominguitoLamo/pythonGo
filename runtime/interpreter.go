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
			inter.pushStack(inter.consts[opArg])
		case code.PRINT_ITEM:
			v := inter.popStack()
			v.Print()
		case code.PRINT_NEWLINE:
			fmt.Printf("\n")
		case code.BINARY_ADD:
			l := inter.popStack()
			r := inter.popStack()
			inter.pushStack(l.Add(r))
		case code.COMPARE_OP:
			r := inter.popStack()
			l := inter.popStack()

			switch (opArg) {
			case code.COMPARE_GREATER:
				inter.pushStack(l.Greater(r))
			case code.COMPARE_LESS:
				inter.pushStack(l.Less(r))
			case code.COMPARE_EQUAL:
				inter.pushStack((l.Equal(r)))
			case code.COMPARE_NOT_EQUAL:
				inter.pushStack((l.NotEqual(r)))
			case code.COMPARE_GREATER_EQUAL:
				inter.pushStack(l.Ge(r))
			case code.COMPARE_LESS_EQUAL:
				inter.pushStack(l.Le(r))
			default:
				panic("Error: Unrecognized compare op")
			}
		case code.RETURN_VALUE:
			inter.popStack()
		case code.POP_JUMP_IF_FALSE:
			v := inter.popStack()
			if pyType.IsHiFalse(v) {
				pc = opArg
			}
		case code.JUMP_FORWARD:
			pc += opArg
		case code.JUMP_ABSOLUTE:
			pc = opArg
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

func (inter *Interpreter) pushStack(obj pyType.HiObject) {
	inter.stack = append(inter.stack, obj)
}