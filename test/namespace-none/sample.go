// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

// GoGen test
type Sample struct {
	// Core data information required for any event
	Header *UnionNullData `json:"header"`
	// Core data information required for any event
	Body *UnionNullData `json:"body"`
}

const SampleAvroCRC64Fingerprint = "\xdf}\x93 \x19f\x18\n"

func NewSample() *Sample {
	return &Sample{}
}

func DeserializeSample(r io.Reader) (*Sample, error) {
	t := NewSample()
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

func DeserializeSampleFromSchema(r io.Reader, schema string) (*Sample, error) {
	t := NewSample()

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

func writeSample(r *Sample, w io.Writer) error {
	var err error
	err = writeUnionNullData(r.Header, w)
	if err != nil {
		return err
	}
	err = writeUnionNullData(r.Body, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Sample) Serialize(w io.Writer) error {
	return writeSample(r, w)
}

func (r *Sample) Schema() string {
	return "{\"doc\":\"GoGen test\",\"fields\":[{\"default\":null,\"doc\":\"Core data information required for any event\",\"name\":\"header\",\"type\":[\"null\",{\"doc\":\"Common information related to the event which must be included in any clean event\",\"fields\":[{\"default\":null,\"doc\":\"Unique identifier for the event used for de-duplication and tracing.\",\"name\":\"uuid\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Fully qualified name of the host that generated the event that generated the data.\",\"name\":\"hostname\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"Trace information not redundant with this object\",\"name\":\"trace\",\"type\":[\"null\",{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",\"headerworks.datatype.UUID\"]}],\"name\":\"Trace\",\"type\":\"record\"}]}],\"name\":\"Data\",\"namespace\":\"headerworks\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Core data information required for any event\",\"name\":\"body\",\"type\":[\"null\",{\"doc\":\"Common information related to the event which must be included in any clean event\",\"fields\":[{\"default\":null,\"doc\":\"Unique identifier for the event used for de-duplication and tracing.\",\"name\":\"uuid\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"bodyworks.datatype\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Fully qualified name of the host that generated the event that generated the data.\",\"name\":\"hostname\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"Trace information not redundant with this object\",\"name\":\"trace\",\"type\":[\"null\",{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",\"headerworks.datatype.UUID\"]}],\"name\":\"Trace\",\"type\":\"record\"}]}],\"name\":\"Data\",\"namespace\":\"bodyworks\",\"type\":\"record\"}]}],\"name\":\"com.avro.test.sample\",\"type\":\"record\"}"
}

func (r *Sample) SchemaName() string {
	return "com.avro.test.sample"
}

func (_ *Sample) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Sample) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Sample) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Sample) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Sample) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Sample) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Sample) SetString(v string)   { panic("Unsupported operation") }
func (_ *Sample) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Sample) Get(i int) types.Field {
	switch i {
	case 0:
		r.Header = NewUnionNullData()

		return r.Header
	case 1:
		r.Body = NewUnionNullData()

		return r.Body
	}
	panic("Unknown field index")
}

func (r *Sample) SetDefault(i int) {
	switch i {
	case 0:
		r.Header = nil
		return
	case 1:
		r.Body = nil
		return
	}
	panic("Unknown field index")
}

func (r *Sample) NullField(i int) {
	switch i {
	case 0:
		r.Header = nil
		return
	case 1:
		r.Body = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ *Sample) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Sample) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Sample) Finalize()                        {}

func (_ *Sample) AvroCRC64Fingerprint() []byte {
	return []byte(SampleAvroCRC64Fingerprint)
}
