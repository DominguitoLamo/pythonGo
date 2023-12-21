package code

import (
	"fmt"
	"chonlam.com/pythonGo/codeStream"
)

type BinaryFileParser struct {
	fileStream *codeStream.StringBuffer;
}

func CreateBinaryFileParser(f *codeStream.StringBuffer) *BinaryFileParser {
	b := new(BinaryFileParser)
	b.fileStream = f
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
		return getCodeObject()
	}

	return nil
}

func getCodeObject() *CodeObject {
	result := new(CodeObject)
	return result
}

func Parse(path string) {
	s := codeStream.CreateStringBuffer(path)
	parser := CreateBinaryFileParser(s)
	parser.Parse()
}