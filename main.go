package main

import (
	"fmt"
	"log"

	"github.com/SitanayaIvan/testing-deploy/config"
	"github.com/SitanayaIvan/testing-deploy/database"
	"github.com/SitanayaIvan/testing-deploy/route"
	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
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

	// Migrate DB
	database.MigrateDb(db)

	// Routing
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(cors.AllowAll())

	v1 := r.Group("/api")

	// All route
	route.Routes(r, v1, db)

	// Midtrans

	// Midtrans connect to server
	midtrans.ServerKey = "SB-Mid-server-m4Qo9eTnsC0OQdJmI8EeXnWW"
	midtrans.Environment = midtrans.Sandbox

	// Midtrans request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "7",
			GrossAmt: 20000,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: "Ivan",
			LName: "Junior",
			Email: "ivansitanaya@gmail.com",
			Phone: "082238366883",
		},
		Callbacks: &snap.Callbacks{
			Finish: "http://localhost:8080/api/user",
		},
	}

	// Midtrans create transaction
	snapResp, _ := snap.CreateTransaction(req)
	fmt.Println("Response:", snapResp)

	// Run App
	log.Println("[SERVER] Server Started!")
	r.Run(":" + conf.App.Port)

}
