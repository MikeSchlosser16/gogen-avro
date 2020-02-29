package types

type Bytes []byte

func (b *Bytes) SetBoolean(v bool) {
	panic("Unable to assign bytes to bytes field")
}

func (b *Bytes) SetInt(v int32) {
	panic("Unable to assign int to bytes field")
}

func (b *Bytes) SetLong(v int64) {
	panic("Unable to assign long to bytes field")
}

func (b *Bytes) SetFloat(v float32) {
	panic("Unable to assign float to bytes field")
}

func (b *Bytes) SetDouble(v float64) {
	panic("Unable to assign double to bytes field")
}

func (b *Bytes) SetUnionElem(v int64) {
	panic("Unable to assign union elem to bytes field")
}

func (b *Bytes) SetBytes(v []byte) {
	*b = v
}

func (b *Bytes) SetString(v string) {
	*b = []byte(v)
}

func (b *Bytes) Get(i int) Field {
	panic("Unable to get field from bytes field")
}

func (b *Bytes) SetDefault(i int) {
	panic("Unable to set default on bytes field")
}

func (b *Bytes) Clear(i int) {
	panic("Unable to clear field from bytes field")
}

func (b *Bytes) AppendMap(key string) Field {
	panic("Unable to append map key to from bytes field")
}

func (b *Bytes) AppendArray() Field {
	panic("Unable to append array element to from bytes field")
}

func (b *Bytes) ClearMap(key string) {
	panic("Unable to clear map key from bytes field")
}

func (b *Bytes) Finalize() {}
