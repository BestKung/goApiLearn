package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Server struct {
		Port  int  `yaml:"port"`
		Debug bool `yaml:"debug"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`

	Mongo struct {
		Connection string `yaml:"connection"`
	} `yaml:"mongo"`
}

func LoadConfig() *AppConfig {
	cfg := AppConfig{}

	file, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("cfg %T \n", cfg)
	fmt.Println("Server Port:", cfg.Server.Port)
	fmt.Println("Database User:", cfg.Database.User)
	return &cfg
}
