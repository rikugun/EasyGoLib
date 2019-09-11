package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rikugun/EasyGoLib/utils"
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
	log.Println("db type ", utils.DBType())
	switch utils.DBType() {
	case "mysql":
		DB, err = gorm.Open("mysql", utils.MysqlConnStr())
	case "sqlite":
		fallthrough
	default:
		{
			log.Println("db file -->", utils.DBFile())
			dbFile := utils.DBFile()
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
