package controllers

import (
	"devops-api/common"
	"fmt"
	"strings"
)

var (
	StorePasswordEntryType = "Store Password"
)

// Update 密码管理 保存/更新密码
func (m *StorePasswordController) Post() {
	man, err := common.NewManagePassword()
	if err != nil {
		m.JsonError(StorePasswordEntryType, fmt.Sprintf("error: %s", err), StringMap{})
		return
	}
	err = man.Update(m.Ctx.Input.RequestBody)
	if err != nil {
		m.JsonError(StorePasswordEntryType, fmt.Sprintf("error: %s", err), StringMap{})
		return
	}
	m.JsonOK(StorePasswordEntryType, StringMap{"update": "ok"})
}

// Get 获取密码
func (m *StorePasswordController) Get() {
	man, err := common.NewManagePassword()
	if err != nil {
		m.LogError(StorePasswordEntryType, StringMap{"errmsg": fmt.Sprintf("error: %s", err)})
		return
	}
	//ids := m.GetString("id")
	ids := m.Ctx.Input.Param(":id")
	if ids == "" {
		m.JsonError(StorePasswordEntryType, "标识不能为空", StringMap{})
		return
	}
	result, err := man.Get(strings.Split(ids, ","))
	if err != nil {
		m.JsonError(StorePasswordEntryType, fmt.Sprintf("error: %s", err), StringMap{})
		return
	}
	m.JsonOK(StorePasswordEntryType, StringMap{"password": result})
}

// Delete 密码管理 删除密码
func (m *StorePasswordController) Delete() {
	man, err := common.NewManagePassword()
	if err != nil {
		m.JsonError(StorePasswordEntryType, fmt.Sprintf("error: %s", err), StringMap{})
		return
	}
	//ids := m.GetString("id")
	ids := m.Ctx.Input.Param(":id")
	if ids == "" {
		m.JsonError(StorePasswordEntryType, "标识不能为空", StringMap{})
		return
	}
	err = man.Delete(strings.Split(ids, ","))
	if err != nil {
		m.JsonError(StorePasswordEntryType, fmt.Sprintf("error: %s", err), StringMap{})
		return
	}
	m.JsonOK(StorePasswordEntryType, StringMap{"delete": "ok", "id": ids})
}
