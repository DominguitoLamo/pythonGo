package pyType

import (
	"fmt"
	"testing"
)

func callPrint(obj HiObject) {
	switch v := obj.(type) {
	case HiInterger:
		fmt.Printf("This is integer %d", HiInterger(v).value)
	case HiString:
		fmt.Printf("This is String %s", HiString(v).Value)
	default:
		fmt.Println("This is nothing")
	}
}

func TestPointerReceiver(t *testing.T) {
	foo := &HiInterger{value: 13}
	callPrint(*foo)

	var bar HiInterger
	bar.value = 13
	callPrint(bar)
}