package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/OmarAouini/thomomys/internal/config"
	_ "github.com/lib/pq"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// postgres connection pool, to be used after the yaml config file is loaded
func ConnectPostgresDB() *sql.DB {
	if config.Config.Database == nil || config.Config.Database.Postgres == nil {
		log.Fatal("postgres database connection details missing from configuration: database.postgres")
	}

	pgRequiredConnectionDetails := map[string]string{
		"postgres.name":           config.Config.Database.Postgres.Name,
		"postgres.user":           config.Config.Database.Postgres.User,
		"postgres.password":       config.Config.Database.Postgres.Password,
		"postgres.ssl_mode":       config.Config.Database.Postgres.SslMode,
		"postgres.schema":         config.Config.Database.Postgres.Schema,
		"postgres.min_connection": fmt.Sprint(config.Config.Database.Postgres.MinimumConnection),
		"postgres.max_connection": fmt.Sprint(config.Config.Database.Postgres.MaximumConnection),
		"postgres.timezone":       config.Config.Database.Postgres.TimeZone,
	}

	missingPgRequiredConnectionDetails := []string{}
	for k, v := range pgRequiredConnectionDetails {
		if v == "" {
			missingPgRequiredConnectionDetails = append(missingPgRequiredConnectionDetails, k)
		}
	}
	if len(missingPgRequiredConnectionDetails) > 0 {
		var str string
		for _, i := range missingPgRequiredConnectionDetails {
			str += fmt.Sprintf("- %s\n", i)
		}
		log.Fatalf("postgres database connection details missing from configuration: \n%v\n", str)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s schema=%s",
		config.Config.Database.Postgres.Name, config.Config.Database.Postgres.User, config.Config.Database.Postgres.Password, config.Config.Database.Postgres.SslMode, config.Config.Database.Postgres.Schema))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connected to postgres database \"%s\"\n", config.Config.Database.Postgres.Name)
	return db
}
