# UTIL/JSON

`utilJSON` is utilities for the JSON files. `utilJSON` implementing core functions from [encoding/json](https://cs.opensource.google/go/go/+/master:/src/encoding/json/).
This package also has some additional I/O features such as JSON writer and JSON reader.

See [BGDK Installation](https://github.com/bearaujus/bgdk#installation)

# IMPORT

- ### Package

```go
import utilJSON "github.com/bearaujus/bgdk/util/json"
```

- ### Mock

```go
import mockUtilJSON "github.com/bearaujus/bgdk/util/json/mock"
```

# DOCUMENTATION

- [Root Functions](#root-functions)
  - [Instance](#instance)
  - [InitTestMode](#inittestmode)
- [Interface Functions](#interface-functions)
  - [JSONMarshal](#jsonmarshal)
  - [JSONMarshalWrite](#jsonmarshalwrite)
  - [JSONMarshalIndent](#jsonmarshalindent)
  - [JSONMarshalIndentWrite](#jsonmarshalindentwrite)
  - [JSONUnmarshal](#jsonunmarshal)
  - [JSONUnmarshalRead](#jsonunmarshalread)

## Root Functions

- ### Instance

```go
// Instance will return the UtilJSON interface.
//
// Note: first time calling this function will create a new instance
// to implement the UtilJSON interface (except if you call InitTestMode).
// The next time this function is called, it will use the previously created instance.
func Instance() UtilJSON
```

&nbsp;

- ### InitTestMode

```go
// InitTestMode will set the UtilJSON instance to mockInstance.
//
// Note: use this function for only testing purposes. After executing InitTestMode,
// whenever your unit tests execute Instance it will call mockInstance.
func InitTestMode(mockInstance *mockUtilJSON.MockUtilJSON)
```

&nbsp;

## Interface Functions

- ### JSONMarshal

```go
// JSONMarshal is used to encode JSON data from v.
//
// Note: v can be a struct or ptr struct (*struct).
func (u *utilJSON) JSONMarshal(v interface{}) ([]byte, error)
```

&nbsp;

- ### JSONMarshalWrite

```go
// JSONMarshalWrite is used to encode JSON data from v and write the encoded JSON data to destPath.
// if pretty is true, it will format the encoded JSON data.
//
// Note: v can be a struct or ptr struct (*struct).
func (u *utilJSON) JSONMarshalWrite(destPath string, v interface{}, pretty bool) error
```

&nbsp;

- ### JSONMarshalIndent

```go
// JSONMarshalIndent is used to encode JSON data from v and format the encoded JSON data.
//
// Note: v can be a struct or ptr struct (*struct).
func (u *utilJSON) JSONMarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
```

&nbsp;

- ### JSONMarshalIndentWrite

```go
// JSONMarshalIndentWrite is used to encode JSON data from v, format the encoded JSON data
// and finally write the encoded JSON data to destPath.
//
// Note: v can be a struct or ptr struct (*struct).
func (u *utilJSON) JSONMarshalIndentWrite(destPath string, v interface{}, prefix, indent string) error
```

&nbsp;

- ### JSONUnmarshal

```go
// JSONUnmarshal is used to store the encoded JSON data to v.
//
// Note: v must be ptr struct (*struct).
func (u *utilJSON) JSONUnmarshal(data []byte, v interface{}) error
```

&nbsp;

- ### JSONUnmarshalRead

```go
// JSONUnmarshalRead is used to read the encoded JSON data from srcPath and store the encoded JSON data to v.
//
// Note: v must be ptr struct (*struct).
func (u *utilJSON) JSONUnmarshalRead(srcPath string, v interface{}) error
```

&nbsp;

[Back to top](#utiljson) 
