// automatically generated, do not modify

package Test

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Person struct {
	_tab flatbuffers.Table
}

func GetRootAsPerson(buf []byte, offset flatbuffers.UOffsetT) *Person {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Person{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Person) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Person) Name() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *Person) Surname() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *Person) Age() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 18
}

func PersonStart(builder *flatbuffers.Builder) { builder.StartObject(3) }
func PersonAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0) }
func PersonAddSurname(builder *flatbuffers.Builder, surname flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(surname), 0) }
func PersonAddAge(builder *flatbuffers.Builder, age int16) { builder.PrependInt16Slot(2, age, 18) }
func PersonEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
