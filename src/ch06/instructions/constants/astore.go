package constants

import (
	"learn/ch06/instructions/base"
	"learn/ch06/rtda"
)

type ASTORE struct{ base.Index8Instruction }
type ASTORE_0 struct{ base.NoOperandsInstruction }
type ASTORE_1 struct{ base.NoOperandsInstruction }
type ASTORE_2 struct{ base.NoOperandsInstruction }
type ASTORE_3 struct{ base.NoOperandsInstruction }

func _astore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}
func (self *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, uint(self.Index))
}
func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}
func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}
