package utilJSON

import mockUtilJSON "github.com/bearaujus/bgdk/util/json/mock"

// UtilJSON is an interface primarily used by Instance.
type UtilJSON interface {
	// JSONMarshal is used to encode JSON data from v.
	//
	// Note: v can be a struct or ptr struct (*struct).
	JSONMarshal(v interface{}) ([]byte, error)

	// JSONMarshalWrite is used to encode JSON data from v and write the encoded JSON data to destPath.
	// if pretty is true, it will format the encoded JSON data.
	//
	// Note: v can be a struct or ptr struct (*struct).
	JSONMarshalWrite(destPath string, v interface{}, pretty bool) error

	// JSONMarshalIndent is used to encode JSON data from v and format the encoded JSON data.
	//
	// Note: v can be a struct or ptr struct (*struct).
	JSONMarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

	// JSONMarshalIndentWrite is used to encode JSON data from v, format the encoded JSON data
	// and finally write the encoded JSON data to destPath.
	//
	// Note: v can be a struct or ptr struct (*struct).
	JSONMarshalIndentWrite(destPath string, v interface{}, prefix, indent string) error

	// JSONUnmarshal is used to store the encoded JSON data to v.
	//
	// Note: v must be ptr struct (*struct).
	JSONUnmarshal(data []byte, v interface{}) error

	// JSONUnmarshalRead is used to read the encoded JSON data from srcPath and store the encoded JSON data to v.
	//
	// Note: v must be ptr struct (*struct).
	JSONUnmarshalRead(srcPath string, v interface{}) error
}

// utilJSON is a struct to implement the UtilJSON interface.
type utilJSON struct{}

// instance hold the UtilJSON interface.
var instance UtilJSON

// Instance will return the UtilJSON interface.
//
// Note: first time calling this function will create a new instance
// to implement the UtilJSON interface (except if you call InitTestMode).
// The next time this function is called, it will use the previously created instance.
func Instance() UtilJSON {
	if instance == nil {
		instance = &utilJSON{}
	}

	return instance
}

// InitTestMode will set the UtilJSON instance to mockInstance.
//
// Note: this function for testing purposes only. After executing InitTestMode,
// whenever your unit tests execute Instance it will call mockInstance.
func InitTestMode(mockInstance *mockUtilJSON.MockUtilJSON) {
	instance = mockInstance
}
