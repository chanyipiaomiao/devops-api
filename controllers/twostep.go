package controllers

import (
	"devops-api/common"
	"log"
)

// Enable 启用2步验证
func (t *TwoStepAuthController) Enable() {

	username := t.GetString("username")
	issuer := t.GetString("issuer")

	two := common.NewTwoStepAuth(username, issuer)
	m, err := two.Enable()
	if err != nil {
		log.Println(err)
	}
	t.Data["json"] = m
	t.ServeJSON()
}

// Auth 验证用户输入的6位数字
func (t *TwoStepAuthController) Auth() {

	username := t.GetString("username")
	issuer := t.GetString("issuer")
	token := t.GetString("token")

	two := common.NewTwoStepAuth(username, issuer)
	ok, err := two.Auth(token)
	if err != nil {
		panic(err)
	}
	t.Data["json"] = map[string]interface{}{
		"auth": ok,
	}
	t.ServeJSON()
}
