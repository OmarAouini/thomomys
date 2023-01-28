package server

import (
	"database/sql"
	"log"

	"github.com/OmarAouini/thomomys/internal/bootstrap"
	"github.com/OmarAouini/thomomys/internal/config"
	"github.com/OmarAouini/thomomys/internal/database"
)

type Server struct {
	DB *sql.DB
	//TODO router here
}

func NewServer() *Server {
	bootstrap.InitApplication()
	server := Server{}
	server.init()
	return &server
}

func (s *Server) init() {
	if config.Config.Port == 0 {
		log.Fatal("\"port\" definition missing from configuration")
	}
	s.DB = database.ConnectPostgresDB()
}
