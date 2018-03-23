package common

import (
	"fmt"

	"github.com/astaxie/beego"
	weixin "github.com/chanyipiaomiao/weixin-kit"
)

var (
	corpID            = beego.AppConfig.String("weixin::corpID")
	appSecret         = beego.AppConfig.String("weixin::warningAppSecret")
	accessTokenAPI    = beego.AppConfig.String("weixin::accessTokenAPI")
	sendMessageAPIURL = beego.AppConfig.String("weixin::sendMessageAPIURL")
)

// SendWeixinMessage 发送消息
func SendWeixinMessage(msgType, msg, toTag, toUser, toParty string) (bool, error) {
	agentID, err := beego.AppConfig.Int64("weixin::warningAppAgentID")
	if err != nil {
		return false, fmt.Errorf("get agentID from app.conf error: %s ", err)
	}

	message := &weixin.Message{
		MsgType: msgType, // 目前只支持发送文本消息
		ToTag:   toTag,   // ToTag 是在企业微信后台定义的标签ID，标签里面可以包含很多人,多个请用|分开
		ToUser:  toUser,  // ToUser 是企业微信后台看到的用户的ID，多个请用|分开
		ToParty: toParty, // ToParty 是企业微信后台看到的部门的ID，多个请用|分开
		AgentID: agentID, // 企业应用的id，整型。可在应用的设置页面查看
		Safe:    0,       // 表示是否是保密消息，0表示否，1表示是，默认0
		Text: &weixin.Text{
			Content: msg,
		},
	}

	client := &weixin.Client{
		AccessTokenAPI: accessTokenAPI,
		APIURL:         sendMessageAPIURL,
		CorpID:         corpID,
		CorpSecret:     appSecret,
		Message:        message,
	}
	_, err = client.SendMessage()
	if err != nil {
		return false, err
	}
	return true, nil
}
