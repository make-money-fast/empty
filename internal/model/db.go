package model

import (
	"github.com/clearcodecn/sqlite"
	"github.com/make-money-fast/empty/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	o   sync.Once
	_db *gorm.DB
)

func Load(conf *config.Config) *gorm.DB {
	o.Do(func() {
		var (
			db  *gorm.DB
			err error
		)
		switch conf.DB.Driver {
		case "sqlite":
			db, err = gorm.Open(sqlite.Open(conf.DB.Dsn), &gorm.Config{})
		case "mysql":
			db, err = gorm.Open(mysql.Open(conf.DB.Dsn), &gorm.Config{})
		default:
			log.Fatal("database driver not set")
		}

		if err != nil {
			log.Fatal("初始化db失败")
			return
		}
		_db = db

		syncTable(_db)
	})

	return _db
}

func syncTable(db *gorm.DB) {
	if err := db.AutoMigrate(
		&Article{},
	); err != nil {
		log.Fatal(err)
	}
}
