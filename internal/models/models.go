package models

import (
	"log"
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"dockernas/internal/config"
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

	newLogger := logger.New(
		log.New(log.Writer(), "\n", 0),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	_db, err := gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{
		Logger:         newLogger,
		NamingStrategy: schema.NamingStrategy{TablePrefix: "t_", SingularTable: true},
	})
	if err != nil || _db == nil {
		log.Println(err)
		panic(err)
	}

	_db.AutoMigrate(&ParamItem{})
	_db.AutoMigrate(&Instance{})
	_db.AutoMigrate(&EventLog{})
	_db.AutoMigrate(&InstancePort{})
	_db.AutoMigrate(&ContainerStat{})
	_db.AutoMigrate(&HttpProxyConfig{})
	_db.AutoMigrate(&Subscribe{})

	db = _db

	return db
}
