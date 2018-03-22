package common

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/hltool"
	"github.com/robfig/cron"
)

var (
	// WillAuthPassword 定时生成的密码
	WillAuthPassword string
	passwordFields   = map[string]interface{}{
		"entryType": "GenPassword",
	}
)

// GenPassword 生成验证密码
func GenPassword() map[string]bool {
	WillAuthPassword = hltool.GenRandomString(32, "no")
	info := fmt.Sprintf("请注意,本次验证密码为: %s 生成时间: %s", WillAuthPassword, hltool.GetNowTime2())
	Logger.Info(passwordFields, info)

	sendResult := make(chan bool)
	if ok, _ := beego.AppConfig.Bool("authpassword::enableDingtalkReciveGenPassword"); ok {
		go func(ch chan bool, messageType, message string) {
			ok, _ := SendByDingTalkRobot(messageType, message, "", "")
			ch <- ok
		}(sendResult, "text", info)
	}

	if ok, _ := beego.AppConfig.Bool("authpassword::enableEmailReciveGenPassword"); ok {
		reciver := beego.AppConfig.String("authpassword::genPasswordReciver")
		if reciver == "" {
			emailFields := map[string]interface{}{
				"entryType": "SendMail",
			}
			Logger.Error(emailFields, "邮件收件人为空!!")
		} else {
			reciverEmail := strings.Split(reciver, ",")
			go func(ch chan bool, subject, content, contentType, attach string, to, cc []string) {
				ok, _ := SendByEmail(subject, content, contentType, attach, to, cc)
				ch <- ok
			}(sendResult, "验证密码", info, "text/plain", "", reciverEmail, []string{})
		}
	}

	m := make(map[string]bool)
	for {
		if <-sendResult {
			m["result"] = true
			break
		}
	}
	return m
}

// ClearPassword 清除验证密码
func ClearPassword() {
	if WillAuthPassword != "" {
		temp := WillAuthPassword
		WillAuthPassword = hltool.GenRandomString(32, "no")
		Logger.Info(passwordFields, fmt.Sprintf("定时清除验证密码(%s)成功", temp))
	}
}

// CronGenAuthPassword 定时生成验证密码
func CronGenAuthPassword() {
	c := cron.New()
	c.AddFunc(beego.AppConfig.String("genAuthPasswordCrontab"), func() {
		GenPassword()
	})
	c.Start()
}

// CronClearAuthPassword 定时清除验证密码
func CronClearAuthPassword() {
	c := cron.New()
	c.AddFunc(beego.AppConfig.String("authpassword::clearAuthPasswordCrontab"), func() {
		ClearPassword()
	})
	c.Start()
}

// ManualGenAuthPassword 手动生成验证密码
func ManualGenAuthPassword() map[string]bool {
	return GenPassword()
}

// GetWiillAuthPassword 获取生成的密码
func GetWiillAuthPassword() string {
	return WillAuthPassword
}
