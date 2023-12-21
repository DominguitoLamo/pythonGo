package codeStream

import (
	"os"
	"strconv"
)

type StringBuffer struct {
	content []byte;
	index int;
}

func CreateStringBuffer(codePath string) *StringBuffer {
	b, err := os.ReadFile(codePath)
	if (err != nil) {
		panic("No Existing File!")
	}

	buffer := new(StringBuffer)
	buffer.content = b
	buffer.index = 0
	return buffer
}

func (s *StringBuffer) Read() byte {
	if (s.index < len(s.content)) {
		result := s.content[s.index]
		s.index++
		return result
	} else {
		s.index = 0
		result := s.content[s.index]
		s.index++
		return result
	}
}

func (s *StringBuffer) Unread() {
	s.index--
}

func str2int(s string) int {
	i, err := strconv.Atoi(s)
	if (err != nil) {
		panic("Read non-number in the sequence")
	}
	return i
}

func (s *StringBuffer) ReadInt() int {
	b1 := int(s.Read())
	b2 := int(s.Read()) << 8 
	b3 := int(s.Read()) << 16
	b4 := int(s.Read()) << 24

	return b4 | b3 | b2 | b1
}