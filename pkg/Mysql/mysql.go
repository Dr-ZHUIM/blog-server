package CustomMysql

import (
	mysqlConfig "article-server/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username = mysqlConfig.Username
	password = mysqlConfig.Password
	hostname = mysqlConfig.Hostname
	port     = mysqlConfig.Port
	dbname   = mysqlConfig.DBName
)

func Connect() {
	dsn := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + dbname
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Auto-migrate the User model
	// db.AutoMigrate(&User{})
}
