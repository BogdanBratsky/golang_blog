package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config структура для представления структуры JSON-конфига
type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

// ServerConfig структура для представления настроек сервера
type ServerConfig struct {
	Port     int    `json:"port"`
	Hostname string `json:"hostname"`
}

// DatabaseConfig структура для представления настроек базы данных
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func GetConfig() Config {
	// Чтение файла конфигурации
	configData, err := os.ReadFile("./config/config.json")
	if err != nil {
		log.Fatal("Unable to read the config file:", err)
	}

	// Распаковка данных из JSON в структуру Config
	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}
	return config
}
