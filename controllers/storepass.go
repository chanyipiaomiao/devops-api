package controllers

import (
	"devops-api/common"
	"fmt"
	"strings"
)

var (
	storePasswordCommonFields = map[string]interface{}{
		"entryType": "Store Password",
	}
)

// Update 密码管理 保存/更新密码
func (m *StorePasswordController) Update() {
	reqID := m.Data["RequestID"].(string)
	storePasswordCommonFields["requestID"] = reqID
	man, err := common.NewManagePassword()
	logger := common.GetLogger()
	if err != nil {
		logger.Error(storePasswordCommonFields, fmt.Sprintf("error: %s", err))
		return
	}
	err = man.Update(m.Ctx.Input.RequestBody)
	if err != nil {
		logger.Error(storePasswordCommonFields, fmt.Sprintf("error: %s", err))
		return
	}
	logger.Info(storePasswordCommonFields, fmt.Sprintf("IP: %s, 更新密码成功.", m.Data["RemoteIP"]))
	m.Data["json"] = map[string]interface{}{"update": true, "requestID": reqID, "statuscode": 0}
	m.ServeJSON()
}

// Get 获取密码
func (m *StorePasswordController) Get() {
	reqID := m.Data["RequestID"].(string)
	storePasswordCommonFields["requestID"] = reqID
	man, err := common.NewManagePassword()
	logger := common.GetLogger()
	if err != nil {
		logger.Error(storePasswordCommonFields, fmt.Sprintf("error: %s", err))
		return
	}
	ids := m.GetString("id")
	if ids == "" {
		m.Data["json"] = map[string]interface{}{"error": "标识不能为空", "requestID": reqID, "statuscode": 1}
		m.ServeJSON()
		return
	}
	result, err := man.Get(strings.Split(ids, ","))
	if err != nil {
		logger.Error(storePasswordCommonFields, fmt.Sprintf("error: %s", err))
		m.Data["json"] = map[string]interface{}{"error": fmt.Sprintf("error: %s", err), "requestID": reqID, "statuscode": 1}
		m.ServeJSON()
		return
	}
	logger.Info(storePasswordCommonFields, fmt.Sprintf("IP: %s, 获取标识: %s 的密码成功", m.Data["RemoteIP"], ids))
	m.Data["json"] = map[string]interface{}{"get": true, "requestID": reqID, "statuscode": 0, "password": result}
	m.ServeJSON()
}

// Delete 密码管理 删除密码
func (m *StorePasswordController) Delete() {
	reqID := m.Data["RequestID"].(string)
	storePasswordCommonFields["requestID"] = reqID
	man, err := common.NewManagePassword()
	logger := common.GetLogger()
	if err != nil {
		logger.Error(storePasswordCommonFields, fmt.Sprintf("%s", err))
		m.Data["json"] = map[string]interface{}{"error": fmt.Sprintf("%s", err), "requestID": reqID, "statuscode": 1}
		m.ServeJSON()
		return
	}
	ids := m.GetString("id")
	if ids == "" {
		m.Data["json"] = map[string]interface{}{"error": "标识不能为空", "requestID": reqID, "statuscode": 1}
		m.ServeJSON()
		return
	}
	err = man.Delete(strings.Split(ids, ","))
	if err != nil {
		logger.Error(storePasswordCommonFields, fmt.Sprintf("error: %s", err))
		m.Data["json"] = map[string]interface{}{"error": fmt.Sprintf("error: %s", err), "requestID": reqID, "statuscode": 1}
		m.ServeJSON()
		return
	}
	logger.Info(storePasswordCommonFields, fmt.Sprintf("IP: %s, 删除密码成功: %s", m.Data["RemoteIP"], ids))
	m.Data["json"] = map[string]interface{}{"delete": true, "requestID": reqID, "statuscode": 0, "id": ids}
	m.ServeJSON()

}
