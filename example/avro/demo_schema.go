// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     example.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

type DemoSchema struct {
	IntField int32 `json:"IntField"`

	DoubleField float64 `json:"DoubleField"`

	StringField string `json:"StringField"`

	BoolField bool `json:"BoolField"`

	BytesField []byte `json:"BytesField"`
}

const DemoSchemaAvroCRC64Fingerprint = "\xc4V\xa9\x04ʛf\xad"

func NewDemoSchema() *DemoSchema {
	return &DemoSchema{}
}

func DeserializeDemoSchema(r io.Reader) (*DemoSchema, error) {
	t := NewDemoSchema()
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

func DeserializeDemoSchemaFromSchema(r io.Reader, schema string) (*DemoSchema, error) {
	t := NewDemoSchema()

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

func writeDemoSchema(r *DemoSchema, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.IntField, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.DoubleField, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.StringField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.BoolField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBytes(r.BytesField, w)
	if err != nil {
		return err
	}
	return err
}

func (r *DemoSchema) Serialize(w io.Writer) error {
	return writeDemoSchema(r, w)
}

func (r *DemoSchema) Schema() string {
	return "{\"fields\":[{\"name\":\"IntField\",\"type\":\"int\"},{\"name\":\"DoubleField\",\"type\":\"double\"},{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"}],\"name\":\"DemoSchema\",\"type\":\"record\"}"
}

func (r *DemoSchema) SchemaName() string {
	return "DemoSchema"
}

func (_ *DemoSchema) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *DemoSchema) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *DemoSchema) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *DemoSchema) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *DemoSchema) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *DemoSchema) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *DemoSchema) SetString(v string)   { panic("Unsupported operation") }
func (_ *DemoSchema) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DemoSchema) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Int{Target: &r.IntField}
	case 1:
		return &types.Double{Target: &r.DoubleField}
	case 2:
		return &types.String{Target: &r.StringField}
	case 3:
		return &types.Boolean{Target: &r.BoolField}
	case 4:
		return &types.Bytes{Target: &r.BytesField}
	}
	panic("Unknown field index")
}

func (r *DemoSchema) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *DemoSchema) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *DemoSchema) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *DemoSchema) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *DemoSchema) Finalize()                        {}

func (_ *DemoSchema) AvroCRC64Fingerprint() []byte {
	return []byte(DemoSchemaAvroCRC64Fingerprint)
}
