# UTIL/YAML

`utilYAML` is utilities for the YAML files. `utilYAML` implementing core functions from [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3).
This package also has some additional I/O features such as YAML writer and YAML reader.

See [BGDK Installation](https://github.com/bearaujus/bgdk#installation)

# IMPORT

- ### Package

```go
import utilYAML "github.com/bearaujus/bgdk/util/yaml"
```

- ### Mock

```go
import mockUtilYAML "github.com/bearaujus/bgdk/util/yaml/mock"
```

# DOCUMENTATION

- [Root Functions](#root-functions)
  - [Instance](#instance)
  - [InitTestMode](#inittestmode)
- [Interface Functions](#interface-functions)
  - [YAMLMarshal](#yamlmarshal)
  - [YAMLMarshalWrite](#yamlmarshalwrite)
  - [YAMLUnmarshal](#yamlunmarshal)
  - [YAMLUnmarshalRead](#yamlunmarshalread)

## Root Functions

- ### Instance

```go
// Instance will return the UtilYAML interface.
//
// Note: first time calling this function will create a new instance
// to implement the UtilYAML interface (except if you call InitTestMode).
// The next time this function is called, it will use the previously created instance.
func Instance() UtilYAML
```

&nbsp;

- ### InitTestMode

```go
// InitTestMode will set the UtilYAML instance to mockInstance.
//
// Note: use this function for only testing purposes. After executing InitTestMode,
// whenever your unit tests execute Instance it will call mockInstance.
func InitTestMode(mockInstance *mockUtilYAML.MockUtilYAML)
```

&nbsp;

## Interface Functions

- ### YAMLMarshal

```go
// YAMLMarshal is used to encode YAML data from v.
//
// Note: v can be a struct or ptr struct (*struct).
func (u *utilYAML) YAMLMarshal(v interface{}) ([]byte, error)
```

&nbsp;


- ### YAMLMarshalWrite

```go
// YAMLMarshalWrite is used to encode YAML data from v and write the encoded YAML data to destPath.
//
// Note: v can be a struct or ptr struct (*struct).
func (u *utilYAML) YAMLMarshalWrite(destPath string, v interface{}) error
```

&nbsp;


- ### YAMLUnmarshal

```go
// YAMLUnmarshal is used to store the encoded YAML data to v.
//
// Note: v must be ptr struct (*struct).
func (u *utilYAML) YAMLUnmarshal(data []byte, v interface{}) error
```

&nbsp;


- ### YAMLUnmarshalRead

```go
// YAMLUnmarshalRead is used to read the encoded YAML data from srcPath and store the encoded YAML data to v.
//
// Note: v must be ptr struct (*struct).
func (u *utilYAML) YAMLUnmarshalRead(srcPath string, v interface{}) error
```

&nbsp;

[Back to top](#utilyaml) 
