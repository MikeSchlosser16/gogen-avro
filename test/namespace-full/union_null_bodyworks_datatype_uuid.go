// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)

type UnionNullBodyworksDatatypeUUIDTypeEnum int

const (
	UnionNullBodyworksDatatypeUUIDTypeEnumBodyworksDatatypeUUID UnionNullBodyworksDatatypeUUIDTypeEnum = 1
)

type UnionNullBodyworksDatatypeUUID struct {
	Null                  *types.NullVal
	BodyworksDatatypeUUID *BodyworksDatatypeUUID
	UnionType             UnionNullBodyworksDatatypeUUIDTypeEnum
}

func writeUnionNullBodyworksDatatypeUUID(r *UnionNullBodyworksDatatypeUUID, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullBodyworksDatatypeUUIDTypeEnumBodyworksDatatypeUUID:
		return writeBodyworksDatatypeUUID(r.BodyworksDatatypeUUID, w)
	}
	return fmt.Errorf("invalid value for *UnionNullBodyworksDatatypeUUID")
}

func NewUnionNullBodyworksDatatypeUUID() *UnionNullBodyworksDatatypeUUID {
	return &UnionNullBodyworksDatatypeUUID{}
}

func (_ *UnionNullBodyworksDatatypeUUID) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionNullBodyworksDatatypeUUID) SetLong(v int64) {
	r.UnionType = (UnionNullBodyworksDatatypeUUIDTypeEnum)(v)
}
func (r *UnionNullBodyworksDatatypeUUID) Get(i int) types.Field {
	switch i {
	case 0:
		return r.Null
	case 1:
		r.BodyworksDatatypeUUID = NewBodyworksDatatypeUUID()
		return r.BodyworksDatatypeUUID
	}
	panic("Unknown field index")
}
func (_ *UnionNullBodyworksDatatypeUUID) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullBodyworksDatatypeUUID) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) Finalize()                {}

func (r *UnionNullBodyworksDatatypeUUID) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionNullBodyworksDatatypeUUIDTypeEnumBodyworksDatatypeUUID:
		return json.Marshal(map[string]interface{}{"UUID": r.BodyworksDatatypeUUID})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullBodyworksDatatypeUUID")
}

func (r *UnionNullBodyworksDatatypeUUID) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["UUID"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.BodyworksDatatypeUUID)
	}
	return fmt.Errorf("invalid value for *UnionNullBodyworksDatatypeUUID")
}
