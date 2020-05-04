package heap

import "learn/ch09/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}
func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}
func (self *Field) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *Field) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}
func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) Descriptor() string {
	return self.descriptor
}

func (self *Field) Class() *Class {
	return self.class
}
