package configManager

import (
	"fmt"
	"strings"

	"github.com/LetsFocus/configManager/pkg/configManager/env"
	"github.com/LetsFocus/configManager/pkg/configManager/json"
	"github.com/LetsFocus/configManager/pkg/configManager/yaml"
)

// LoaderFactory creates ConfigLoader instances based on file extensions
func LoaderFactory(filePath string) (ConfigManager, error) {
	switch {
	case strings.HasSuffix(filePath, ".env"):
		return &env.EnvLoader{}, nil
	case strings.HasSuffix(filePath, ".yaml"):
		return &yaml.YAMLLoader{}, nil
	case strings.HasSuffix(filePath, ".json"):
		return &json.JSONLoader{}, nil
	default:
		return nil, fmt.Errorf("unsupported file type: %s", filePath)
	}
}
