package utilYAML

import (
	"gopkg.in/yaml.v3"
	"os"
)

func (u *utilYAML) YAMLMarshal(v interface{}) ([]byte, error) {
	// encode the yaml data from v
	return yaml.Marshal(v)
}

func (u *utilYAML) YAMLMarshalWrite(destPath string, v interface{}) error {
	// encode the yaml data from v
	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	// write the encoded yaml data to the destination path
	return os.WriteFile(destPath, data, os.ModePerm)
}

func (u *utilYAML) YAMLUnmarshal(data []byte, v interface{}) error {
	// store the encoded yaml data to v
	return yaml.Unmarshal(data, v)
}

func (u *utilYAML) YAMLUnmarshalRead(srcPath string, v interface{}) error {
	// read the encoded yaml data from the source path
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	// store the encoded yaml data to v
	return yaml.Unmarshal(data, v)
}
