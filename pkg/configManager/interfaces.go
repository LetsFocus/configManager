package configManager

// ConfigManager is the interface for loading configuration files
type ConfigManager interface {
	Load(filePath string) (map[string]string, error)
}

// CacheManager is the interface for managing in-memory cache
type CacheManager interface {
	Get(key string) (string, bool)
	Set(key, value string)
	Flush()
}
