package utilYAML

import mockUtilYAML "github.com/bearaujus/bgdk/util/yaml/mock"

// UtilYAML is an interface primarily used by Instance.
type UtilYAML interface {
	// YAMLMarshal is used to encode YAML data from v.
	//
	// Note: v can be a struct or ptr struct (*struct).
	YAMLMarshal(v interface{}) ([]byte, error)

	// YAMLMarshalWrite is used to encode YAML data from v and write the encoded YAML data to destPath.
	//
	// Note: v can be a struct or ptr struct (*struct).
	YAMLMarshalWrite(destPath string, v interface{}) error

	// YAMLUnmarshal is used to store the encoded YAML data to v.
	//
	// Note: v must be ptr struct (*struct).
	YAMLUnmarshal(data []byte, v interface{}) error

	// YAMLUnmarshalRead is used to read the encoded YAML data from srcPath and store the encoded YAML data to v.
	//
	// Note: v must be ptr struct (*struct).
	YAMLUnmarshalRead(srcPath string, v interface{}) error
}

// utilYAML is a struct to implement the UtilYAML interface.
type utilYAML struct{}

// instance hold the UtilYAML interface.
var instance UtilYAML

// Instance will return the UtilYAML interface.
//
// Note: first time calling this function will create a new instance
// to implement the UtilYAML interface (except if you call InitTestMode).
// The next time this function is called, it will use the previously created instance.
func Instance() UtilYAML {
	if instance == nil {
		instance = &utilYAML{}
	}

	return instance
}

// InitTestMode will set the UtilYAML instance to mockInstance.
//
// Note: this function for testing purposes only. After executing InitTestMode,
// whenever your unit tests execute Instance it will call mockInstance.
func InitTestMode(mockInstance *mockUtilYAML.MockUtilYAML) {
	instance = mockInstance
}
