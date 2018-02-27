package common

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/hltool"
)

// InitLog 初始化日志
func InitLog() {
	var logpath string
	if LogPathFromCli == "" {
		logpath = beego.AppConfig.String("log::logPath")
	} else {
		logpath = LogPathFromCli
	}
	hlog, err := hltool.NewHLog(logpath)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	Logger, err = hlog.GetLogger()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}

// GetLogger 返回Logger
func GetLogger() *hltool.HLogger {
	return Logger
}
