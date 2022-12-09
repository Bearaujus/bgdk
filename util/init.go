package util

import mockUtil "github.com/bearaujus/bgdk/util/util.mock"

// Util is interface primarily used by util.
type Util interface {
	// WriteMarshalJSON is used to encode JSON from v and write the data to destPath.
	// if prettyFormat value is true, it will format the JSON.
	//
	// Note: v can be a struct or ptr struct (*struct).
	WriteMarshalJSON(destPath string, v interface{}, prettyFormat bool) error

	// WriteMarshalYAML is used to encode YAML from v and write the data to destPath.
	//
	// Note: v can be a struct or ptr struct (*struct).
	WriteMarshalYAML(destPath string, v interface{}) error

	// ReadUnmarshalJSON is used to read the data to srcPath and store the data to v.
	//
	// Note: v must be ptr struct (*struct).
	ReadUnmarshalJSON(srcPath string, v interface{}) error

	// ReadUnmarshalYAML is used to read the data to srcPath and store the data to v.
	//
	// Note: v must be ptr struct (*struct).
	ReadUnmarshalYAML(srcPath string, v interface{}) error
}

// util is a struct to implement the Util interface.
type util struct{}

// instance hold the Util interface.
var instance Util

// Instance will return Util interface.
//
// Note: if this is your first time calling this function and never executed InitTestMode before.
// it will create a new instance to implement the Util interface. The next time you call Instance,
// it will use previously created instance.
func Instance() Util {
	if instance == nil {
		instance = &util{}
	}

	return instance
}

// InitTestMode will set the instance to mockInstance.
//
// Note: you can use this function for testing purposes. After you execute InitTestMode,
// when your unit tests execute Instance, it will call mockInstance.
func InitTestMode(mockInstance *mockUtil.MockUtil) {
	instance = mockInstance
}
