package configManager

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {

	type configStruct struct {
		Name     string `env:"NAME"`
		Port     int    `env:"PORT"`
		Required string `env:"REQUIRED" required:"true"`
		Optional string `env:"OPTIONAL" default:"default_value"`
	}

	config := New()

	// Set environment variables
	os.Setenv("NAME", "test")
	os.Setenv("PORT", "8080")
	os.Setenv("REQUIRED", "required_value")
	defer func() {
		os.Unsetenv("NAME")
		os.Unsetenv("PORT")
		os.Unsetenv("REQUIRED")
	}()

	var cfg configStruct
	err := config.Unmarshal(&cfg)
	assert.NoError(t, err, "Unmarshal should not return an error")
	assert.Equal(t, "test", cfg.Name, "Unmarshal should set the Name field")
	assert.Equal(t, 8080, cfg.Port, "Unmarshal should set the Port field")
	assert.Equal(t, "required_value", cfg.Required, "Unmarshal should set the Required field")
	assert.Equal(t, "default_value", cfg.Optional, "Unmarshal should set the Optional field to its default")
}

func TestGetConfigWithDefault(t *testing.T) {
	config := New()

	// Set a value in the cache
	config.cache.Set("KEY", "value")

	// Set an environment variable
	os.Setenv("ENV_KEY", "env_value")
	defer os.Unsetenv("ENV_KEY")

	// Test cache value retrieval
	assert.Equal(t, "value", config.GetConfigWithDefault("KEY", "default"), "GetConfigWithDefault should return the cached value")

	// Test environment variable retrieval
	assert.Equal(t, "env_value", config.GetConfigWithDefault("ENV_KEY", "default"), "GetConfigWithDefault should return the environment variable value")

	// Test missing key with default
	assert.Equal(t, "default", config.GetConfigWithDefault("MISSING_KEY", "default"), "GetConfigWithDefault should return the default value for missing keys")
}

func TestGetConfig(t *testing.T) {
	config := New()

	// Set a value in the cache
	config.cache.Set("KEY", "value")

	// Set an environment variable
	os.Setenv("ENV_KEY", "env_value")
	defer os.Unsetenv("ENV_KEY")

	// Test cache value retrieval
	assert.Equal(t, "value", config.GetConfig("KEY"), "GetConfig should return the cached value")

	// Test environment variable retrieval
	assert.Equal(t, "env_value", config.GetConfig("ENV_KEY"), "GetConfig should return the environment variable value")

	// Test missing key
	assert.Equal(t, "", config.GetConfig("MISSING_KEY"), "GetConfig should return an empty string for missing keys")
}

func TestLoadFile(t *testing.T) {
	config := New()

	// Create a mock file
	filePath := "./mock_config.json"
	os.WriteFile(filePath, []byte("{}"), 0644)
	defer os.Remove(filePath)

	err := config.loadFile(filePath)
	assert.NoError(t, err, "loadFile should not return an error for valid files")
}

func TestLoadConfigs(t *testing.T) {
	config := New()

	// Set up a mock configuration directory
	basePath := "./mock_configs"
	os.Mkdir(basePath, 0755)
	defer os.RemoveAll(basePath)

	// Create dummy files
	os.WriteFile(basePath+"/local.yaml", []byte("key: value"), 0644)
	os.Setenv("APP_ENV", "local")
	defer os.Unsetenv("APP_ENV")

	err := config.LoadConfigs(basePath)
	assert.NoError(t, err, "LoadConfigs should not return an error")
}