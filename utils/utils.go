package utils

import (
	"github.com/go-pg/pg"
	"github.com/joho/godotenv"
	"os"
)

// Decimals
const (
	WavesDecimal = 1e8
	EthDecimal   = 1e18
)

func GetDBCredentials () (string, string, string, string, string) {
	envLoadErr := godotenv.Load(".env")
	if envLoadErr != nil {
		_ = godotenv.Load(".env.example")
	}

	dbhost := "localhost"
	dbport := "5432"
	if os.Getenv("DB_HOST") != "" {
		dbhost = os.Getenv("DB_HOST")
	}
	if os.Getenv("DB_PORT") != "" {
		dbport = os.Getenv("DB_PORT")
	}
	dbuser := os.Getenv("DB_USERNAME")
	dbpass := os.Getenv("DB_PASS")
	dbdatabase := os.Getenv("DB_NAME")

	return dbhost, dbport, dbuser, dbpass, dbdatabase
}

func ConnectToPG () *pg.DB {
	dbhost, dbport, dbuser, dbpass, dbdatabase := GetDBCredentials()

	db := pg.Connect(&pg.Options{
		Addr: dbhost + ":" + dbport,
		User:     dbuser,
		Password: dbpass,
		Database: dbdatabase,
		//TLSConfig: &tls.Config{
		//	InsecureSkipVerify: true,
		//},
	})
	return db
}