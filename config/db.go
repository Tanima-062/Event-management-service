package infra

import (
	"fmt"
	"net/url"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConnection
func DBCon() (*gorm.DB, error) {

	username := "root"
	password := ""
	database := "event_management_system"
	host := "127.0.0.1"
	port := "3306"
	protocol := fmt.Sprintf("tcp(%s:%s)", host, port)

	connectInfo := fmt.Sprintf("%s:%s@%s/%s?parseTime=True&loc=%s", username, password, protocol, database, url.PathEscape("Asia/Dhaka"))
	DB, err := gorm.Open(mysql.Open(connectInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return DB, nil
}