package controllers

import (
	"devops-api/common"
	"fmt"
)

// Get Get方法
func (q *QueryIPController) Get() {
	requestID := q.Data["RequestID"].(string)
	queryIPLog := map[string]interface{}{
		"entryType": "query ip",
		"requestID": requestID,
	}
	ip := q.GetString("ip")
	qip := common.NewQueryIP("data/ip2region.db")
	r, err := qip.Query(ip)
	if err != nil {
		queryIPLog["statuscode"] = 1
		queryIPLog["errmsg"] = fmt.Sprintf("%s", err)
		common.GetLogger().Error(queryIPLog, "查询IP地址区域")
		q.Data["json"] = queryIPLog
		q.ServeJSON()
		return
	}
	queryIPLog["statuscode"] = 0
	queryIPLog["errmsg"] = ""
	queryIPLog["ip"] = ip
	queryIPLog["ipInfo"] = r
	common.GetLogger().Info(queryIPLog, "查询IP地址区域")
	q.Data["json"] = queryIPLog
	q.ServeJSON()
}
