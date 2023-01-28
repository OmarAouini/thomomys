package config

import (
	"fmt"
	"log"
	"os"

	"github.com/OmarAouini/thomomys/pkg/utilities"
	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Environment string `yaml:"environment"`
	Key         string `yaml:"key"`
	AppName     string `yaml:"application_name"`
	Port        int    `yaml:"port"`
	Database    *struct {
		Postgres *struct {
			Name              string `yaml:"name"`
			User              string `yaml:"user"`
			Password          string `yaml:"password"`
			Schema            string `yaml:"schema"`
			MinimumConnection int    `yaml:"min_connection"`
			MaximumConnection int    `yaml:"max_connection"`
			SslMode           string `yaml:"ssl_mode"`
			TimeZone          string `yaml:"timezone"`
		}
		MySql *struct {
			Name              string `yaml:"name"`
			User              string `yaml:"user"`
			Password          string `yaml:"password"`
			MinimumConnection int    `yaml:"min_connection"`
			MaximumConnection int    `yaml:"max_connection"`
			SslMode           string `yaml:"ssl_mode"`
			TimeZone          string `yaml:"timezone"`
		}
		//TODO mongo
	}
	Kafka *struct {
		Brokers []string `yaml:"brokers"`
	}
	RabbitMQ *struct {
	}
}

const (
	dev   string = "development"
	prod  string = "production"
	local string = "local"
)

var (
	environments []string = []string{local, dev, prod}
)

// global app config wrapper
var Config YamlConfig

func ParseYamlConfigFile() {

	env := os.Getenv("profile")
	if env == "" {
		log.Fatal("\"profile\" environment variable not set!")
	}

	if !utilities.ContainsString(environments, env) {
		log.Fatalf("environment \"%s\" not available", env)
	}

	log.Println("parsing yaml config file...")

	confContent, err := os.ReadFile(fmt.Sprintf("config-%s.yml", env))
	if err != nil {
		log.Fatal(err)
	}

	if env == prod {
		// expand environment variables for replacing placeholder in production config file
		// with injected variables from the environment
		confContent = []byte(os.ExpandEnv(string(confContent)))
	}

	conf := &YamlConfig{}
	if err := yaml.Unmarshal(confContent, conf); err != nil {
		log.Fatal(err)
	}
	Config = *conf
	log.Println("yaml config file loaded.")
}

// utility for printing on console during development the yaml config file loaded into the app
func DebugYamlConfig() {
	if Config.Environment != prod {
		conf, _ := utilities.PrettyStruct(Config)
		fmt.Println("\n############################################################################")
		fmt.Println("############################################################################")
		fmt.Println("CONFIG:")
		fmt.Println(conf)
		fmt.Println("############################################################################")
		fmt.Println("############################################################################")
	}
}
