package constants

import (
	"learn/ch09/instructions/base"
	"learn/ch09/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// 什么也不用做
}
