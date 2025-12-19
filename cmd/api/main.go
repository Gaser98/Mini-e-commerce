package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"db-design-project/internal/api"
	"db-design-project/internal/db"
)

func main() {
	dbConn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(dbConn)
	router := gin.Default()

	api.RegisterRoutes(router, queries)

	router.Run(":8080")
}
