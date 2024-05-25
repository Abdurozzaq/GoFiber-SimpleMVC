package database

import (
	"fmt"
	"github.com/Abdurozzaq/GoFiber-SimpleMVC/src/models"
	"log"
	"strconv"

	"github.com/Abdurozzaq/GoFiber-SimpleMVC/src/config"
	_ "github.com/Abdurozzaq/GoFiber-SimpleMVC/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {

		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	errMigrate := DB.AutoMigrate(&models.GlobalInformation{})
	if errMigrate != nil {
		panic("failed to auto migrate")
	} else {
		fmt.Println("Database Migrated")
	}
}
