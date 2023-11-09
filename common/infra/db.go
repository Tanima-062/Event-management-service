package infra

import (
	"fmt"
	"net/url"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConnection
func DBCon() (*gorm.DB, error) {

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	protocol := fmt.Sprintf("tcp(%s:%s)", host, port)

	connectInfo := fmt.Sprintf("%s:%s@%s/%s?parseTime=True&loc=%s", username, password, protocol, database, url.PathEscape("Asia/Dhaka"))
	DB, err := gorm.Open(mysql.Open(connectInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return DB, nil
}