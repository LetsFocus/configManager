package yaml

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYAMLLoader_Load(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		expected    map[string]string
		expectError bool
	}{
		{
			name: "Valid YAML with flat structure",
			fileContent: `
key1: value1
key2: value2
`,
			expected: map[string]string{
				"KEY1": "value1",
				"KEY2": "value2",
			},
			expectError: false,
		},
		{
			name: "Valid YAML with nested structure",
			fileContent: `
parent:
  child1: value1
  child2: value2
`,
			expected: map[string]string{
				"PARENT_CHILD1": "value1",
				"PARENT_CHILD2": "value2",
			},
			expectError: false,
		},
		{
			name:        "Invalid YAML content",
			fileContent: `: invalid YAML`,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Empty YAML content",
			fileContent: ``,
			expected:    map[string]string{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.CreateTemp("", "test_*.yaml")
			assert.NoError(t, err, "Failed to create temp file")
			defer os.Remove(file.Name())

			_, err = file.WriteString(tt.fileContent)
			assert.NoError(t, err, "Failed to write to temp file")

			file.Close()

			loader := &YAMLLoader{}
			result, err := loader.Load(file.Name())

			if tt.expectError {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Did not expect an error but got one")
				assert.Equal(t, tt.expected, result, "Loaded configuration did not match expected")
			}
		})
	}
}
