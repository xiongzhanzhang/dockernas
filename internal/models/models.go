package models

import (
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"tinycloud/internal/config"
)

var db *gorm.DB = nil
var lock sync.Mutex

func GetDb() *gorm.DB {
	lock.Lock()
	defer lock.Unlock()

	if db != nil {
		return db
	}

	dbFilePath := config.GetDBFilePath()

	_db, err := gorm.Open("sqlite3", dbFilePath)
	if err != nil || _db == nil {
		log.Println(err)
		panic(err)
	}

	_db.SingularTable(true)
	_db.AutoMigrate(&ParamItem{})
	_db.AutoMigrate(&Instance{})
	_db.AutoMigrate(&EventLog{})
	_db.AutoMigrate(&InstancePort{})

	_db.DB().SetMaxIdleConns(3)
	_db.DB().SetMaxOpenConns(20)

	db = _db

	return db
}
