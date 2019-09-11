package db

import (
	"../utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type Model struct {
	ID        string         `structs:"id" gorm:"primary_key" form:"id" json:"id"`
	CreatedAt utils.DateTime `structs:"-" json:"createdAt" gorm:"type:datetime"`
	UpdatedAt utils.DateTime `structs:"-" json:"updatedAt" gorm:"type:datetime"`
	// DeletedAt *time.Time `sql:"index" structs:"-"`
}

var DB *gorm.DB

func Init() (err error) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTablename string) string {
		return "t_" + defaultTablename
	}
	switch utils.DBType() {
	case "mysql":
		DB, err = gorm.Open("mysql", "<user>:<password>/<database>?charset=utf8&parseTime=True&loc=Local")
	case "sqlite":
	default:
		{
			dbFile := utils.DBFile()
			log.Println("db file -->", utils.DBFile())
			DB, err = gorm.Open("sqlite3", fmt.Sprintf("%s?loc=Asia/Shanghai", dbFile))
			// Sqlite cannot handle concurrent writes, so we limit sqlite to one connection.
			// see https://github.com/mattn/go-sqlite3/issues/274
			DB.DB().SetMaxOpenConns(1)
		}
	}

	if err != nil {
		return
	}

	DB.SetLogger(DefaultGormLogger)
	DB.LogMode(false)
	return
}

func Close() {
	if DB != nil {
		DB.Close()
		DB = nil
	}
}
