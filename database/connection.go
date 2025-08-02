package database

import(
	"fmt"
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"github.com/valentinoanugerah/crud_go/models"
)

var DB *gorm.DB

func Connect(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env ")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"),

)

DB, err = gorm.Open(mysql.Open(dsn) &gorm.Config{})

if err != nil {
	log.Fatal("Failed to connect to database")
}

DB.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Supplier{}, &models.Sale{}, &models.Sale_item{}, &models.Role{}, &models.Purchase{}, &models.Purchase_item{}, &models.Product{}, &models.Customer{}, &models.Audit_log{})
}