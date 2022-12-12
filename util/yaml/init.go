package yaml

import "github.com/bearaujus/bgdk/util/yaml/mock"

// YAML is an interface primarily used by Instance.
type YAML interface {
	// Marshal is used to encode the data from v.
	//
	// Note: v can be a struct or a ptr struct (*struct).
	Marshal(v interface{}) ([]byte, error)

	// MarshalWrite is used to encode data from v and write the encoded data to destPath.
	//
	// Note: v can be a struct or a ptr struct (*struct).
	MarshalWrite(destPath string, v interface{}) error

	// Unmarshal is used to store the encoded data to v.
	//
	// Note: v must be a ptr struct (*struct).
	Unmarshal(data []byte, v interface{}) error

	// UnmarshalRead is used to read the encoded data from srcPath and store the encoded data to v.
	//
	// Note: v must be a ptr struct (*struct).
	UnmarshalRead(srcPath string, v interface{}) error
}

// yaml is a struct to implement the YAML interface.
type yaml struct{}

// instance hold the YAML interface.
var instance YAML

// Instance will return the YAML interface.
//
// Note: first time calling this function will create a new instance
// to implement the YAML interface (except if you call InitTestMode).
// The next time this function is called, it will use the previously created instance.
func Instance() YAML {
	if instance == nil {
		instance = &yaml{}
	}

	return instance
}

// InitTestMode will set the YAML instance to mockInstance.
//
// Note: this function for testing purposes only. After executing InitTestMode,
// whenever your unit tests execute Instance it will call mockInstance.
func InitTestMode(mockInstance *mock.MockYAML) {
	instance = mockInstance
}
