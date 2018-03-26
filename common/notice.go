package common

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/hltool"
)

// SendByDingTalkRobot 通过钉钉发送消息通知
func SendByDingTalkRobot(messageType, message, title, robotURL string) (bool, error) {
	var url string
	if robotURL == "" {
		url = beego.AppConfig.String("dingTalkRobot")
	} else {
		url = robotURL
	}

	dingtalk := &hltool.DingTalkClient{
		RobotURL: url,
		Message: &hltool.DingTalkMessage{
			Type:    messageType,
			Message: message,
			Title:   title,
		},
	}
	ok, err := hltool.SendMessage(dingtalk)
	if err != nil {
		dingFields := map[string]interface{}{
			"entryType":     "DingTalkRobot",
			"dingTalkRobot": url,
		}
		Logger.Error(dingFields, fmt.Sprintf("发送钉钉通知失败了: %s", err))
		return false, err
	}
	return ok, nil
}

// SendByEmail 通过Email发送消息通知
func SendByEmail(subject, content, contentType, attach string, to, cc []string) (bool, error) {
	username := beego.AppConfig.String("email::username")
	host := beego.AppConfig.String("email::host")
	password := beego.AppConfig.String("email::password")
	port, err := beego.AppConfig.Int("email::port")
	if err != nil {
		confFields := map[string]interface{}{
			"entryType": "Parse Configure File",
		}
		Logger.Error(confFields, fmt.Sprintf("从配置文件中解析邮件端口失败: %s", err))
		return false, err
	}

	message := hltool.NewEmailMessage(username, subject, contentType, content, attach, to, cc)
	email := hltool.NewEmailClient(host, username, password, port, message)
	ok, err := hltool.SendMessage(email)
	if err != nil {
		emailFields := map[string]interface{}{
			"entryType": "SendMail",
			"mail": map[string]interface{}{
				"Username":    username,
				"Host":        host,
				"Port":        port,
				"ContentType": contentType,
				"Attach":      attach,
				"To":          to,
				"Cc":          cc,
			},
		}
		Logger.Error(emailFields, fmt.Sprintf("发送邮件失败了: %s", err))
		return false, err
	}
	return ok, nil
}
