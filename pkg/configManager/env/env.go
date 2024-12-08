package env

import (
	"bufio"
	"os"
	"strings"
)

// EnvLoader implements ConfigLoader for .env files
type EnvLoader struct{}

// Load parses .env files and returns key-value pairs
func (e *EnvLoader) Load(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	configs := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			configs[key] = value
		}
	}
	return configs, scanner.Err()
}
