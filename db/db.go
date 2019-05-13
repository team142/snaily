package db

import (
	"flag"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

var DatabaseHost = flag.String("pghost", "localhost", "PG hostname")
var DatabaseUser = flag.String("pguser", "snaily", "PG username")
var DatabasePassword = flag.String("pgpassword", "snaily", "PG password")
var DatabaseDatabase = "madast"
var Port = flag.Uint64("pgport", 5000, "PG port")

func Connect() (conn *pgx.Conn, err error) {
	conn, err = pgx.Connect(pgx.ConnConfig{
		Host:     *DatabaseHost,
		Port:     uint16(*Port),
		User:     *DatabaseUser,
		Password: *DatabasePassword,
		Database: DatabaseDatabase,
	})
	if err != nil {
		logrus.Fatal("Unable to connection to database: %v\n", err)
	}
	return
}
