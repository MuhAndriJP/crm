package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init() {
	host := fmt.Sprint(os.Getenv("POSTGRE_DB_HOST"))
	user := fmt.Sprint(os.Getenv("POSTGRE_DB_USER"))
	password := fmt.Sprint(os.Getenv("POSTGRE_DB_PASSWORD"))
	port := fmt.Sprint(os.Getenv("POSTGRE_DB_PORT"))
	dbName := fmt.Sprint(os.Getenv("POSTGRE_DB_DATABASE"))

	urlConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ", host, port, user, password, dbName)

	driverName := os.Getenv("POSTGRE_DB_DRIVER")
	db, err := gorm.Open(driverName, urlConnection)
	if err != nil {
		log.Fatalf("⇨ %s Data source %s:%s , Failed : %s \n", driverName, host, port, err.Error())
	}

	fmt.Printf("⇨ %s Data source %s:%s , Successfully connected! \n", driverName, host, port)
	DB = db
}
