// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     arrays.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

type MapTestRecord struct {
	IntField map[string]*UnionNullInt `json:"IntField"`
}

const MapTestRecordAvroCRC64Fingerprint = "\xf7\xdb\x00\xb2n\xa8u\xbf"

func NewMapTestRecord() *MapTestRecord {
	return &MapTestRecord{}
}

func DeserializeMapTestRecord(r io.Reader) (*MapTestRecord, error) {
	t := NewMapTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeMapTestRecordFromSchema(r io.Reader, schema string) (*MapTestRecord, error) {
	t := NewMapTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeMapTestRecord(r *MapTestRecord, w io.Writer) error {
	var err error
	err = writeMapUnionNullInt(r.IntField, w)
	if err != nil {
		return err
	}
	return err
}

func (r *MapTestRecord) Serialize(w io.Writer) error {
	return writeMapTestRecord(r, w)
}

func (r *MapTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"IntField\",\"type\":{\"type\":\"map\",\"values\":[\"null\",\"int\"]}}],\"name\":\"MapTestRecord\",\"type\":\"record\"}"
}

func (r *MapTestRecord) SchemaName() string {
	return "MapTestRecord"
}

func (_ *MapTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *MapTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *MapTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *MapTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *MapTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *MapTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *MapTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *MapTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MapTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		r.IntField = make(map[string]*UnionNullInt)

		return &MapUnionNullIntWrapper{Target: &r.IntField}
	}
	panic("Unknown field index")
}

func (r *MapTestRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *MapTestRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *MapTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *MapTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *MapTestRecord) Finalize()                        {}

func (_ *MapTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(MapTestRecordAvroCRC64Fingerprint)
}
