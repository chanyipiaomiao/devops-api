package controllers

import (
	"devops-api/common"
	"fmt"
)

var (
	weixinEntryType = "SendWeixinMessage"
)

// SendMessage 发送消息
func (w *WeixinController) SendMessage() {
	msgType := w.GetString("msgType")
	toTag := w.GetString("toTag")
	toUser := w.GetString("toUser")
	toParty := w.GetString("toParty")
	msg := w.GetString("msg")

	_, err := common.SendWeixinMessage(msgType, msg, toTag, toUser, toParty)
	if err != nil {
		w.JsonError(weixinEntryType, fmt.Sprintf("%s", err), StringMap{"result": "send fail"}, true)
		return
	}

	w.JsonOK(weixinEntryType, StringMap{"result": "send ok"}, true)
}
