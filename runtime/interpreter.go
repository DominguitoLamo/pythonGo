package runtime

import (
	"fmt"

	"chonlam.com/pythonGo/code"
	"chonlam.com/pythonGo/pyType"
)

type Interpreter struct {
	stack []pyType.HiObject;
	consts []pyType.HiObject;
	loopStack []Block;
}

func (inter *Interpreter) Run(codes code.CodeObject) {
	pc := 0
	codeLength := codes.ByteCodes.Length

	inter.stack = make([]pyType.HiObject, 0)
	inter.consts = codes.Consts
	inter.loopStack = make([]Block, 0)
	names := codes.Names
	locals := make(map[pyType.HiObject] pyType.HiObject)

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
		case code.LOAD_NAME:
			v := names[opArg]
			w := locals[v]
			_, ok := w.(pyType.HiNil)
			if (!ok) {
				inter.pushStack(w)
			}
		case code.STORE_NAME:
			v := names[opArg]
			locals[v] = inter.popStack()
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
		case code.SETUP_LOOP:
			inter.pushLoop(Block{
				BType: opCode,
				Target: pc + opArg,
				Level: inter.getStackLevel(),
			})
		case code.POP_BLOCK:
			b := inter.popLoop()
			for (inter.getStackLevel() > b.Level) {
				inter.popStack()
			}
		case code.BREAK_LOOP:
			b := inter.popLoop()
			for (inter.getStackLevel() > b.Level) {
				inter.popStack()
			}
			pc = b.Target
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

func (inter *Interpreter) popLoop() Block {
	length := len(inter.loopStack)
	result := inter.loopStack[length - 1]
	inter.loopStack = inter.loopStack[:length - 1]
	return result
}

func (inter *Interpreter) pushLoop(b Block) {
	inter.loopStack = append(inter.loopStack, b)
}

func (inter *Interpreter) getStackLevel() int {
	return len(inter.stack)
}