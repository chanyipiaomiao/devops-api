package controllers

import (
	"fmt"

	"devops-api/common"
)

// Enable 启用2步验证
func (t *TwoStepAuthController) Enable() {

	username := t.GetString("username")
	issuer := t.GetString("issuer")

	two := common.NewTwoStepAuth(username)
	two.Issuer = issuer
	two.Digits = common.TwoStepAuthDigits

	requestID := t.Data["RequestID"].(string)
	twoStepAuthLog := map[string]interface{}{
		"entryType": "TwoStepAuth",
		"username":  username,
		"issuer":    issuer,
		"requestId": requestID,
	}

	m, err := two.Enable()
	if err != nil {
		common.GetLogger().Error(twoStepAuthLog, fmt.Sprintf("%s", err))
		t.Data["json"] = map[string]interface{}{
			"requestId":  requestID,
			"statuscode": 1,
			"enable":     false,
		}
		t.ServeJSON()
		return
	}

	m["enable"] = true
	m["statuscode"] = 0
	m["requestId"] = requestID
	t.Data["json"] = m
	common.GetLogger().Info(twoStepAuthLog, "启用2步验证")
	t.ServeJSON()
}

// Disable 禁用2步验证
func (t *TwoStepAuthController) Disable() {
	username := t.GetString("username")
	two := common.NewTwoStepAuth(username)
	err := two.Disable()

	requestID := t.Data["RequestID"].(string)
	twoStepAuthLog := map[string]interface{}{
		"entryType": "TwoStepAuth",
		"username":  username,
		"requestId": requestID,
	}

	result := map[string]interface{}{
		"requestId": requestID,
		"username":  username,
	}
	t.Data["json"] = result
	if err != nil {
		result["statuscode"] = 1
		result["disable"] = false
		common.GetLogger().Error(twoStepAuthLog, fmt.Sprintf("%s", err))
		t.ServeJSON()
		return
	}

	result["statuscode"] = 0
	result["disable"] = true
	common.GetLogger().Info(twoStepAuthLog, "禁用2步验证")
	t.ServeJSON()
}

// Auth 验证用户输入的6位数字
func (t *TwoStepAuthController) Auth() {

	username := t.GetString("username")
	issuer := t.GetString("issuer")
	token := t.GetString("token")

	two := common.NewTwoStepAuth(username)
	two.Issuer = issuer
	isok, err := two.Auth(token)

	requestID := t.Data["RequestID"].(string)
	twoStepAuthLog := map[string]interface{}{
		"entryType": "TwoStepAuth",
		"username":  username,
		"issuer":    issuer,
		"token":     token,
		"requestId": requestID,
	}

	result := map[string]interface{}{
		"requestId": requestID,
		"auth":      isok,
		"username":  username,
	}
	t.Data["json"] = result

	if err != nil {
		result["statuscode"] = 1
		common.GetLogger().Error(twoStepAuthLog, fmt.Sprintf("%s", err))
		t.ServeJSON()
		return
	}
	result["statuscode"] = 0
	common.GetLogger().Info(twoStepAuthLog, "验证2步验证")
	t.ServeJSON()
}
