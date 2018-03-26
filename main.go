package main

import (
	"devops-api/common"
	_ "devops-api/routers"

	"github.com/astaxie/beego"
)

func main() {

	// 是否启用 定时生成验证密码 功能
	if ok, _ := beego.AppConfig.Bool("authpassword::enableCrontabAuthPassword"); ok {
		common.CronGenAuthPassword()
	}

	// 是否启用 定时清除验证密码 功能
	if ok, _ := beego.AppConfig.Bool("authpassword::enableManualGenAuthPassword"); ok {
		common.CronClearAuthPassword()
	}

	// 初始化获取命令行参数
	common.InitCli()

}
