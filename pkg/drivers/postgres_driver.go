package drivers

import (
	"database/sql"
	"fmt"
	"go-demoproject/pkg/config"
	"log"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	SQL *sql.DB
}

func (p *PostgresDB) Connect() *PostgresDB {
	if config.DB_TYPE != "postgres" {
		log.Fatal("Database type not found!! ")
	}

	// Initialize connection string.
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.POSTGRES_SSL)

	// Initialize connection object.
	db, err := sql.Open(config.DB_TYPE, connectionString)
	if err != nil {
		log.Fatal(err)
	}

	p.SQL = db
	return p
}
