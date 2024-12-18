package internal

import (
	"reflect"
	"testing"
)


func TestFlattenMap(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		prefix   string
		expected map[string]string
	}{
		{
			name: "Empty map",
			input: map[string]interface{}{},
			prefix: "",
			expected: map[string]string{},
		},
		{
			name: "Flat map",
			input: map[string]interface{}{
				"key1": "value1",
				"key2": 42,
			},
			prefix: "",
			expected: map[string]string{
				"KEY1": "value1",
				"KEY2": "42",
			},
		},
		{
			name: "Nested map",
			input: map[string]interface{}{
				"key1": "value1",
				"nested": map[string]interface{}{
					"key2": true,
					"key3": 3.14,
				},
			},
			prefix: "",
			expected: map[string]string{
				"KEY1": "value1",
				"NESTED_KEY2": "true",
				"NESTED_KEY3": "3.14",
			},
		},
		{
			name: "Nested map with prefix",
			input: map[string]interface{}{
				"key1": "value1",
				"nested": map[string]interface{}{
					"key2": false,
				},
			},
			prefix: "PREFIX_",
			expected: map[string]string{
				"PREFIX_KEY1": "value1",
				"PREFIX_NESTED_KEY2": "false",
			},
		},
		{
			name: "Multiple nested maps",
			input: map[string]interface{}{
				"key1": map[string]interface{}{
					"key2": map[string]interface{}{
						"key3": "value3",
					},
				},
			},
			prefix: "",
			expected: map[string]string{
				"KEY1_KEY2_KEY3": "value3",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FlattenMap(test.input, test.prefix)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input %v and prefix %q, expected %v but got %v", test.input, test.prefix, test.expected, result)
			}
		})
	}
}