# Config Loader for Go

This Go package provides a simple configuration loader that reads YAML or JSON configuration files using the `cleanenv` library. It includes support for database and host configurations, as well as robust error handling and logging.

## Features

- Loads configuration from a file (YAML or JSON).
- Supports database (`DBConfig`) and host (`HostConfig`) configurations.
- Uses `cleanenv` for structured, environment-friendly configuration parsing.
- Logs errors using `slog` when configuration loading fails.

## Installation

To install the package, you can use:

```bash
go get github.com/TrueIronEnjoyer/config
```

## Usage

### Configuration Structure

The package defines the following structure for configuration:

```go
type Config struct {
    DB   DBConfig   `json:"db,omitempty" yaml:"db"`
    Host HostConfig `json:"host,omitempty" yaml:"host"`
}
```

- **DBConfig**: Holds database connection details such as `url`, `port`, `username`, `password`, and `full_url`.
- **HostConfig**: Holds the host settings like `url` and `port`.

### Example Configuration File (YAML)

```yaml
db:
  url: "localhost"
  port: "5432"
  username: "user"
  password: "password"
  full_url: "postgres://user:password@localhost:5432/dbname"

host:
  url: "http://localhost"
  port: "8080"
```

### Loading the Configuration

To load a configuration file, call the `Load` method from the `Config` struct and pass the file path:

```go
package main

import (
    "log"
    "github.com/yourusername/yourproject/config"
)

func main() {
    var cfg config.Config
    err := cfg.Load("config.yaml")
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Use the config
    log.Printf("DB URL: %s", cfg.DB.Url)
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Author

Aziz Umerov  
Date: October 12, 2024
