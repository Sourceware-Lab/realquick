package dbpg

import (
	"strings"

	"github.com/google/uuid"

	"github.com/Sourceware-Lab/realquick/config"
)

func Setup() string {
	config.LoadConfig()

	Open(config.Config.DatabaseDSN)

	dbDSNString := config.Config.DatabaseDSN
	dbDSN := config.DBDSN{}
	dbDSN.ParseDSN(dbDSNString)

	dbName := strings.ReplaceAll("testdb-"+uuid.New().String(), "-", "")
	dbDSN.DBName = dbName

	CreateDB(dbName)

	Open(dbDSN.String())
	RunMigrations()

	return dbName
}

func Teardown(dbName string) {
	Open(config.Config.DatabaseDSN)
	DeleteDB(dbName)
}
