package internal

import (
	"fmt"
	"strings"
)

// FlattenMap recursively flattens nested maps into a single-level map
func FlattenMap(data map[string]interface{}, prefix string) map[string]string {
	result := make(map[string]string)
	for key, value := range data {
		envKey := prefix + strings.ToUpper(key)
		switch v := value.(type) {
		case string:
			result[envKey] = v
		case int, int64, float64, bool:
			result[envKey] = fmt.Sprintf("%v", v)
		case map[string]interface{}:
			subMap := FlattenMap(v, envKey+"_")
			for subKey, subValue := range subMap {
				result[subKey] = subValue
			}
		}
	}

	return result
}
