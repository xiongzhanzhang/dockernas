package models

import (
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"tinycloud/internal/config"
	"tinycloud/internal/utils"
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
	isInited := utils.IsFileExist(dbFilePath) != false

	_db, err := gorm.Open("sqlite3", dbFilePath)
	if err != nil || _db == nil {
		log.Println(err)
		panic(err)
	}

	_db.SingularTable(true)
	_db.DB().SetMaxIdleConns(3)
	_db.DB().SetMaxOpenConns(20)

	if isInited == false {
		_db.AutoMigrate(&ParamItem{})
	}
	db = _db

	return db
}
