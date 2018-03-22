package controllers

import (
	"devops-api/common"
	"fmt"
)

// SendMessage 发送钉钉消息
func (d *DingdingController) SendMessage() {
	msgType := d.GetString("msgType")
	msg := d.GetString("msg")
	title := d.GetString("title")
	robotURL := d.GetString("url")

	requestID := d.Data["RequestID"].(string)
	sendDingdingMessageLog := map[string]interface{}{
		"entryType": "SendDingdingMessage",
		"requestId": requestID,
	}
	_, err := common.SendByDingTalkRobot(msgType, msg, title, robotURL)
	if err != nil {
		sendDingdingMessageLog["statuscode"] = 1
		sendDingdingMessageLog["errmsg"] = fmt.Sprintf("%s", err)
		common.GetLogger().Error(sendDingdingMessageLog, "发送钉钉消息")
		d.Data["json"] = sendDingdingMessageLog
		d.ServeJSON()
		return
	}
	sendDingdingMessageLog["statuscode"] = 0
	sendDingdingMessageLog["errmsg"] = ""
	sendDingdingMessageLog["result"] = "发送成功"
	common.GetLogger().Info(sendDingdingMessageLog, msg)
	d.Data["json"] = sendDingdingMessageLog
	d.ServeJSON()
}
