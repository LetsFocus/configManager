
# configManager: Advanced Configuration Loader for Go

## Overview
`configManager` is a Go module that simplifies the management of configuration data from various sources such as environment variables, JSON, YAML, and `.env` files. It provides a flexible way to handle complex configurations, including nested structs, validation, default values, and custom parsing. This module also caches configuration values in memory for efficient access.

## Features
- **Multiple Sources**: Supports `.env`, JSON, and YAML files.
- **Default Values**: Automatically applies default values when environment variables are missing.
- **Validation**: Supports required fields and throws errors for missing variables.
- **Nested Structs**: Handles deeply nested structs with ease.
- **Custom Parsing**: Allows custom parsing (e.g., JSON strings).
- **Environment-Specific Files**: Supports app-specific configurations based on the `APP_ENV` variable (e.g., `.dev.env`, `.prod.env`).
- **Caching**: Caches frequently accessed configuration data to improve performance.
- **Recursive Directory Search**: Searches up to 3 levels of subdirectories to find configuration files.
- **Automatic Binding**: Automatically binds configuration data to struct fields.

## Installation

To install the module, use the following command:

```bash
go get github.com/LetsFocus/configManager
```

## Example Usage

### 1. Load Configurations

```go
package main

import (
    "fmt"
    "log"
    "github.com/LetsFocus/configManager"
)

func main() {
    // Create a new config manager
    cm := configManager.New()
	
    // Define a struct to hold your configuration
    var config AppConfig

    // Unmarshal data from environment variables into the config struct
    err = cm.Unmarshal(&config)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Config Loaded: %+v\n", config)
}
```

### 2. Define Configuration Struct

```go
type AppConfig struct {
    DBUrl      string `env:"DB_URL" default:"postgres://localhost:5432" required:"true"`
    APIToken   string `env:"API_TOKEN"`
    ServerPort int    `env:"SERVER_PORT" default:"8080"`
    DebugMode  bool   `env:"DEBUG_MODE" default:"false"`
}
```

### 3. Unmarshal Function

The `Unmarshal` function automatically populates your struct fields based on environment variables or configuration file data.

```go
err := cm.Unmarshal(&config)
if err != nil {
    log.Fatalf("Error loading config: %v", err)
}
```

### 4. Caching Configuration Data

To optimize access to frequently used configuration data, the module provides in-memory caching. Data is stored in a small in-memory cache to avoid repeatedly accessing the OS or reading from files.

### 5. Clearing Cache

You can clear the in-memory cache if needed:

```go
cm.ClearCache()
```

## Configuration File Support

This module supports three main configuration file types:

1. **`.env` Files**
    - Each line in the `.env` file contains a key-value pair.
    - Example:
      ```
      DB_URL=postgres://localhost:5432
      API_TOKEN=my-api-token
      SERVER_PORT=8080
      DEBUG_MODE=true
      ```

2. **JSON Files**
    - Configuration data can be in standard JSON format.
    - Example:
      ```json
      {
        "DB_URL": "postgres://localhost:5432",
        "API_TOKEN": "my-api-token",
        "SERVER_PORT": 8080,
        "DEBUG_MODE": true
      }
      ```

3. **YAML Files**
    - Configuration data in YAML format.
    - Example:
      ```yaml
      DB_URL: postgres://localhost:5432
      API_TOKEN: my-api-token
      SERVER_PORT: 8080
      DEBUG_MODE: true
      ```
## How It Works

The `configManager` module performs the following actions:

1. It searches for configuration files (up to 3 levels deep) in a specified directory (`basePath`).
2. It loads configuration files based on a priority order: first `.env`, then `.json`, and lastly `.yaml`. If a configuration file is found, it will not search for other file types.
3. It supports environment-specific configurations using the `APP_ENV` environment variable. For example, if `APP_ENV` is set to `dev`, the module will look for `.dev.env`, `.dev.json`, or `.dev.yaml` files in the specified directory.
4. It parses the configuration files and environment variables.
5. It binds the data to a provided struct using reflection.
6. It supports caching to avoid re-reading the same configuration multiple times, ensuring better performance.

### Configuration Loading Process:

1. **Search for Config Files**: The module scans the specified directory for valid `.env`, `.json`, or `.yaml` files.
   - First, it looks for `.env`, `.json`, or `.yaml` files in the base directory in the order of priority.
   - Then, if the `APP_ENV` environment variable is set, it looks for environment-specific files, such as `.dev.env`, `.prod.env`, `.dev.json`, or `.prod.json`, in the same priority order.
2. **Load Data**: It reads the file content, parses the data, and loads it into memory.
3. **Environment Variables**: Configuration values are loaded into environment variables, and the struct fields are populated from these values using reflection.
4. **Cache**: Frequently accessed configuration data is cached in memory to avoid reloading it repeatedly, improving performance.
5. **Validation and Defaults**: It ensures that required fields are set and assigns default values to fields that are missing.

## Advanced Features

- **File Priority**: The module loads configuration files based on a defined priority: `.env` > `.json` > `.yaml`. If a file is found in one of these formats, it will stop searching for the other formats.
- **Environment-Specific Files**: Supports loading different configuration files based on the environment (e.g., `.dev.env`, `.prod.env`). If the `APP_ENV` environment variable is set, it will attempt to load corresponding environment-specific files.
- **Nested Structs**: Supports nested structs, allowing for more complex configuration structures (e.g., YAML, JSON files with nested fields).
- **Custom Parsing**: Supports custom types by implementing the `Unmarshal` interface. This allows for more advanced data manipulation during the unmarshalling process.
- **Validation**: Ensures that required configuration fields are set and validates their values, ensuring that no required configurations are missing.
- **Cache Management**: The cache can be cleared manually or set to expire after a certain period, ensuring that configuration data remains up-to-date.
- **Flexible Configuration Sources**: The module supports loading configuration data from `.env`, `.json`, and `.yaml` files. The order of loading is flexible based on whether `APP_ENV` is set or not.

## Contributions

Feel free to fork the repository, make changes, and create pull requests! We welcome contributions that improve functionality or fix bugs.
