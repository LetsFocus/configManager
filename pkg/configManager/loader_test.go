package configManager

import (
	"errors"
	"testing"

	"github.com/LetsFocus/configManager/pkg/configManager/env"
	"github.com/LetsFocus/configManager/pkg/configManager/json"
	"github.com/LetsFocus/configManager/pkg/configManager/yaml"
	"github.com/stretchr/testify/assert"
)

func TestLoaderFactory(t *testing.T) {
	tests := []struct {
		name          string
		filePath      string
		expectedType  interface{}
		expectedError error
	}{
		{
			name:          "Valid .env file",
			filePath:      "config.env",
			expectedType:  &env.EnvLoader{},
			expectedError: nil,
		},
		{
			name:          "Valid .yaml file",
			filePath:      "config.yaml",
			expectedType:  &yaml.YAMLLoader{},
			expectedError: nil,
		},
		{
			name:          "Valid .json file",
			filePath:      "config.json",
			expectedType:  &json.JSONLoader{},
			expectedError: nil,
		},
		{
			name:          "Unsupported file type",
			filePath:      "config.txt",
			expectedType:  nil,
			expectedError: errors.New("unsupported file type: config.txt"),
		},
		{
			name:          "Empty file name",
			filePath:      "",
			expectedType:  nil,
			expectedError: errors.New("unsupported file type: "),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loader, err := LoaderFactory(tt.filePath)

			// Assert error
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error(), "error mismatch")
			} else {
				assert.NoError(t, err, "unexpected error")
			}

			// Assert loader type
			if tt.expectedType != nil {
				assert.IsType(t, tt.expectedType, loader, "loader type mismatch")
			} else {
				assert.Nil(t, loader, "expected loader to be nil")
			}
		})
	}
}
