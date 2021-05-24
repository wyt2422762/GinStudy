package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wyt/GinStudy/conf"
)

func init() {
	var sb strings.Builder
	sb.WriteString(conf.DbUesr)
	sb.WriteString(":")
	sb.WriteString(conf.DbPwssword)
	sb.WriteString("@tcp(")
	sb.WriteString(conf.DbHost)
	sb.WriteString(")/")
	sb.WriteString(conf.DbName)
	sb.WriteString("?charset=utf8")

	fmt.Println("连接mysql数据库")
	DB, err := sql.Open(conf.DbType, sb.String())
	if err != nil {
		panic(err.Error())
	}
	//最大空闲连接数
	DB.SetMaxIdleConns(10)
	//最大连接数
	DB.SetMaxOpenConns(10)
	//最大空闲时间
	DB.SetConnMaxIdleTime(time.Second * 30)
}
