package main

import (
	"log"

	"github.com/SitanayaIvan/testing-deploy/config"
	"github.com/SitanayaIvan/testing-deploy/database"
	"github.com/SitanayaIvan/testing-deploy/route"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	conf, _ := config.LoadVariable()

	// Connection to Database
	db, err := database.ConnectDB()
	if err != nil {
		log.Println("[ERROR] Failed to connect database!")
	}
	log.Println("[SUCCESS] Connect database success!")

	// Migrate Db
	// database.MigrateDb(db)

	// Routing
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(cors.AllowAll())

	v1 := r.Group("/api")

	// All route
	route.Routes(r, v1, db)

	// Run App
	log.Println("[SERVER] Server Started!")
	r.Run(":" + conf.App.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
