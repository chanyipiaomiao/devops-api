package common

import (
	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/hltool"
)

var (

	// Logger 日志对象
	Logger *hltool.HLogger

	// LogPathFromCli 从命令行传入日志路径
	LogPathFromCli string

	// EnableToken 读取是否启用 Token 认证配置
	EnableToken = getEnableToken()
)

// getEnableToken 读取是否启用 Token 认证配置
func getEnableToken() bool {
	enableToken, err := beego.AppConfig.Bool("security::enableToken")
	if err != nil {
		enableToken = true
	}

	return enableToken
}
