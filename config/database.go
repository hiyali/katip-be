package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB() *gorm.DB {
	cf := Get() // config from config.yml

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true", cf.DB.User, cf.DB.Password, cf.DB.Name))
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	if err != nil {
		fmt.Printf("%v\n", err)
		panic("failed to connect database")
	}

	// gorm DB settings ----------------------------------------------
	// Disable table name's pluralization globally
	// db.SingularTable(true) // if set this to true, `User`'s default table name will be `user`, table name setted with `TableName` won't be affected

	// Enable Logger, show detailed log
	// Disable Logger, don't show any log
	db.LogMode(true)

	// Debug a single operation, show detailed log for this operation
	// db.Debug().Where("name = ?", "jinzhu").First(&User{})

	// db.SetLogger(gorm.Logger{revel.TRACE})
	// db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	// Ping
	// db.DB().Ping()

	// Connection Pool
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// gorm DB settings ----------------------------------------------

	// gorm Transaction example ----------------------------------------------
	/*
	  // begin a transaction
	  tx := db.Begin()
	  // do some database operations in the transaction (use 'tx' from this point, not 'db')
	  tx.Create(...)
	  // ...
	  // rollback the transaction in case of error
	  tx.Rollback()
	  // Or commit the transaction
	  tx.Commit()
	  // gorm Transaction example ----------------------------------------------
	*/

	return db
}

func GetTable(tableName string) *gorm.DB {
	db := GetDB()
	db.Select(tableName)
	return db
}
