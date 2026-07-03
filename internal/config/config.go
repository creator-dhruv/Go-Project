package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// We use Clean Env package here to load the variable from yaml file or .env file
// Key denotes the type of file used to load the value
// Value denotes the name of environment variable
// These annotations (`yaml:"port" env:"PORT" ...`) are called struct tags

type HttpServer struct {
	Address string `yaml:"address" env:"ADDRESS" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	Port        string `yaml:"port" env:"PORT"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

//

func MustLoad() *Config {
	var configPath string

	// Get Config File (.env) Path from os
	configPath = os.Getenv("CONFIG_PATH")

	// Get Config File Path from flags
	if configPath == "" {

		// Return Flag pointer
		flags := flag.String("config", "", "path to the config file")
		flag.Parse()

		configPath = *flags
		if configPath == "" {
			log.Fatal("No Config path not set yet")
		}
	}

	// Check the file will exist or not
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("Config File is not existed")
	}

	// Read the Config File and Set the variable into struct
	var clg Config

	err := cleanenv.ReadConfig(configPath, &clg)
	if err != nil {
		log.Fatalf("cannot read Config File: %s", err.Error())
	}

	return &clg

}
