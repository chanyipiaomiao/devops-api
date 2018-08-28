package controllers

import (
	"devops-api/common"
	"fmt"
)

var (
	dingdingEntryType = "SendDingdingMessage"
)

// SendMessage 发送钉钉消息
func (d *DingdingController) SendMessage() {
	msgType := d.GetString("msgType")
	msg := d.GetString("msg")
	title := d.GetString("title")
	robotURL := d.GetString("url")

	_, err := common.SendByDingTalkRobot(msgType, msg, title, robotURL)
	if err != nil {
		d.JsonError(dingdingEntryType, fmt.Sprintf("%s", err),
			StringMap{"result": "send fail"}, true)
		return
	}
	d.JsonOK(dingdingEntryType, StringMap{"result": "send ok"}, true)
}
