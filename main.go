package main

import (
	"database/sql"

	"github.com/piipiets/go-cinema-restapi/config"
	"github.com/piipiets/go-cinema-restapi/database/connection"
	"github.com/piipiets/go-cinema-restapi/database/migration"
	"github.com/piipiets/go-cinema-restapi/handler"
	"github.com/piipiets/go-cinema-restapi/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Initiator()

	connection.Initiator()
	defer connection.DBConnections.Close()

	migration.Initiator(connection.DBConnections)

	InitiateRouter(connection.DBConnections)
}

func InitiateRouter(db *sql.DB) {
	router := gin.Default()

	// Service
	cinemaService := service.NewCinemaService(db)

	// Handler
	cinemaHandler := handler.NewCinemaHandler(
		cinemaService,
	)

	router.POST(
		"/cinema",
		cinemaHandler.CreateCinema,
	)

	router.Run(":8080")
}
