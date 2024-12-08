package json

import (
	"encoding/json"
	"io/ioutil"

	"github.com/LetsFocus/configManager/internal"
)

// JSONLoader implements ConfigLoader for .json files
type JSONLoader struct{}

// Load parses JSON files and returns key-value pairs
func (j *JSONLoader) Load(filePath string) (map[string]string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return internal.FlattenMap(data, ""), nil
}
