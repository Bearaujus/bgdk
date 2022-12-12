package yaml

import (
	goPkgInYaml "gopkg.in/yaml.v3"
	"os"
)

func (u *yaml) Marshal(v interface{}) ([]byte, error) {
	// encode the yaml data from v
	return goPkgInYaml.Marshal(v)
}

func (u *yaml) MarshalWrite(destPath string, v interface{}) error {
	// encode the yaml data from v
	data, err := goPkgInYaml.Marshal(v)
	if err != nil {
		return err
	}

	// write the encoded yaml data to the destination path
	return os.WriteFile(destPath, data, os.ModePerm)
}

func (u *yaml) Unmarshal(data []byte, v interface{}) error {
	// store the encoded yaml data to v
	return goPkgInYaml.Unmarshal(data, v)
}

func (u *yaml) UnmarshalRead(srcPath string, v interface{}) error {
	// read the encoded yaml data from the source path
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	// store the encoded yaml data to v
	return goPkgInYaml.Unmarshal(data, v)
}
