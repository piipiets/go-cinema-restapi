package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Initiator() {
	viper.AutomaticEnv()

	viper.BindEnv("db.postgres.db_host", "PGHOST")
	viper.BindEnv("db.postgres.db_port", "PGPORT")
	viper.BindEnv("db.postgres.db_name", "PGDATABASE")
	viper.BindEnv("db.postgres.db_user", "PGUSER")
	viper.BindEnv("db.postgres.db_password", "PGPASSWORD")
	viper.BindEnv("db.postgres.db_engine", "DB_ENGINE")

	requiredEnv := []string{
		"db.postgres.db_host",
		"db.postgres.db_port",
		"db.postgres.db_name",
		"db.postgres.db_user",
		"db.postgres.db_password",
		"db.postgres.db_engine",
	}

	for _, env := range requiredEnv {
		if viper.GetString(env) == "" {
			panic(fmt.Sprintf("%s environment variable is required", env))
		}
	}

	fmt.Println("Successfully read environment config")
}
