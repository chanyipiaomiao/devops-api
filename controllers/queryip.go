package controllers

import (
	"devops-api/common"
	"fmt"
)

// Get Get方法
func (q *QueryIPController) Get() {

	ip := q.GetString("ip")
	qip := common.NewQueryIP("data/ip2region.db")
	r, err := qip.Query(ip)
	if err != nil {
		q.JsonError("query ip", fmt.Sprintf("%s", err), NullStringMap{})
		return
	}

	data := map[string]interface{}{
		"ip":     ip,
		"ipInfo": r,
	}
	q.JsonOK("query ip", data)
}
