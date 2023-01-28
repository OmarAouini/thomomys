package bootstrap

import (
	"github.com/OmarAouini/thomomys/internal/config"
	"github.com/OmarAouini/thomomys/internal/kafka"
	"github.com/OmarAouini/thomomys/internal/rabbit"
)

func InitApplication() {
	//env vars
	config.ParseYamlConfigFile()
	// print config (local and development mode)
	config.DebugYamlConfig()
	//kafka
	kafka.InitKafka()
	//rabbit
	rabbit.InitRabbit()
}
