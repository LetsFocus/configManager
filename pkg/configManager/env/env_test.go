package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvLoader_Load(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		expected    map[string]string
		expectError bool
	}{
		{
			name: "Valid .env file with simple key-value pairs",
			fileContent: `
				KEY1=value1
				KEY2=value2
			`,
			expected: map[string]string{
				"KEY1": "value1",
				"KEY2": "value2",
			},
			expectError: false,
		},
		{
			name: "Valid .env file with comments and blank lines",
			fileContent: `
				# This is a comment
				KEY1=value1

				# Another comment
				KEY2=value2
			`,
			expected: map[string]string{
				"KEY1": "value1",
				"KEY2": "value2",
			},
			expectError: false,
		},
		{
			name: "Valid .env file with spaces around keys and values",
			fileContent: `
				KEY1 = value1
				KEY2= value2
				KEY3 =value3
			`,
			expected: map[string]string{
				"KEY1": "value1",
				"KEY2": "value2",
				"KEY3": "value3",
			},
			expectError: false,
		},
		{
			name: "Valid .env file with empty values",
			fileContent: `
				KEY1=
				KEY2=value2
			`,
			expected: map[string]string{
				"KEY1": "",
				"KEY2": "value2",
			},
			expectError: false,
		},
		{
			name: "Invalid .env file with missing `=`",
			fileContent: `
				KEY1=value1
				KEY2
			`,
			expected:    map[string]string{"KEY1": "value1"},
			expectError: false, // Should handle missing "=" gracefully, ignoring the malformed line
		},
		{
			name: "Empty .env file",
			fileContent: `
			`,
			expected:    map[string]string{},
			expectError: false,
		},
		{
			name:        "File does not exist",
			fileContent: "",
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file for testing if fileContent is provided
			var filePath string
			var err error
			if tt.fileContent != "" {
				file, err := os.CreateTemp("", "test_*.env")
				assert.NoError(t, err, "Failed to create temp file")
				defer os.Remove(file.Name())
				filePath = file.Name()

				_, err = file.WriteString(tt.fileContent)
				assert.NoError(t, err, "Failed to write to temp file")

				file.Close()
			} else {
				filePath = "nonexistent.env"
			}

			loader := &EnvLoader{}
			result, err := loader.Load(filePath)

			if tt.expectError {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.Equal(t, tt.expected, result, "Loaded configuration did not match expected")
			}
		})
	}
}
