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
	text := w.GetString("text")
	agentID, err := w.GetInt64("agentID")
	if err != nil {
		agentID = 0
	}

	requestID := w.Data["RequestID"].(string)
	sendWeixinMessageLog := map[string]interface{}{
		"entryType": "SendWeixinMessage",
		"requestId": requestID,
	}
	_, err = common.SendWeixinMessage(msgType, text, toTag, toUser, toParty, agentID)
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
	common.GetLogger().Info(sendWeixinMessageLog, text)
	w.Data["json"] = sendWeixinMessageLog
	w.ServeJSON()
	return
}
