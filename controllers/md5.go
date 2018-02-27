package controllers

import (
	"devops-api/common"

	"github.com/chanyipiaomiao/hltool"
)

// Get 方法
func (m *MD5Controller) Get() {
	rawString := m.GetString("string")
	rawStringMD5 := hltool.GetMD5(rawString)
	m.Data["json"] = map[string]string{"rawString": rawString, "md5String": rawStringMD5, "requestID": m.Data["RequestID"].(string)}
	logFields := map[string]interface{}{
		"entryType":    "Get String MD5",
		"rawString":    rawString,
		"rawStringMD5": rawStringMD5,
		"requestID":    m.Data["RequestID"],
	}
	common.GetLogger().Info(logFields, "获取MD5值")
	m.ServeJSON()
}
