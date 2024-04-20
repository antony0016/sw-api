package inits

import (
	"strings"

	"github.com/antony0016/sw-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ToDSN(host string, port string, username string, password string, dbName string) string {
	DBStrs := []string{
		"host=" + host,
		"port=" + port,
		"user=" + username,
		"password=" + password,
		"dbname=" + dbName,
		"sslmode=disable",
	}
	return strings.Join(DBStrs, "\n")
}

func DBInit(env config.Env) (*gorm.DB, error) {
	dsn := ToDSN(env.DB_HOST, env.DB_PORT, env.DB_USERNAME, env.DB_PASSWORD, env.DB_DATABASE)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
