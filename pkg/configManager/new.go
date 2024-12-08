package configManager

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/LetsFocus/configManager/pkg/cache"
)

// Config manages loading and caching configurations
type Config struct {
	cache CacheManager
}

// New initializes a new ConfigManager instance
func New() *Config {
	memoryCache := cache.NewInMemoryCache()
	configManager := &Config{cache: memoryCache}
	basePath := "./configs"
	if err := configManager.LoadConfigs(basePath); err != nil {
		fmt.Printf("Error loading configuration files: %v\n", err)
	}

	return configManager
}

// LoadConfigs loads configuration files with the following rules:
// 1. Checks for `.env`, `.json`, `.yaml` in the given order; stops if one is found and loaded.
// 2. If APP_ENV is set, checks for `APP_ENV.env`, `APP_ENV.json`, `APP_ENV.yaml` in the given order; stops if one is found and loaded.
func (cm *Config) LoadConfigs(basePath string) error {
	if basePath == "" {
		return errors.New("basePath cannot be empty")
	}

	// Load base files in priority order
	cm.loadFirstAvailableFile(basePath, []string{".env", ".json", ".yaml"})

	// Load environment-specific files in priority order, if APP_ENV is set
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "local"
	}

	if appEnv != "" {
		cm.loadFirstAvailableFile(basePath, []string{
			fmt.Sprintf("%s.env", appEnv),
			fmt.Sprintf("%s.json", appEnv),
			fmt.Sprintf("%s.yaml", appEnv),
		})
	}

	return nil
}

// loadFirstAvailableFile checks and loads the first available file from the list
func (cm *Config) loadFirstAvailableFile(basePath string, files []string) bool {
	for _, file := range files {
		fullPath := filepath.Join(basePath, file)
		if _, err := os.Stat(fullPath); err == nil {
			err := cm.loadFile(fullPath)
			if err != nil {
				fmt.Printf("Error loading file %s: %v\n", fullPath, err)
			} else {
				fmt.Printf("Loaded configuration from %s\n", fullPath)
			}
			return true // Stop after the first successfully loaded file
		}
	}
	return false // No file found
}

// loadFile uses the appropriate loader to load a configuration file
func (cm *Config) loadFile(file string) error {
	loader, err := LoaderFactory(file)
	if err != nil {
		return fmt.Errorf("unsupported file type for %s: %v", file, err)
	}

	configs, err := loader.Load(file)
	if err != nil {
		return fmt.Errorf("error loading file %s: %v", file, err)
	}

	for key, value := range configs {
		os.Setenv(key, value) // Update environment variables
		cm.cache.Set(key, value)
	}

	return nil
}

// GetConfig retrieves a configuration value from the cache or environment variables
func (cm *Config) GetConfig(key string) string {
	if value, found := cm.cache.Get(key); found {
		return value
	}
	return os.Getenv(key)
}

// Unmarshal binds configuration values to a given struct using tags or field names
func (cm *Config) Unmarshal(target interface{}) error {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.New("target must be a pointer to a struct")
	}

	v = v.Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Recursive call for nested structs
		if field.Kind() == reflect.Struct {
			if err := cm.Unmarshal(field.Addr().Interface()); err != nil {
				return err
			}
			continue
		}

		// Retrieve environment variable key
		envKey := fieldType.Tag.Get("env")
		if envKey == "" {
			envKey = strings.ToUpper(fieldType.Name)
		}

		// Retrieve environment variable value
		envValue, found := os.LookupEnv(envKey)
		if !found {
			defaultValue := fieldType.Tag.Get("default")
			if defaultValue != "" {
				envValue = defaultValue
			} else if fieldType.Tag.Get("required") == "true" {
				return fmt.Errorf("missing required environment variable: %s", envKey)
			} else {
				continue
			}
		}

		// Set field value
		if err := setFieldValue(field, envValue); err != nil {
			return fmt.Errorf("error setting field %s: %v", fieldType.Name, err)
		}
	}

	return nil
}

// setFieldValue sets a value to a struct field based on its type
func setFieldValue(field reflect.Value, value string) error {
	if !field.CanSet() {
		return errors.New("field cannot be set")
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(intValue)
	case reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(floatValue)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(boolValue)
	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}

	return nil
}
