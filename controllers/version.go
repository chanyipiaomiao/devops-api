package controllers

import (
	"devops-api/common"
	"fmt"
)

// Get 获取程序版本号
func (v *VersionController) Get() {
	v.Data["json"] = common.GetVersion()
	versionFields := map[string]interface{}{
		"entryType": "Get App Version",
		"requestID": v.Data["RequestID"].(string),
	}
	common.GetLogger().Info(versionFields, fmt.Sprintf("IP(%s)获取程序版本号.", v.Data["RemoteIP"]))
	v.ServeJSON()
}
