package controllers

import (
	"fmt"
	"strings"

	"github.com/chanyipiaomiao/hltool"

	"devops-api/common"
	"github.com/astaxie/beego"
)

var (
	genPasswordEntryType = "GenPassword"
)

// GenPassword 生成指定长度的密码
func (p *PasswordController) GenPassword() {
	length, err := p.GetInt("length")
	if err != nil {
		length = 32
	}
	name := p.GetString("name")
	special := strings.ToLower(p.GetString("specialChar"))

	if name == "" {
		p.JsonOK(genPasswordEntryType,
			StringMap{"password": hltool.GenRandomString(length, special), "length": length}, true)
	} else {
		m := make(map[string]string)
		names := strings.Split(name, ",")
		for _, name := range names {
			m[name] = hltool.GenRandomString(length, special)
		}
		p.JsonOK(genPasswordEntryType, StringMap{"password": m}, true)
	}
}

// ManualGenAuthPassword 手动生成验证密码
func (p *PasswordController) ManualGenAuthPassword() {
	ok, err := beego.AppConfig.Bool("authpassword::enableManualGenAuthPassword")
	if err != nil {
		p.JsonError(genPasswordEntryType, fmt.Sprintf("%s", err), StringMap{"manualGenAuthPassword": false}, true)
		return
	}
	if !ok {
		p.LogError(genPasswordEntryType, StringMap{"errmsg": "app.conf enableManualGenAuthPassword not true"})
		p.JsonError(genPasswordEntryType, "ManualGenAuthPassword error", StringMap{"manualGenAuthPassword": false}, false)
		return
	}

	m := common.ManualGenAuthPassword()
	if result, _ := m["result"]; result {
		p.JsonOK(genPasswordEntryType, StringMap{"manualGenAuthPassword": true}, true)
		return
	}

}

// AuthGenPassword 验证生成上面生成的密码
func (p *PasswordController) AuthGenPassword() {
	password := p.GetString("password")
	if "" == password {
		p.LogError(genPasswordEntryType, StringMap{"errmsg": "使用空密码进行验证, 验证失败."})
		p.JsonError(genPasswordEntryType, "auth fail", StringMap{"auth": false}, false)
		return
	}
	willAuthPassword := common.GetWiillAuthPassword()
	if willAuthPassword == password {
		p.LogInfo(genPasswordEntryType, StringMap{"result": "验证成功."})
		p.JsonOK(genPasswordEntryType, StringMap{"auth": true}, false)
	} else {
		p.LogError(genPasswordEntryType,
			StringMap{"errmsg": fmt.Sprintf("验证失败.要验证的密码是: (%s) 正确的密码是: (%s)", password, willAuthPassword)})
		p.JsonError(genPasswordEntryType, "auth fail", StringMap{"auth": false}, false)
	}
}
