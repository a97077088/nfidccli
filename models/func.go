package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-adodb"
	"github.com/satori/go.uuid"
	"time"
)

var db *gorm.DB

//接口
func Ctx()*gorm.DB{
	return db
}

//初始化数据库
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

//生成任务编号
func Build_taskid()string{
	r:=fmt.Sprintf("国抽导入%d",time.Now().UnixNano())
	return r
}
//生成数据库id
func Build_id()string{
	ud:=uuid.NewV4().String()
	return ud
}