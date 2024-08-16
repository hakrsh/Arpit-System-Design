package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/harikrishnanum/pg_read_repilic/db"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	masterDb  *db.Queries
	replicaDb *db.Queries
	server    *gin.Engine
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: %v", err)
	}
	connectDb()
	server = gin.New()
	setupRoutes()

}

func getConnection(dbConnStr string) *db.Queries {
	dbDriver := os.Getenv("POSTGRES_DRIVER")
	dbConn, err := sql.Open(dbDriver, dbConnStr)
	if err != nil {
		log.Fatal("Error connecing to master database: %v", err)
	}
	return db.New(dbConn)
}
func connectDb() {
	masterDbConnStr := os.Getenv("POSTGRES_PRIMARY_SOURCE")
	replicaDbConnStr := os.Getenv("POSTGRES_REPILICA_SOURCE")

	if masterDbConnStr == "" || replicaDbConnStr == "" {
		log.Fatalf("Database connection strings are not set properly")
	}

	masterDb = getConnection(masterDbConnStr)
	log.Println("Master DB connected successfully")

	replicaDb = getConnection(replicaDbConnStr)
	log.Println("Replica DB connected successfully")
}
func setupRoutes() {
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "PG read replica test",
		})
	})
	server.POST("/user", func(ctx *gin.Context) {
		name := faker.Name()
		user, err := masterDb.CreateUser(ctx, name)
		if err != nil {
			log.Printf("Failed to create user: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create user",
			})
			return
		}
		ctx.JSON(http.StatusOK, user)
	})
	server.GET("/users", func(ctx *gin.Context) {
		users, err := replicaDb.ListUsers(ctx)
		if err != nil {
			log.Printf("Failed to fetch users: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch users",
			})
			return
		}
		ctx.JSON(http.StatusOK, users)
	})
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Fatal(server.Run(":" + port))
}
