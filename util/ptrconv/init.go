package ptrconv

import "github.com/bearaujus/bgdk/util/ptrconv/mock"

// PtrConv is an interface primarily used by Instance.
type PtrConv interface {
	// ConvBool convert bool object to ptr bool (*bool)
	ConvBool(v bool) *bool

	// ConvByte convert byte object to ptr byte (*byte)
	ConvByte(v byte) *byte

	// ConvComplex64 convert complex64 object to ptr complex64 (*complex64)
	ConvComplex64(v complex64) *complex64

	// ConvComplex128 convert complex128 object to ptr complex128 (*complex128)
	ConvComplex128(v complex128) *complex128

	// ConvFloat32 convert float32 object to ptr float32 (*float32)
	ConvFloat32(v float32) *float32

	// ConvFloat64 convert float64 object to ptr float64 (*float64)
	ConvFloat64(v float64) *float64

	// ConvInt convert int object to ptr int (*int)
	ConvInt(v int) *int

	// ConvInt8 convert int8 object to ptr int8 (*int8)
	ConvInt8(v int8) *int8

	// ConvInt16 convert int16 object to ptr int16 (*int16)
	ConvInt16(v int16) *int16

	// ConvInt32 convert int32 object to ptr int32 (*int32)
	ConvInt32(v int32) *int32

	// ConvInt64 convert int64 object to ptr int64 (*int64)
	ConvInt64(v int64) *int64

	// ConvRune convert rune object to ptr rune (*rune)
	ConvRune(v rune) *rune

	// ConvString convert string object to ptr string (*string)
	ConvString(v string) *string

	// ConvUint convert uint object to ptr uint (*uint)
	ConvUint(v uint) *uint

	// ConvUint8 convert uint8 object to ptr uint8 (*uint8)
	ConvUint8(v uint8) *uint8

	// ConvUint16 convert uint16 object to ptr uint16 (*uint16)
	ConvUint16(v uint16) *uint16

	// ConvUint32 convert uint32 object to ptr uint32 (*uint32)
	ConvUint32(v uint32) *uint32

	// ConvUint64 convert uint64 object to ptr uint64 (*uint64)
	ConvUint64(v uint64) *uint64

	// ConvPtrBool convert ptr bool (*bool) to bool object
	//
	// Note: it is safe when v is nil. It will create a new bool object and return it
	ConvPtrBool(v *bool) bool

	// ConvPtrByte convert ptr byte (*byte) to byte object
	//
	// Note: it is safe when v is nil. It will create a new byte object and return it
	ConvPtrByte(v *byte) byte

	// ConvPtrComplex64 convert ptr complex64 (*complex64) to complex64 object
	//
	// Note: it is safe when v is nil. It will create a new complex64 object and return it
	ConvPtrComplex64(v *complex64) complex64

	// ConvPtrComplex128 convert ptr complex128 (*complex128) to complex128 object
	//
	// Note: it is safe when v is nil. It will create a new complex128 object and return it
	ConvPtrComplex128(v *complex128) complex128

	// ConvPtrFloat32 convert ptr float32 (*float32) to float32 object
	//
	// Note: it is safe when v is nil. It will create a new float32 object and return it
	ConvPtrFloat32(v *float32) float32

	// ConvPtrFloat64 convert ptr float64 (*float64) to float64 object
	//
	// Note: it is safe when v is nil. It will create a new float64 object and return it
	ConvPtrFloat64(v *float64) float64

	// ConvPtrInt convert ptr int (*int) to int object
	//
	// Note: it is safe when v is nil. It will create a new int object and return it
	ConvPtrInt(v *int) int

	// ConvPtrInt8 convert ptr int8 (*int8) to int8 object
	//
	// Note: it is safe when v is nil. It will create a new int8 object and return it
	ConvPtrInt8(v *int8) int8

	// ConvPtrInt16 convert ptr int16 (*int16) to int16 object
	//
	// Note: it is safe when v is nil. It will create a new int16 object and return it
	ConvPtrInt16(v *int16) int16

	// ConvPtrInt32 convert ptr int32 (*int32) to int32 object
	//
	// Note: it is safe when v is nil. It will create a new int32 object and return it
	ConvPtrInt32(v *int32) int32

	// ConvPtrInt64 convert ptr int64 (*int64) to int64 object
	//
	// Note: it is safe when v is nil. It will create a new int64 object and return it
	ConvPtrInt64(v *int64) int64

	// ConvPtrRune convert ptr rune (*rune) to rune object
	//
	// Note: it is safe when v is nil. It will create a new rune object and return it
	ConvPtrRune(v *rune) rune

	// ConvPtrString convert ptr string (*string) to string object
	//
	// Note: it is safe when v is nil. It will create a new string object and return it
	ConvPtrString(v *string) string

	// ConvPtrUint convert ptr uint (*uint) to uint object
	//
	// Note: it is safe when v is nil. It will create a new uint object and return it
	ConvPtrUint(v *uint) uint

	// ConvPtrUint8 convert ptr uint8 (*uint8) to uint8 object
	//
	// Note: it is safe when v is nil. It will create a new uint8 object and return it
	ConvPtrUint8(v *uint8) uint8

	// ConvPtrUint16 convert ptr uint16 (*uint16) to uint16 object
	//
	// Note: it is safe when v is nil. It will create a new uint16 object and return it
	ConvPtrUint16(v *uint16) uint16

	// ConvPtrUint32 convert ptr uint32 (*uint32) to uint32 object
	//
	// Note: it is safe when v is nil. It will create a new uint32 object and return it
	ConvPtrUint32(v *uint32) uint32

	// ConvPtrUint64 convert ptr uint64 (*uint64) to uint64 object
	//
	// Note: it is safe when v is nil. It will create a new uint64 object and return it
	ConvPtrUint64(v *uint64) uint64
}

// ptrConv is a struct to implement the PtrConv interface.
type ptrConv struct{}

// instance hold the PtrConv interface.
var instance PtrConv

// Instance will return the PtrConv interface.
//
// Note: first time calling this function will create a new instance
// to implement the PtrConv interface (except if you call InitTestMode).
// The next time this function is called, it will use the previously created instance.
func Instance() PtrConv {
	if instance == nil {
		instance = &ptrConv{}
	}

	return instance
}

// InitTestMode will set the PtrConv instance to mockInstance.
//
// Note: this function for testing purposes only. After executing InitTestMode,
// whenever your unit tests execute Instance it will call mockInstance.
func InitTestMode(mockInstance *mock.MockPtrConv) {
	instance = mockInstance
}
