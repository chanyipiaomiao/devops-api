package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/hltool"

	"devops-api/common"
)

var (
	passwordFields = map[string]interface{}{
		"entryType": "GenPassword",
	}
)

// GenPassword 生成指定长度的密码
func (p *PasswordController) GenPassword() {
	length, err := p.GetInt("length")
	if err != nil {
		length = 32
	}
	name := p.GetString("name")
	special := strings.ToLower(p.GetString("specialChar"))
	passwordFields["requestID"] = p.Data["RequestID"].(string)

	if name == "" {
		p.Data["json"] = map[string]string{"password": hltool.GenRandomString(length, special), "requestID": p.Data["RequestID"].(string)}
		common.GetLogger().Info(passwordFields, fmt.Sprintf("IP(%s)请求生成密码,长度(%d)", p.Data["RemoteIP"], length))
		p.ServeJSON()
	} else {
		m := make(map[string]string)
		names := strings.Split(name, ",")
		for _, name := range names {
			m[name] = hltool.GenRandomString(length, special)
		}
		common.GetLogger().Info(passwordFields, fmt.Sprintf("IP(%s)请求生成密码,长度(%d),名称(%s)", p.Data["RemoteIP"], length, name))
		p.Data["json"] = map[string]interface{}{
			"requestID": p.Data["RequestID"].(string),
			"password":  m,
		}
		p.ServeJSON()
	}
}

// AuthGenPassword 验证生成的密码
func (p *PasswordController) AuthGenPassword() {
	password := p.GetString("password")
	passwordFields["requestID"] = p.Data["RequestID"].(string)
	if "" == password {
		p.Data["json"] = map[string]interface{}{"auth": false, "requestID": p.Data["RequestID"].(string)}
		common.GetLogger().Info(passwordFields, fmt.Sprintf("IP(%s)请求验证密码,但是使用空密码进行, 验证失败.", p.Data["RemoteIP"]))
		p.ServeJSON()
		return
	}
	willAuthPassword := common.GetWiillAuthPassword()
	if willAuthPassword == password {
		p.Data["json"] = map[string]interface{}{"auth": true, "requestID": p.Data["RequestID"].(string)}
		common.GetLogger().Info(passwordFields, fmt.Sprintf("IP(%s)请求验证密码,验证成功.", p.Data["RemoteIP"]))
		p.ServeJSON()
	} else {
		p.Data["json"] = map[string]interface{}{"auth": false, "requestID": p.Data["RequestID"].(string)}
		common.GetLogger().Info(passwordFields, fmt.Sprintf("IP(%s)请求验证密码,验证失败.要验证的密码是: (%s) 正确的密码是: (%s)", p.Data["RemoteIP"], password, willAuthPassword))
		p.ServeJSON()
	}
}

// ManualGenAuthPassword 手动生成验证密码
func (p *PasswordController) ManualGenAuthPassword() {
	passwordFields["requestID"] = p.Data["RequestID"].(string)
	ok, err := beego.AppConfig.Bool("authpassword::enableManualGenAuthPassword")
	if err != nil {
		p.Data["json"] = map[string]interface{}{"manualGenAuthPassword": false, "requestID": p.Data["RequestID"].(string)}
		common.GetLogger().Error(passwordFields, fmt.Sprintf("%s", err))
		p.ServeJSON()
		return
	}
	if !ok {
		p.Data["json"] = map[string]string{"error": "app.conf enableManualGenAuthPassword not true", "requestID": p.Data["RequestID"].(string)}
		p.ServeJSON()
		return
	}

	m := common.ManualGenAuthPassword()
	if result, _ := m["result"]; result {
		p.Data["json"] = map[string]interface{}{"manualGenAuthPassword": true, "requestID": p.Data["RequestID"].(string)}
		p.ServeJSON()
		return
	}

}
