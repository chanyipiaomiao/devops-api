package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"

	"devops-api/common"
)

// BaseController 基础控制器
type BaseController struct {
	beego.Controller
}

// Prepare 覆盖Controller的方法
func (b *BaseController) Prepare() {

	// 获取客户端IP
	remoteIP := b.Ctx.Input.IP()
	b.Data["RemoteIP"] = remoteIP

	// 从配置文件中获取 RequestID或者TraceID,如果配置文件中没有配置默认就是 RequestId
	uniqueIDName := beego.AppConfig.String("uniqueIDName")
	if uniqueIDName == "" {
		uniqueIDName = "RequestID"
	}
	uniqueID := b.Ctx.Input.Header(uniqueIDName)
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
	b.Data["RequestID"] = uniqueID

	// 配置文件文件中启用了token功能,才验证token
	if common.EnableToken {

		// 获取 DEVOPS-API-TOKEN 头信息
		token := b.Ctx.Input.Header("DEVOPS-API-TOKEN")
		if token == "" {
			b.Data["json"] = map[string]string{"result": "need DEVOPS-API-TOKEN header", "statuscode": "1"}
			b.ServeJSON()
			b.StopRun()
		}

		// 验证 DEVOPS-API-TOKEN 是否有效
		jwtoken, err := common.NewToken()
		logFields := map[string]interface{}{
			"entryType": "JWToken Auth",
			"requestID": b.Data["RequestID"],
		}
		if err != nil {
			common.GetLogger().Error(logFields, fmt.Sprintf("%s", err))
			b.Data["json"] = map[string]string{"result": "DEVOPS-API-TOKEN auth fail", "statuscode": "1"}
			b.ServeJSON()
			b.StopRun()
		}

		// 验证是否是root token 不能使用root token
		isroot, err := jwtoken.IsRootToken(token)
		if err != nil {
			common.GetLogger().Error(logFields, fmt.Sprintf("%s", err))
			b.Data["json"] = map[string]string{"result": "DEVOPS-API-TOKEN auth fail", "statuscode": "1"}
			b.ServeJSON()
			b.StopRun()
		}
		if isroot {
			common.GetLogger().Error(logFields, "can't use root token")
			b.Data["json"] = map[string]string{"result": "DEVOPS-API-TOKEN auth fail", "statuscode": "1"}
			b.ServeJSON()
			b.StopRun()
		}

		_, err = jwtoken.IsTokenValid(token)
		if err != nil {
			common.GetLogger().Error(logFields, fmt.Sprintf("%s", err))
			b.Data["json"] = map[string]string{"result": "DEVOPS-API-TOKEN auth fail", "statuscode": "1"}
			b.ServeJSON()
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
