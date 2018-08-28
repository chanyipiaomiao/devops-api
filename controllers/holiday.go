package controllers

import (
	"fmt"

	"devops-api/common"
)

// Post 接收中国的节假日安排, 为判断节假日和工作日准备
func (h *HolidayController) Post() {
	entryType := "Setting Holiday/Workday"
	holiday := &common.HoliWorkday{}
	err := holiday.Setting(h.Ctx.Input.RequestBody)
	if err != nil {
		h.JsonError(entryType, fmt.Sprintf("%s", err), StringMap{}, true)
		return
	}
	h.JsonOK(entryType, StringMap{"setting": "ok"}, true)
}

// Get 接收一个日期，判断是节假日还是工作日
func (h *HolidayController) Get() {
	entryType := "Get Holiday/Workday"
	date := h.GetString("date")
	holiworkday := &common.HoliWorkday{}
	r, err := holiworkday.IsHoliWorkday(date)
	if err != nil {
		h.JsonError(entryType, fmt.Sprintf("%s", err), StringMap{}, true)
		return
	}
	h.JsonOK(entryType, StringMap{"date": date, "dateType": r}, true)
}
