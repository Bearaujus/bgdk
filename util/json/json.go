package json

import (
	"bytes"
	encodingJSON "encoding/json"
	"os"
)

func (u *json) Marshal(v interface{}) ([]byte, error) {
	// encode the json data from v
	return encodingJSON.Marshal(v)
}

func (u *json) MarshalWrite(destPath string, v interface{}, pretty bool) error {
	// encode the json data from v
	data, err := encodingJSON.Marshal(v)
	if err != nil {
		return err
	}

	// if pretty format is true
	if pretty {
		// format the encoded json data
		buff := bytes.NewBuffer(nil)
		if err = encodingJSON.Indent(buff, data, "", "\t"); err != nil {
			return err
		}
		data = buff.Bytes()
	}

	// write the encoded json data to the destination path
	return os.WriteFile(destPath, data, os.ModePerm)
}

func (u *json) Unmarshal(data []byte, v interface{}) error {
	// store the encoded json data to v
	return encodingJSON.Unmarshal(data, v)
}

func (u *json) UnmarshalRead(srcPath string, v interface{}) error {
	// read the encoded json data from the source path
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	// store the encoded json data to v
	return encodingJSON.Unmarshal(data, v)
}
