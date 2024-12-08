package yaml

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"

	"github.com/LetsFocus/configManager/internal"
)

// YAMLLoader implements ConfigLoader for .yaml files
type YAMLLoader struct{}

// Load parses YAML files and returns key-value pairs
func (y *YAMLLoader) Load(filePath string) (map[string]string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return internal.FlattenMap(data, ""), nil
}
