package ptrconv

func (p *ptrConv) ConvBool(v bool) *bool {
	// reference the bool object to ptr bool
	return &v
}

func (p *ptrConv) ConvByte(v byte) *byte {
	// reference the byte object to ptr byte
	return &v
}

func (p *ptrConv) ConvComplex64(v complex64) *complex64 {
	// reference the complex64 object to ptr complex64
	return &v
}

func (p *ptrConv) ConvComplex128(v complex128) *complex128 {
	// reference the complex128 object to ptr complex128
	return &v
}

func (p *ptrConv) ConvFloat32(v float32) *float32 {
	// reference the float32 object to ptr float32
	return &v
}

func (p *ptrConv) ConvFloat64(v float64) *float64 {
	// reference the float64 object to ptr float64
	return &v
}

func (p *ptrConv) ConvInt(v int) *int {
	// reference the int object to ptr int
	return &v
}

func (p *ptrConv) ConvInt8(v int8) *int8 {
	// reference the int8 object to ptr int8
	return &v
}

func (p *ptrConv) ConvInt16(v int16) *int16 {
	// reference the int16 object to ptr int16
	return &v
}

func (p *ptrConv) ConvInt32(v int32) *int32 {
	// reference the int32 object to ptr int32
	return &v
}

func (p *ptrConv) ConvInt64(v int64) *int64 {
	// reference the int64 object to ptr int64
	return &v
}

func (p *ptrConv) ConvRune(v rune) *rune {
	// reference the rune object to ptr rune
	return &v
}

func (p *ptrConv) ConvString(v string) *string {
	// reference the string object to ptr string
	return &v
}

func (p *ptrConv) ConvUint(v uint) *uint {
	// reference the uint object to ptr uint
	return &v
}

func (p *ptrConv) ConvUint8(v uint8) *uint8 {
	// reference the uint8 object to ptr uint8
	return &v
}

func (p *ptrConv) ConvUint16(v uint16) *uint16 {
	// reference the uint16 object to ptr uint16
	return &v
}

func (p *ptrConv) ConvUint32(v uint32) *uint32 {
	// reference the uint32 object to ptr uint32
	return &v
}

func (p *ptrConv) ConvUint64(v uint64) *uint64 {
	// reference the uint64 object to ptr uint64
	return &v
}

func (p *ptrConv) ConvPtrBool(v *bool) bool {
	// if v is nil ptr bool, create new bool object
	if v == nil {
		var newV bool
		return newV
	}

	// if v is not nil ptr bool, dereference it
	return *v
}

func (p *ptrConv) ConvPtrByte(v *byte) byte {
	// if v is nil ptr byte, create new byte object
	if v == nil {
		var newV byte
		return newV
	}

	// if v is not nil ptr byte, dereference it
	return *v
}

func (p *ptrConv) ConvPtrComplex64(v *complex64) complex64 {
	// if v is nil ptr complex64, create new complex64 object
	if v == nil {
		var newV complex64
		return newV
	}

	// if v is not nil ptr complex64, dereference it
	return *v
}

func (p *ptrConv) ConvPtrComplex128(v *complex128) complex128 {
	// if v is nil ptr complex128, create new complex128 object
	if v == nil {
		var newV complex128
		return newV
	}

	// if v is not nil ptr complex128, dereference it
	return *v
}

func (p *ptrConv) ConvPtrFloat32(v *float32) float32 {
	// if v is nil ptr float32, create new float32 object
	if v == nil {
		var newV float32
		return newV
	}

	// if v is not nil ptr float32, dereference it
	return *v
}

func (p *ptrConv) ConvPtrFloat64(v *float64) float64 {
	// if v is nil ptr float64, create new float64 object
	if v == nil {
		var newV float64
		return newV
	}

	// if v is not nil ptr float64, dereference it
	return *v
}

func (p *ptrConv) ConvPtrInt(v *int) int {
	// if v is nil ptr int, create new int object
	if v == nil {
		var newV int
		return newV
	}

	// if v is not nil ptr int, dereference it
	return *v
}

func (p *ptrConv) ConvPtrInt8(v *int8) int8 {
	// if v is nil ptr int8, create new int8 object
	if v == nil {
		var newV int8
		return newV
	}

	// if v is not nil ptr int8, dereference it
	return *v
}

func (p *ptrConv) ConvPtrInt16(v *int16) int16 {
	// if v is nil ptr int16, create new int16 object
	if v == nil {
		var newV int16
		return newV
	}

	// if v is not nil ptr int16, dereference it
	return *v
}

func (p *ptrConv) ConvPtrInt32(v *int32) int32 {
	// if v is nil ptr int32, create new int32 object
	if v == nil {
		var newV int32
		return newV
	}

	// if v is not nil ptr int32, dereference it
	return *v
}

func (p *ptrConv) ConvPtrInt64(v *int64) int64 {
	// if v is nil ptr int64, create new int64 object
	if v == nil {
		var newV int64
		return newV
	}

	// if v is not nil ptr int64, dereference it
	return *v
}

func (p *ptrConv) ConvPtrRune(v *rune) rune {
	// if v is nil ptr rune, create new rune object
	if v == nil {
		var newV rune
		return newV
	}

	// if v is not nil ptr rune, dereference it
	return *v
}

func (p *ptrConv) ConvPtrString(v *string) string {
	// if v is nil ptr string, create new string object
	if v == nil {
		var newV string
		return newV
	}

	// if v is not nil ptr string, dereference it
	return *v
}

func (p *ptrConv) ConvPtrUint(v *uint) uint {
	// if v is nil ptr uint, create new uint object
	if v == nil {
		var newV uint
		return newV
	}

	// if v is not nil ptr uint, dereference it
	return *v
}

func (p *ptrConv) ConvPtrUint8(v *uint8) uint8 {
	// if v is nil ptr uint8, create new uint8 object
	if v == nil {
		var newV uint8
		return newV
	}

	// if v is not nil ptr uint8, dereference it
	return *v
}

func (p *ptrConv) ConvPtrUint16(v *uint16) uint16 {
	// if v is nil ptr uint16, create new uint16 object
	if v == nil {
		var newV uint16
		return newV
	}

	// if v is not nil ptr uint16, dereference it
	return *v
}

func (p *ptrConv) ConvPtrUint32(v *uint32) uint32 {
	// if v is nil ptr uint32, create new uint32 object
	if v == nil {
		var newV uint32
		return newV
	}

	// if v is not nil ptr uint32, dereference it
	return *v
}

func (p *ptrConv) ConvPtrUint64(v *uint64) uint64 {
	// if v is nil ptr uint64, create new uint64 object
	if v == nil {
		var newV uint64
		return newV
	}

	// if v is not nil ptr uint64, dereference it
	return *v
}
