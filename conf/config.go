package conf

import (
	"fmt"
	"github.com/go-ini/ini"
	"time"
)

//这个文件用来读取ini文件中的信息

var (
	Cfg *ini.File

	RunMode string

	PageSize  int
	JwtSecret string

	//端口
	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration

	DbType     string
	DbUesr     string
	DbPwssword string
	DbHost     string
	DbName     string

	LogPath  string
	LogLevel string
)

func init() {

	fmt.Println("config.go init 方法")

	var err error
	Cfg, err = ini.Load("conf/config.ini")
	if err != nil {
		panic("读取启动配置文件失败")
	}

	loadApp()
	loadBase()
	loadServer()
	loadDataBase()
	loadLog()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadApp() {
	PageSize = Cfg.Section("app").Key("PAGE_SIZE").MustInt(10)
	JwtSecret = Cfg.Section("app").Key("JWT_SECRET").MustString("")
}

func loadServer() {
	HttpPort = Cfg.Section("server").Key("HTTP_PORT").MustInt(8080)
	ReadTimeOut = time.Duration(Cfg.Section("server").Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(Cfg.Section("server").Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadDataBase() {
	DbType = Cfg.Section("dataBase").Key("DbType").MustString("")
	DbName = Cfg.Section("dataBase").Key("DbName").MustString("")
	DbUesr = Cfg.Section("dataBase").Key("DbUesr").MustString("")
	DbPwssword = Cfg.Section("dataBase").Key("DbPwssword").MustString("")
	DbHost = Cfg.Section("dataBase").Key("DbHost").MustString("")
	DbHost = Cfg.Section("dataBase").Key("DbHost").MustString("")
}

func loadLog() {
	LogPath = Cfg.Section("log").Key("LogPath").MustString("")
	LogLevel = Cfg.Section("log").Key("LogLevel").MustString("")
}
