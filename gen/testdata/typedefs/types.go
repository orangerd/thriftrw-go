// Code generated by thriftrw
// @generated

package typedefs

import (
	"errors"
	"fmt"
	"github.com/thriftrw/thriftrw-go/gen/testdata/enums"
	"github.com/thriftrw/thriftrw-go/gen/testdata/structs"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type _Set_Binary_ValueList [][]byte

func (v _Set_Binary_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		if x == nil {
			return fmt.Errorf("invalid set item: value is nil")
		}
		w, err := wire.NewValueBinary(x), error(nil)
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_Binary_ValueList) Size() int {
	return len(v)
}

func (_Set_Binary_ValueList) ValueType() wire.Type {
	return wire.TBinary
}

func (_Set_Binary_ValueList) Close() {
}

func _Set_Binary_Read(s wire.ValueList) ([][]byte, error) {
	if s.ValueType() != wire.TBinary {
		return nil, nil
	}
	o := make([][]byte, 0, s.Size())
	err := s.ForEach(func(x wire.Value) error {
		i, err := x.GetBinary(), error(nil)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	s.Close()
	return o, err
}

type BinarySet [][]byte

func (v BinarySet) ToWire() (wire.Value, error) {
	x := ([][]byte)(v)
	return wire.NewValueSet(_Set_Binary_ValueList(x)), error(nil)
}

func (v BinarySet) String() string {
	x := ([][]byte)(v)
	return fmt.Sprint(x)
}

func (v *BinarySet) FromWire(w wire.Value) error {
	x, err := _Set_Binary_Read(w.GetSet())
	*v = (BinarySet)(x)
	return err
}

type _Map_Edge_Edge_MapItemList []struct {
	Key   *structs.Edge
	Value *structs.Edge
}

func (m _Map_Edge_Edge_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for _, i := range m {
		k := i.Key
		v := i.Value
		if k == nil {
			return fmt.Errorf("invalid map key: value is nil")
		}
		if v == nil {
			return fmt.Errorf("invalid [%v]: value is nil", k)
		}
		kw, err := k.ToWire()
		if err != nil {
			return err
		}
		vw, err := v.ToWire()
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_Edge_Edge_MapItemList) Size() int {
	return len(m)
}

func (_Map_Edge_Edge_MapItemList) KeyType() wire.Type {
	return wire.TStruct
}

func (_Map_Edge_Edge_MapItemList) ValueType() wire.Type {
	return wire.TStruct
}

func (_Map_Edge_Edge_MapItemList) Close() {
}

func _Edge_Read(w wire.Value) (*structs.Edge, error) {
	var v structs.Edge
	err := v.FromWire(w)
	return &v, err
}

func _Map_Edge_Edge_Read(m wire.MapItemList) ([]struct {
	Key   *structs.Edge
	Value *structs.Edge
}, error) {
	if m.KeyType() != wire.TStruct {
		return nil, nil
	}
	if m.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make([]struct {
		Key   *structs.Edge
		Value *structs.Edge
	}, 0, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := _Edge_Read(x.Key)
		if err != nil {
			return err
		}
		v, err := _Edge_Read(x.Value)
		if err != nil {
			return err
		}
		o = append(o, struct {
			Key   *structs.Edge
			Value *structs.Edge
		}{k, v})
		return nil
	})
	m.Close()
	return o, err
}

type EdgeMap []struct {
	Key   *structs.Edge
	Value *structs.Edge
}

func (v EdgeMap) ToWire() (wire.Value, error) {
	x := ([]struct {
		Key   *structs.Edge
		Value *structs.Edge
	})(v)
	return wire.NewValueMap(_Map_Edge_Edge_MapItemList(x)), error(nil)
}

func (v EdgeMap) String() string {
	x := ([]struct {
		Key   *structs.Edge
		Value *structs.Edge
	})(v)
	return fmt.Sprint(x)
}

func (v *EdgeMap) FromWire(w wire.Value) error {
	x, err := _Map_Edge_Edge_Read(w.GetMap())
	*v = (EdgeMap)(x)
	return err
}

type Event struct {
	UUID *UUID      `json:"uuid"`
	Time *Timestamp `json:"time,omitempty"`
}

func (v *Event) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.UUID == nil {
		return w, errors.New("field UUID of Event is required")
	}
	w, err = v.UUID.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Time != nil {
		w, err = v.Time.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _UUID_Read(w wire.Value) (*UUID, error) {
	var x UUID
	err := x.FromWire(w)
	return &x, err
}

func _Timestamp_Read(w wire.Value) (Timestamp, error) {
	var x Timestamp
	err := x.FromWire(w)
	return x, err
}

func (v *Event) FromWire(w wire.Value) error {
	var err error
	uuidIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.UUID, err = _UUID_Read(field.Value)
				if err != nil {
					return err
				}
				uuidIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI64 {
				var x Timestamp
				x, err = _Timestamp_Read(field.Value)
				v.Time = &x
				if err != nil {
					return err
				}
			}
		}
	}
	if !uuidIsSet {
		return errors.New("field UUID of Event is required")
	}
	return nil
}

func (v *Event) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("UUID: %v", v.UUID)
	i++
	if v.Time != nil {
		fields[i] = fmt.Sprintf("Time: %v", *(v.Time))
		i++
	}
	return fmt.Sprintf("Event{%v}", strings.Join(fields[:i], ", "))
}

type _List_Event_ValueList []*Event

func (v _List_Event_ValueList) ForEach(f func(wire.Value) error) error {
	for i, x := range v {
		if x == nil {
			return fmt.Errorf("invalid [%v]: value is nil", i)
		}
		w, err := x.ToWire()
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_Event_ValueList) Size() int {
	return len(v)
}

func (_List_Event_ValueList) ValueType() wire.Type {
	return wire.TStruct
}

func (_List_Event_ValueList) Close() {
}

func _Event_Read(w wire.Value) (*Event, error) {
	var v Event
	err := v.FromWire(w)
	return &v, err
}

func _List_Event_Read(l wire.ValueList) ([]*Event, error) {
	if l.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make([]*Event, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := _Event_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

type EventGroup []*Event

func (v EventGroup) ToWire() (wire.Value, error) {
	x := ([]*Event)(v)
	return wire.NewValueList(_List_Event_ValueList(x)), error(nil)
}

func (v EventGroup) String() string {
	x := ([]*Event)(v)
	return fmt.Sprint(x)
}

func (v *EventGroup) FromWire(w wire.Value) error {
	x, err := _List_Event_Read(w.GetList())
	*v = (EventGroup)(x)
	return err
}

type _Set_Frame_ValueList []*structs.Frame

func (v _Set_Frame_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		if x == nil {
			return fmt.Errorf("invalid set item: value is nil")
		}
		w, err := x.ToWire()
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_Frame_ValueList) Size() int {
	return len(v)
}

func (_Set_Frame_ValueList) ValueType() wire.Type {
	return wire.TStruct
}

func (_Set_Frame_ValueList) Close() {
}

func _Frame_Read(w wire.Value) (*structs.Frame, error) {
	var v structs.Frame
	err := v.FromWire(w)
	return &v, err
}

func _Set_Frame_Read(s wire.ValueList) ([]*structs.Frame, error) {
	if s.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make([]*structs.Frame, 0, s.Size())
	err := s.ForEach(func(x wire.Value) error {
		i, err := _Frame_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	s.Close()
	return o, err
}

type FrameGroup []*structs.Frame

func (v FrameGroup) ToWire() (wire.Value, error) {
	x := ([]*structs.Frame)(v)
	return wire.NewValueSet(_Set_Frame_ValueList(x)), error(nil)
}

func (v FrameGroup) String() string {
	x := ([]*structs.Frame)(v)
	return fmt.Sprint(x)
}

func (v *FrameGroup) FromWire(w wire.Value) error {
	x, err := _Set_Frame_Read(w.GetSet())
	*v = (FrameGroup)(x)
	return err
}

func _EnumWithValues_Read(w wire.Value) (enums.EnumWithValues, error) {
	var v enums.EnumWithValues
	err := v.FromWire(w)
	return v, err
}

type MyEnum enums.EnumWithValues

func (v MyEnum) ToWire() (wire.Value, error) {
	x := (enums.EnumWithValues)(v)
	return x.ToWire()
}

func (v MyEnum) String() string {
	x := (enums.EnumWithValues)(v)
	return fmt.Sprint(x)
}

func (v *MyEnum) FromWire(w wire.Value) error {
	x, err := _EnumWithValues_Read(w)
	*v = (MyEnum)(x)
	return err
}

type Pdf []byte

func (v Pdf) ToWire() (wire.Value, error) {
	x := ([]byte)(v)
	return wire.NewValueBinary(x), error(nil)
}

func (v Pdf) String() string {
	x := ([]byte)(v)
	return fmt.Sprint(x)
}

func (v *Pdf) FromWire(w wire.Value) error {
	x, err := w.GetBinary(), error(nil)
	*v = (Pdf)(x)
	return err
}

type _Map_Point_Point_MapItemList []struct {
	Key   *structs.Point
	Value *structs.Point
}

func (m _Map_Point_Point_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for _, i := range m {
		k := i.Key
		v := i.Value
		if k == nil {
			return fmt.Errorf("invalid map key: value is nil")
		}
		if v == nil {
			return fmt.Errorf("invalid [%v]: value is nil", k)
		}
		kw, err := k.ToWire()
		if err != nil {
			return err
		}
		vw, err := v.ToWire()
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_Point_Point_MapItemList) Size() int {
	return len(m)
}

func (_Map_Point_Point_MapItemList) KeyType() wire.Type {
	return wire.TStruct
}

func (_Map_Point_Point_MapItemList) ValueType() wire.Type {
	return wire.TStruct
}

func (_Map_Point_Point_MapItemList) Close() {
}

func _Point_Read(w wire.Value) (*structs.Point, error) {
	var v structs.Point
	err := v.FromWire(w)
	return &v, err
}

func _Map_Point_Point_Read(m wire.MapItemList) ([]struct {
	Key   *structs.Point
	Value *structs.Point
}, error) {
	if m.KeyType() != wire.TStruct {
		return nil, nil
	}
	if m.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make([]struct {
		Key   *structs.Point
		Value *structs.Point
	}, 0, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := _Point_Read(x.Key)
		if err != nil {
			return err
		}
		v, err := _Point_Read(x.Value)
		if err != nil {
			return err
		}
		o = append(o, struct {
			Key   *structs.Point
			Value *structs.Point
		}{k, v})
		return nil
	})
	m.Close()
	return o, err
}

type PointMap []struct {
	Key   *structs.Point
	Value *structs.Point
}

func (v PointMap) ToWire() (wire.Value, error) {
	x := ([]struct {
		Key   *structs.Point
		Value *structs.Point
	})(v)
	return wire.NewValueMap(_Map_Point_Point_MapItemList(x)), error(nil)
}

func (v PointMap) String() string {
	x := ([]struct {
		Key   *structs.Point
		Value *structs.Point
	})(v)
	return fmt.Sprint(x)
}

func (v *PointMap) FromWire(w wire.Value) error {
	x, err := _Map_Point_Point_Read(w.GetMap())
	*v = (PointMap)(x)
	return err
}

type State string

func (v State) ToWire() (wire.Value, error) {
	x := (string)(v)
	return wire.NewValueString(x), error(nil)
}

func (v State) String() string {
	x := (string)(v)
	return fmt.Sprint(x)
}

func (v *State) FromWire(w wire.Value) error {
	x, err := w.GetString(), error(nil)
	*v = (State)(x)
	return err
}

type Timestamp int64

func (v Timestamp) ToWire() (wire.Value, error) {
	x := (int64)(v)
	return wire.NewValueI64(x), error(nil)
}

func (v Timestamp) String() string {
	x := (int64)(v)
	return fmt.Sprint(x)
}

func (v *Timestamp) FromWire(w wire.Value) error {
	x, err := w.GetI64(), error(nil)
	*v = (Timestamp)(x)
	return err
}

type Transition struct {
	From   State      `json:"from"`
	To     State      `json:"to"`
	Events EventGroup `json:"events"`
}

func (v *Transition) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = v.From.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = v.To.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	if v.Events != nil {
		w, err = v.Events.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _State_Read(w wire.Value) (State, error) {
	var x State
	err := x.FromWire(w)
	return x, err
}

func _EventGroup_Read(w wire.Value) (EventGroup, error) {
	var x EventGroup
	err := x.FromWire(w)
	return x, err
}

func (v *Transition) FromWire(w wire.Value) error {
	var err error
	fromIsSet := false
	toIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.From, err = _State_Read(field.Value)
				if err != nil {
					return err
				}
				fromIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.To, err = _State_Read(field.Value)
				if err != nil {
					return err
				}
				toIsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TList {
				v.Events, err = _EventGroup_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	if !fromIsSet {
		return errors.New("field From of Transition is required")
	}
	if !toIsSet {
		return errors.New("field To of Transition is required")
	}
	return nil
}

func (v *Transition) String() string {
	var fields [3]string
	i := 0
	fields[i] = fmt.Sprintf("From: %v", v.From)
	i++
	fields[i] = fmt.Sprintf("To: %v", v.To)
	i++
	if v.Events != nil {
		fields[i] = fmt.Sprintf("Events: %v", v.Events)
		i++
	}
	return fmt.Sprintf("Transition{%v}", strings.Join(fields[:i], ", "))
}

type UUID I128

func (v *UUID) ToWire() (wire.Value, error) {
	x := (*I128)(v)
	return x.ToWire()
}

func (v *UUID) String() string {
	x := (*I128)(v)
	return fmt.Sprint(x)
}

func (v *UUID) FromWire(w wire.Value) error {
	return (*I128)(v).FromWire(w)
}

type I128 struct {
	High int64 `json:"high"`
	Low  int64 `json:"low"`
}

func (v *I128) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueI64(v.High), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueI64(v.Low), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *I128) FromWire(w wire.Value) error {
	var err error
	highIsSet := false
	lowIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI64 {
				v.High, err = field.Value.GetI64(), error(nil)
				if err != nil {
					return err
				}
				highIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI64 {
				v.Low, err = field.Value.GetI64(), error(nil)
				if err != nil {
					return err
				}
				lowIsSet = true
			}
		}
	}
	if !highIsSet {
		return errors.New("field High of I128 is required")
	}
	if !lowIsSet {
		return errors.New("field Low of I128 is required")
	}
	return nil
}

func (v *I128) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("High: %v", v.High)
	i++
	fields[i] = fmt.Sprintf("Low: %v", v.Low)
	i++
	return fmt.Sprintf("I128{%v}", strings.Join(fields[:i], ", "))
}
