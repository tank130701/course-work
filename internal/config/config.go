package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ConnectionType string   `json:"connection_type"`
	HttpPort       string   `json:"http_port"`
	PostgreSQL     DBConfig `json:"postgresql"`
	MySQL          DBConfig `json:"mysql"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"ssl_mode"`
}

func InitConfig() (*Config, error) {
	// Read the JSON file (replace "config.json" with your actual file path).
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("error reading config file %w", err)
	}

	// Parse the JSON data into a Config struct.
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {

		return nil, fmt.Errorf("error parsing config: %w", err)
	}
	return &config, nil
}
