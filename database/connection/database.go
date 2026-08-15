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

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("db.postgres.db_host"),
		viper.GetInt("db.postgres.db_port"),
		viper.GetString("db.postgres.db_user"),
		viper.GetString("db.postgres.db_password"),
		viper.GetString("db.postgres.db_name"),
	)

	DBConnections, err = sql.Open(dbEngine, dsn)

	// check connection
	err = DBConnections.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
}
