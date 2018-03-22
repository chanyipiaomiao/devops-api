package controllers

import (
	"devops-api/common"
	"fmt"
)

// SendMessage 发送消息
func (w *WeixinController) SendMessage() {
	msgType := w.GetString("msgType")
	toTag := w.GetString("toTag")
	toUser := w.GetString("toUser")
	toParty := w.GetString("toParty")
	msg := w.GetString("msg")

	requestID := w.Data["RequestID"].(string)
	sendWeixinMessageLog := map[string]interface{}{
		"entryType": "SendWeixinMessage",
		"requestId": requestID,
	}
	_, err := common.SendWeixinMessage(msgType, msg, toTag, toUser, toParty)
	if err != nil {
		sendWeixinMessageLog["statuscode"] = 1
		sendWeixinMessageLog["errmsg"] = fmt.Sprintf("%s", err)
		common.GetLogger().Error(sendWeixinMessageLog, "发送微信消息")
		w.Data["json"] = sendWeixinMessageLog
		w.ServeJSON()
		return
	}
	sendWeixinMessageLog["statuscode"] = 0
	sendWeixinMessageLog["errmsg"] = ""
	sendWeixinMessageLog["result"] = "发送成功"
	common.GetLogger().Info(sendWeixinMessageLog, msg)
	w.Data["json"] = sendWeixinMessageLog
	w.ServeJSON()
	return
}
