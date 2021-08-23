package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

//InitDb initializes the connection to the db and the db object, which through we interact with the DB
func Init() *gorm.DB{
	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Recipe{})

	return db
}

type Recipe struct {

}
