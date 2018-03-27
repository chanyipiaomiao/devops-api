package controllers

import (
	"devops-api/common"
	"fmt"
)

// Post 接收中国的节假日安排, 为判断节假日和工作日准备
func (h *HolidayController) Post() {
	holiday := &common.HoliWorkday{}
	holiday.Setting(h.Ctx.Input.RequestBody)
	h.Data["json"] = map[string]interface{}{
		"setting": "Ok",
	}
	h.ServeJSON()
}

// Get 接收一个日期，判断是节假日还是工作日
func (h *HolidayController) Get() {
	date := h.GetString("date")
	holiworkday := &common.HoliWorkday{}
	r, err := holiworkday.IsHoliWorkday(date)
	if err != nil {
		fmt.Println(err)
	}
	h.Data["json"] = map[string]interface{}{
		"date": date,
		"type": r,
	}
	h.ServeJSON()
}
