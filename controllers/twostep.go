package controllers

import (
	"fmt"

	"devops-api/common"
)

var (
	twoStepEntryType = "TwoStepAuth"
)

// Enable 启用2步验证
func (t *TwoStepAuthController) Enable() {

	username := t.GetString("username")
	issuer := t.GetString("issuer")

	two := common.NewTwoStepAuth(username)
	two.Issuer = issuer
	two.Digits = common.TwoStepAuthDigits

	m, err := two.Enable()
	if err != nil {
		t.JsonError(twoStepEntryType,
			fmt.Sprintf("%s", err), StringMap{"enable": "no", "username": username}, true)
		return
	}
	t.JsonOK(twoStepEntryType, m, true)
}

// Disable 禁用2步验证
func (t *TwoStepAuthController) Disable() {
	username := t.GetString("username")
	two := common.NewTwoStepAuth(username)
	err := two.Disable()

	if err != nil {
		t.JsonError(twoStepEntryType,
			fmt.Sprintf("%s", err), StringMap{"disable": "no", "username": username}, true)
		return
	}
	t.JsonOK(twoStepEntryType, StringMap{"disable": "yes", "username": username}, true)
}

// Auth 验证用户输入的6位数字
func (t *TwoStepAuthController) Auth() {

	username := t.GetString("username")
	issuer := t.GetString("issuer")
	token := t.GetString("token")

	two := common.NewTwoStepAuth(username)
	two.Issuer = issuer
	isok, err := two.Auth(token)

	if err == nil {
		t.JsonOK(twoStepEntryType, StringMap{"username": username, "issuer": issuer, "auth": isok}, true)
		return
	}
	t.JsonError(twoStepEntryType, fmt.Sprintf("%s", err),
		StringMap{"username": username, "issuer": issuer, "auth": isok}, true)
}
