package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Initiator() {
	viper.AutomaticEnv()

	viper.BindEnv("PGHOST")
	viper.BindEnv("PGPORT")
	viper.BindEnv("PGDATABASE")
	viper.BindEnv("PGUSER")
	viper.BindEnv("PGPASSWORD")
	viper.BindEnv("DB_ENGINE")

	requiredEnv := []string{
		"PGHOST",
		"PGPORT",
		"PGDATABASE",
		"PGUSER",
		"PGPASSWORD",
		"DB_ENGINE",
	}

	for _, env := range requiredEnv {
		if viper.GetString(env) == "" {
			panic(fmt.Sprintf("%s environment variable is required", env))
		}
	}

	fmt.Println("Successfully read environment config")
}
