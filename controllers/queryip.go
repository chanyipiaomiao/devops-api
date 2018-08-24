package controllers

import (
	"devops-api/common"
	"fmt"
)

// Get Get方法
func (q *QueryIPController) Get() {

	ip := q.Ctx.Input.Param(":ip")
	qip := common.NewQueryIP("data/ip2region.db")
	r, err := qip.Query(ip)
	if err != nil {
		q.JsonError("Query IP", fmt.Sprintf("%s", err), StringMap{})
		return
	}

	data := map[string]interface{}{
		"ip":     ip,
		"ipInfo": r,
	}
	q.JsonOK("query ip", data)
}
