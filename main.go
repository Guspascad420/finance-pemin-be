package main

import (
	"finance-be/controllers"
	"finance-be/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "POST", "GET", "OPTIONS", "DELETE"},
		AllowHeaders:  []string{"Content-Type", "X-Requested-With", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	godotenv.Load()

	database.Connect()
	database.Migrate()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	api := r.Group("/api")
	api.POST("/user/register", controllers.RegisterUser)
	api.POST("/user/login", controllers.GenerateToken)
	api.GET("/user/profile", controllers.GetUserProfile)
	api.GET("/balance", controllers.GetBalances)
	api.POST("/balance/add", controllers.CreateNewBalance)
	api.DELETE("/balance/:id", controllers.DeleteBalance)
	api.GET("/transaction", controllers.GetTransactions)
	api.POST("/transaction/add", controllers.CreateNewTransaction)
	api.PUT("/balance/update", controllers.UpdateBalance)
	api.DELETE("/transaction/:id", controllers.DeleteTransaction)

	r.Run()
}
