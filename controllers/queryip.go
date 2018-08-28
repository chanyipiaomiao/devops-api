package controllers

import (
	"devops-api/common"
	"fmt"
)

var (
	queryIPEntryType = "Query IP"
)

// Get Get方法
func (q *QueryIPController) Get() {

	ip := q.GetString("ip")
	qip := common.NewQueryIP("data/ip2region.db")
	r, err := qip.Query(ip)
	if err != nil {
		q.JsonError(queryIPEntryType, fmt.Sprintf("%s", err), StringMap{}, true)
		return
	}

	data := map[string]interface{}{
		"ip":     ip,
		"ipInfo": r,
	}
	q.JsonOK(queryIPEntryType, data, true)
}
