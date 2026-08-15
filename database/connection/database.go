package connection

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

var (
	DBConnections *sql.DB
	err           error
)

func Initiator() {
	dbEngine := viper.GetString("db.postgres.db_engine")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		viper.GetString("PGHOST"),
		viper.GetString("PGPORT"),
		viper.GetString("PGUSER"),
		viper.GetString("PGPASSWORD"),
		viper.GetString("PGDATABASE"),
	)

	DBConnections, err = sql.Open(dbEngine, dsn)

	// check connection
	err = DBConnections.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
}
