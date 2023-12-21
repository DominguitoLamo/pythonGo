package code

import (
	"fmt"

	"chonlam.com/pythonGo/codeStream"
	"chonlam.com/pythonGo/pyType"
)

type BinaryFileParser struct {
	fileStream *codeStream.StringBuffer;
	cur int;
	stringTable []*pyType.HiString
}

func CreateBinaryFileParser(f *codeStream.StringBuffer) *BinaryFileParser {
	b := new(BinaryFileParser)
	b.fileStream = f
	b.cur = 0
	b.stringTable = make([]*pyType.HiString, 0)
	return b
}

func (b *BinaryFileParser) Parse() *CodeObject {
	magicNum := b.fileStream.ReadInt()
	fmt.Printf("magic number is %x\n", magicNum);
	modDate := b.fileStream.ReadInt();
	fmt.Printf("mod date is %x\n", modDate)

	objectType := string(b.fileStream.Read())
	if (objectType == "c") {
		fmt.Printf("parse Ok!\n")
		return b.getCodeObject()
	}

	return nil
}

func (b *BinaryFileParser) getCodeObject() *CodeObject {
	result := new(CodeObject)
	result.ArgCount = b.fileStream.ReadInt()
	fmt.Printf("arg count is %d\n", result.ArgCount)
	result.NLocals = b.fileStream.ReadInt()
	result.StackSize = b.fileStream.ReadInt()
	result.Flags = b.fileStream.ReadInt()
	fmt.Printf("flags is 0x%x\n", result.Flags)

	result.ByteCodes = b.getByteCodes()
	result.Consts = b.getConsts()
	result.Names = b.getNames()
	result.VarNames = b.getVarNames()
	result.FreeVars = b.getFreeVars()
	result.CellVars = b.getCellVars()

	result.FileName = b.getFileName()
	result.CoName = b.getName()
	result.LineNum = b.fileStream.ReadInt()
	result.NoTable = b.getNoTable()

	return result
}

func (b *BinaryFileParser) getByteCodes() *pyType.HiString {
	str := string(b.fileStream.Read())
	if (str != "s") {
		panic("There is no string")
	}

	return b.getString()
}

func (b *BinaryFileParser) getString() *pyType.HiString {
	length := b.fileStream.ReadInt()
	strValue := ""
	for i:=0; i < length; i++ {
		strValue += string(b.fileStream.Read())
	}

	return pyType.CreateHiString(strValue)
}

func (b *BinaryFileParser) getName() *pyType.HiString {
	ch := string(b.fileStream.Read())

	if (ch == "s") {
		return b.getString()
	} else if (ch == "t") {
		str := b.getString()
		b.stringTable = append(b.stringTable, str)
		return str
	} else if (ch == "R") {
		return b.stringTable[b.fileStream.ReadInt()]
	}
	return nil
}

func (b *BinaryFileParser) getFileName() *pyType.HiString {
	return b.getName()
}

func (b *BinaryFileParser) getNoTable() *pyType.HiString {
	ch := string(b.fileStream.Read())

	if (ch != "s" && ch != "t") {
		b.fileStream.Unread()
		return nil
	}

	return b.getString()
}

func (b *BinaryFileParser) getNames() []pyType.HiObject {
	if (string(b.fileStream.Read()) == "(") {
		return b.getTuple()
	}

	b.fileStream.Unread()
	return nil
}

func (b *BinaryFileParser) getVarNames() []pyType.HiObject {
	if (string(b.fileStream.Read()) == "(") {
		return b.getTuple()
	}

	b.fileStream.Unread()
	return nil
}

func (b *BinaryFileParser) getFreeVars() []pyType.HiObject {
	if (string(b.fileStream.Read()) == "(") {
		return b.getTuple()
	}

	b.fileStream.Unread()
	return nil
}

func (b *BinaryFileParser) getCellVars() []pyType.HiObject {
	if (string(b.fileStream.Read()) == "(") {
		return b.getTuple()
	}

	b.fileStream.Unread()
	return nil
}

func (b *BinaryFileParser) getConsts() []pyType.HiObject {
	if (string(b.fileStream.Read()) == "(") {
		return b.getTuple()
	}

	b.fileStream.Unread()
	return nil
}

func (b *BinaryFileParser) getTuple() []pyType.HiObject {
	length := b.fileStream.ReadInt()

	list := make([]pyType.HiObject, 0)
	for i := 0; i < length; i++ {
		objType := string(b.fileStream.Read())

		switch (objType) {
		case "c":
			fmt.Printf("get a code object\n")
			result := b.getCodeObject()
			list = append(list, result)
		case "i":
			result := pyType.CreateInteger(b.fileStream.ReadInt())
			list = append(list, result)
		case "N":
			list = append(list, nil)
		case "t":
			str := b.getString()
			list = append(list, str)
			b.stringTable = append(b.stringTable, str)
		case "s":
			str := b.getString()
			list = append(list, str)
		case "R":
			index := b.fileStream.ReadInt()
			str := b.stringTable[index]
			list = append(list, str)
		default:
			fmt.Printf("parser, unrecognized type: %s\n", objType)
		}
	}

	return list
}

func Parse(path string) *CodeObject {
	s := codeStream.CreateStringBuffer(path)
	parser := CreateBinaryFileParser(s)
	return parser.Parse()
}