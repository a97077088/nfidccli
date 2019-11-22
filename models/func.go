package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-adodb"
)

var db *gorm.DB

//接口
func Ctx()*gorm.DB{
	return db
}

func InitDb(host string,port int,user,pass string,dbname string)error{
	connectionString := fmt.Sprintf("Provider=SQLOLEDB;Data Source=%s,%d;user id=%s;password=%s;Initial Catalog=%s;",
		host,port,user, pass, dbname,)
	var err error
	db,err=gorm.Open("adodb",connectionString)
	if err != nil {
		return err
	}
	// 设置空闲连接池中的最大连接数
	db.DB().SetMaxIdleConns(10)
	// 设置到数据库的最大打开连接数
	db.DB().SetMaxOpenConns(100)

	db.LogMode(true)
	return nil
}