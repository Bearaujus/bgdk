package utilJSON

import (
	"encoding/json"
	"os"
)

func (u *utilJSON) JSONMarshal(v interface{}) ([]byte, error) {
	// encode the json data from v
	return json.Marshal(v)
}

func (u *utilJSON) JSONMarshalWrite(destPath string, v interface{}, pretty bool) error {
	// if pretty format is true
	if pretty {
		// encode the json data from v and format the encoded json data
		data, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			return err
		}

		// write the formatted encoded json data to the destination path
		return os.WriteFile(destPath, data, os.ModePerm)
	}

	// encode the json data from v
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	// write the encoded json data to the destination path
	return os.WriteFile(destPath, data, os.ModePerm)
}

func (u *utilJSON) JSONMarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	// encode the json data from v and format the encoded json data
	return json.MarshalIndent(v, prefix, indent)
}

func (u *utilJSON) JSONMarshalIndentWrite(destPath string, v interface{}, prefix, indent string) error {
	// encode the json data from v
	data, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return err
	}

	// write the encoded json data to the destination path
	return os.WriteFile(destPath, data, os.ModePerm)
}

func (u *utilJSON) JSONUnmarshal(data []byte, v interface{}) error {
	// store the encoded json data to v
	return json.Unmarshal(data, v)
}

func (u *utilJSON) JSONUnmarshalRead(srcPath string, v interface{}) error {
	// read the encoded json data from the source path
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	// store the encoded json data to v
	return json.Unmarshal(data, v)
}
