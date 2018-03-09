package controllers

import (
	"devops-api/common"
)

// Update 密码管理 保存/更新密码
func (m *ManagePasswordController) Update() {
	man, err := common.NewManagePassword()
	if err != nil {
		return
	}
	err = man.Update(m.Ctx.Input.RequestBody)
	if err != nil {
		return
	}
}

// Delete 密码管理 删除密码
func (m *ManagePasswordController) Delete() {

}
