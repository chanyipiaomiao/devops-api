package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"

	"devops-api/common"
)

func getUniqueIDName() string {

	// 从配置文件中获取 RequestID或者TraceID,如果配置文件中没有配置默认就是 RequestId
	uniqueIDName := beego.AppConfig.String("uniqueIDName")
	if uniqueIDName == "" {
		uniqueIDName = "RequestID"
	}
	return uniqueIDName
}

var (
	UniQueIDName   = getUniqueIDName()
	NeedTokenError = "need DEVOPS-API-TOKEN header"
	TokenAuthError = "DEVOPS-API-TOKEN auth fail"
)

type StringMap map[string]interface{}

// BaseController 基础控制器
type BaseController struct {
	beego.Controller
}

func (b *BaseController) log(msg StringMap) StringMap {
	if _, ok := msg["requestId"]; !ok {
		msg["requestId"] = b.Data[UniQueIDName]
	}

	if _, ok := msg["clientIP"]; !ok {
		msg["clientIP"] = b.Data["RemoteIP"]
	}

	if _, ok := msg["token"]; !ok {
		msg["token"] = b.Data["token"]
	}
	return msg
}

func (b *BaseController) LogInfo(entryType string, msg StringMap) {
	message := b.log(msg)
	if _, ok := msg["statuscode"]; !ok {
		message["statuscode"] = 0
	}
	common.GetLogger().Info(message, entryType)
}

func (b *BaseController) LogError(entryType string, msg StringMap) {
	message := b.log(msg)
	if _, ok := msg["statuscode"]; !ok {
		message["statuscode"] = 1
	}
	common.GetLogger().Error(message, entryType)
}

func (b *BaseController) json(entryType, errmsg string, statuscode int, data interface{}, isLog bool) {
	msg := map[string]interface{}{
		"entryType":  entryType,
		"requestId":  b.Data[UniQueIDName],
		"errmsg":     errmsg,
		"statuscode": statuscode,
		"data":       data,
	}
	b.Data["json"] = msg
	b.ServeJSON()

	msg["clientIP"] = b.Data["RemoteIP"]
	msg["token"] = b.Data["token"]

	if isLog {
		go func() {
			if statuscode == 1 {
				b.LogError(entryType, msg)
			} else {
				b.LogInfo(entryType, msg)
			}
		}()
	}

}

func (b *BaseController) JsonError(entryType, errmsg string, data interface{}, isLog bool) {
	b.json(entryType, errmsg, 1, data, isLog)
}

func (b *BaseController) JsonOK(entryType string, data interface{}, isLog bool) {
	b.json(entryType, "", 0, data, isLog)
}

// Prepare 覆盖Controller的方法
func (b *BaseController) Prepare() {

	// 获取客户端IP
	b.Data["RemoteIP"] = b.Ctx.Input.IP()

	uniqueID := b.Ctx.Input.Header(UniQueIDName)
	if uniqueID == "" {
		uid, err := uuid.NewV4()
		if err != nil {
			common.GetLogger().Error(map[string]interface{}{
				"entryType": "Get UUID",
			}, fmt.Sprintf("%s", err))
			uniqueID = ""
		} else {
			uniqueID = fmt.Sprintf("%s", uid)
		}
	}
	b.Data[UniQueIDName] = uniqueID

	// 配置文件文件中启用了token功能,才验证token
	if common.EnableToken {

		// 获取 DEVOPS-API-TOKEN 头信息
		token := b.Ctx.Input.Header("DEVOPS-API-TOKEN")
		if token == "" {
			b.JsonError("JWToken Auth", NeedTokenError, StringMap{}, true)
			b.StopRun()
		}
		b.Data["token"] = token

		// 验证 DEVOPS-API-TOKEN 是否有效
		jwtoken, err := common.NewToken()
		if err != nil {
			b.JsonError("JWToken Auth", TokenAuthError, StringMap{}, true)
			b.StopRun()
		}

		// 验证是否是root token 不能使用root token
		isroot, err := jwtoken.IsRootToken(token)
		if err != nil {
			b.JsonError("JWToken Auth", TokenAuthError, StringMap{}, true)
			b.StopRun()
		}
		if isroot {
			b.JsonError("JWToken Auth", TokenAuthError, StringMap{}, true)
			b.StopRun()
		}

		_, err = jwtoken.IsTokenValid(token)
		if err != nil {
			b.JsonError("JWToken Auth", TokenAuthError, StringMap{}, true)
			b.StopRun()
		}
	}

}

// PasswordController 密码管理控制器
type PasswordController struct {
	BaseController
}

// MD5Controller MD5管理控制器
type MD5Controller struct {
	BaseController
}

// EmailController  发送邮件控制器
type EmailController struct {
	BaseController
}

// VersionController 程序自身版本管理控制器
type VersionController struct {
	BaseController
}

// TwoStepAuthController 二步验证控制器
type TwoStepAuthController struct {
	BaseController
}

// StorePasswordController 密码管理控制器
type StorePasswordController struct {
	BaseController
}

// WeixinController 发送微信消息控制器
type WeixinController struct {
	BaseController
}

// DingdingController 发送钉钉消息控制器
type DingdingController struct {
	BaseController
}

// HolidayController 节假日工作日判断
type HolidayController struct {
	BaseController
}

// QueryIPController IP地址查询
type QueryIPController struct {
	BaseController
}

// PhoneController 手机归属地查询
type PhoneController struct {
	BaseController
}
