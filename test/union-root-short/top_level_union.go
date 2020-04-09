// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     union.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)

type TopLevelUnionTypeEnum int

const (
	TopLevelUnionTypeEnumIp_address TopLevelUnionTypeEnum = 0

	TopLevelUnionTypeEnumEvent TopLevelUnionTypeEnum = 1
)

type TopLevelUnion struct {
	Ip_address Ip_address
	Event      *Event
	UnionType  TopLevelUnionTypeEnum
}

func writeTopLevelUnion(r *TopLevelUnion, w io.Writer) error {

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case TopLevelUnionTypeEnumIp_address:
		return writeIp_address(r.Ip_address, w)
	case TopLevelUnionTypeEnumEvent:
		return writeEvent(r.Event, w)
	}
	return fmt.Errorf("invalid value for *TopLevelUnion")
}

func NewTopLevelUnion() *TopLevelUnion {
	return &TopLevelUnion{}
}

func (_ *TopLevelUnion) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *TopLevelUnion) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *TopLevelUnion) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *TopLevelUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *TopLevelUnion) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *TopLevelUnion) SetString(v string)  { panic("Unsupported operation") }
func (r *TopLevelUnion) SetLong(v int64) {
	r.UnionType = (TopLevelUnionTypeEnum)(v)
}
func (r *TopLevelUnion) Get(i int) types.Field {
	switch i {
	case 0:
		return &Ip_addressWrapper{Target: (&r.Ip_address)}
	case 1:
		r.Event = NewEvent()
		return r.Event
	}
	panic("Unknown field index")
}
func (_ *TopLevelUnion) NullField(i int)                  { panic("Unsupported operation") }
func (_ *TopLevelUnion) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *TopLevelUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *TopLevelUnion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *TopLevelUnion) Finalize()                        {}

func (r *TopLevelUnion) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case TopLevelUnionTypeEnumIp_address:
		return json.Marshal(map[string]interface{}{"ip_address": r.Ip_address})
	case TopLevelUnionTypeEnumEvent:
		return json.Marshal(map[string]interface{}{"event": r.Event})
	}
	return nil, fmt.Errorf("invalid value for *TopLevelUnion")
}

func (r *TopLevelUnion) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["ip_address"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.Ip_address)
	}
	if value, ok := fields["event"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Event)
	}
	return fmt.Errorf("invalid value for *TopLevelUnion")
}
