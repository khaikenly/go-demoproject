package drivers

import (
	"database/sql"
	"go-demoproject/pkg/config"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MysqlDB struct {
	SQL *sql.DB
}

func (m *MysqlDB) Connect() *MysqlDB {

	if config.DB_TYPE != "mysql" {
		log.Fatal("Type DB not found!! ")
	}

	// Specify connection properties.
	cfgConnectDb := mysql.Config{
		User:      config.DB_USER,
		Passwd:    config.DB_PASSWORD,
		Net:       "tcp",
		Addr:      config.DB_HOST + ":" + config.DB_PORT,
		DBName:    config.DB_NAME,
		ParseTime: true,
	}

	// Get a database handle.
	db, err := sql.Open(config.DB_TYPE, cfgConnectDb.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	m.SQL = db
	return m
}
