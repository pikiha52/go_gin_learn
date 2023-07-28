package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"clean_architecture_gin/api/routes"
	"clean_architecture_gin/config"
	"clean_architecture_gin/pkg/user"

)

func main() {
	app := gin.Default()

	connection := ConnectDB()
	userRepo := user.NewRepo(connection)
	userService := user.NewService(userRepo)

	api := app.Group("/api")
	routes.Routes(api, userService)

	app.Run(":3000")
}

func ConnectDB() *gorm.DB {
	var DB *gorm.DB

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Error parsing DB port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to open database!")
	}

	fmt.Println("Connection Opened to Database: ")

	// fmt.Println("Migration table users completed")

	return DB
}
