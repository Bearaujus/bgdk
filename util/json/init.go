package json

import "github.com/bearaujus/bgdk/util/json/mock"

// JSON is an interface primarily used by Instance.
type JSON interface {
	// Marshal is used to encode the data from v.
	//
	// Note: v can be a struct or a ptr struct (*struct).
	Marshal(v interface{}) ([]byte, error)

	// MarshalWrite is used to encode data from v and write the encoded data to destPath.
	// if pretty is true, it will format the output.
	//
	// Note: v can be a struct or a ptr struct (*struct).
	MarshalWrite(destPath string, v interface{}, pretty bool) error

	// Unmarshal is used to store the encoded data to v.
	//
	// Note: v must be a ptr struct (*struct).
	Unmarshal(data []byte, v interface{}) error

	// UnmarshalRead is used to read the encoded data from srcPath and store the encoded data to v.
	//
	// Note: v must be a ptr struct (*struct).
	UnmarshalRead(srcPath string, v interface{}) error
}

// json is a struct to implement the JSON interface.
type json struct{}

// instance hold the JSON interface.
var instance JSON

// Instance will return the JSON interface.
//
// Note: first time calling this function will create a new instance
// to implement the JSON interface (except if you call InitTestMode).
// The next time this function is called, it will use the previously created instance.
func Instance() JSON {
	if instance == nil {
		instance = &json{}
	}

	return instance
}

// InitTestMode will set the JSON instance to mockInstance.
//
// Note: this function for testing purposes only. After executing InitTestMode,
// whenever your unit tests execute Instance it will call mockInstance.
func InitTestMode(mockInstance *mock.MockJSON) {
	instance = mockInstance
}
