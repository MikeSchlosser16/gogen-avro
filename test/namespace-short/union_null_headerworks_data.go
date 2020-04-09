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

type UnionNullHeaderworksDataTypeEnum int

const (
	UnionNullHeaderworksDataTypeEnumHeaderworksData UnionNullHeaderworksDataTypeEnum = 1
)

type UnionNullHeaderworksData struct {
	Null            *types.NullVal
	HeaderworksData *HeaderworksData
	UnionType       UnionNullHeaderworksDataTypeEnum
}

func writeUnionNullHeaderworksData(r *UnionNullHeaderworksData, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullHeaderworksDataTypeEnumHeaderworksData:
		return writeHeaderworksData(r.HeaderworksData, w)
	}
	return fmt.Errorf("invalid value for *UnionNullHeaderworksData")
}

func NewUnionNullHeaderworksData() *UnionNullHeaderworksData {
	return &UnionNullHeaderworksData{}
}

func (_ *UnionNullHeaderworksData) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionNullHeaderworksData) SetLong(v int64) {
	r.UnionType = (UnionNullHeaderworksDataTypeEnum)(v)
}
func (r *UnionNullHeaderworksData) Get(i int) types.Field {
	switch i {
	case 0:
		return r.Null
	case 1:
		r.HeaderworksData = NewHeaderworksData()
		return r.HeaderworksData
	}
	panic("Unknown field index")
}
func (_ *UnionNullHeaderworksData) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) Finalize()                        {}

func (r *UnionNullHeaderworksData) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionNullHeaderworksDataTypeEnumHeaderworksData:
		return json.Marshal(map[string]interface{}{"Data": r.HeaderworksData})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullHeaderworksData")
}

func (r *UnionNullHeaderworksData) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["Data"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.HeaderworksData)
	}
	return fmt.Errorf("invalid value for *UnionNullHeaderworksData")
}
