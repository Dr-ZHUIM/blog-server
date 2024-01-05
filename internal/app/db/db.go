package CustomMysql

import (
	mysqlConfig "article-server/config"
	DTOs "article-server/internal/app/dtos"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	username = mysqlConfig.Username
	password = mysqlConfig.Password
	hostname = mysqlConfig.Hostname
	port     = mysqlConfig.Port
	dbname   = mysqlConfig.DBName
)

var DB *gorm.DB

func Connect() {
	dsn := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + dbname

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   true,
			SingularTable: true,
		},
		// Logger: Log.GormLogger,
	})

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	// Auto-migrate the User model
	db.AutoMigrate(&DTOs.Article{}, &DTOs.User{})
}
