package util

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
)

func (u *util) WriteMarshalJSON(destPath string, v interface{}, prettyFormat bool) error {
	// if prettyFormat is true
	if prettyFormat {
		// encode the json from v formatted
		data, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			return err
		}

		// write the formatted encoded json to the destination path
		return os.WriteFile(destPath, data, os.ModePerm)
	}

	// encode the json from v
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	// write the encoded json to the destination path
	return os.WriteFile(destPath, data, os.ModePerm)
}

func (u *util) WriteMarshalYAML(destPath string, v interface{}) error {
	// encode the yaml from v
	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	// write the encoded yaml to the destination path
	return os.WriteFile(destPath, data, os.ModePerm)
}

func (u *util) ReadUnmarshalJSON(srcPath string, v interface{}) error {
	// read the json from the source path
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	// store the encoded json to v
	return json.Unmarshal(data, v)
}

func (u *util) ReadUnmarshalYAML(srcPath string, v interface{}) error {
	// read the yaml from the source path
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	// store the encoded yaml to v
	return yaml.Unmarshal(data, v)
}
