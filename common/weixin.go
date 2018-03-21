package common

import (
	"github.com/astaxie/beego"
	weixin "github.com/chanyipiaomiao/weixin-kit"
)

var (
	corpID            = beego.AppConfig.String("weixin::corpID")
	appSecret         = beego.AppConfig.String("weixin::appSecret")
	accessTokenAPI    = beego.AppConfig.String("weixin::accessTokenAPI")
	sendMessageAPIURL = beego.AppConfig.String("weixin::sendMessageAPIURL")
)

// SendWeixinMessage 发送消息
func SendWeixinMessage(msgType, text, toTag, toUser, toParty string, agentIDFromReq int64) (bool, error) {
	var agentID int64
	if agentIDFromReq == 0 {
		agentIDFromConf, err := beego.AppConfig.Int64("weixin::agentID")
		if err != nil {
			return false, err
		}
		agentID = agentIDFromConf
	} else {
		agentID = agentIDFromReq
	}

	message := &weixin.Message{
		MsgType: msgType, // 目前只支持发送文本消息
		ToTag:   toTag,   // ToTag 是在企业微信后台定义的标签ID，标签里面可以包含很多人,多个请用|分开
		ToUser:  toUser,  // ToUser 是企业微信后台看到的用户的ID，多个请用|分开
		ToParty: toParty, // ToParty 是企业微信后台看到的部门的ID，多个请用|分开
		AgentID: agentID, // 企业应用的id，整型。可在应用的设置页面查看
		Safe:    0,       // 表示是否是保密消息，0表示否，1表示是，默认0
		Text: &weixin.Text{
			Content: text,
		},
	}

	client := &weixin.Client{
		AccessTokenAPI: accessTokenAPI,
		APIURL:         sendMessageAPIURL,
		CorpID:         corpID,
		CorpSecret:     appSecret,
		Message:        message,
	}
	_, err := client.SendMessage()
	if err != nil {
		return false, err
	}
	return true, nil
}
