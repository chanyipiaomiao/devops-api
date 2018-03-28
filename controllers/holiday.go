package controllers

import (
	"fmt"

	"devops-api/common"
)

// Post 接收中国的节假日安排, 为判断节假日和工作日准备
func (h *HolidayController) Post() {
	requestID := h.Data["RequestID"].(string)
	holidayLog := map[string]interface{}{
		"entryType": "setting holiday/workday",
		"requestID": requestID,
	}
	holiday := &common.HoliWorkday{}
	err := holiday.Setting(h.Ctx.Input.RequestBody)
	if err != nil {
		holidayLog["statuscode"] = 1
		holidayLog["errmsg"] = fmt.Sprintf("%s", err)
		holidayLog["result"] = "error"
		common.GetLogger().Error(holidayLog, "设置节假日和工作日")
		h.Data["json"] = holidayLog
		h.ServeJSON()
		return
	}

	holidayLog["statuscode"] = 0
	holidayLog["errmsg"] = ""
	holidayLog["result"] = "ok"
	common.GetLogger().Info(holidayLog, "设置节假日和工作日")
	h.Data["json"] = holidayLog
	h.ServeJSON()
}

// Get 接收一个日期，判断是节假日还是工作日
func (h *HolidayController) Get() {
	requestID := h.Data["RequestID"].(string)
	holidayLog := map[string]interface{}{
		"entryType": "judgment holiday/workday",
		"requestID": requestID,
	}
	date := h.GetString("date")
	holiworkday := &common.HoliWorkday{}
	r, err := holiworkday.IsHoliWorkday(date)
	if err != nil {
		holidayLog["statuscode"] = 1
		holidayLog["errmsg"] = fmt.Sprintf("%s", err)
		common.GetLogger().Error(holidayLog, "判断节假日和工作日")
		h.Data["json"] = holidayLog
		h.ServeJSON()
		return
	}

	holidayLog["statuscode"] = 0
	holidayLog["errmsg"] = ""
	holidayLog["date"] = date
	holidayLog["dateType"] = r
	common.GetLogger().Info(holidayLog, "判断节假日和工作日")
	h.Data["json"] = holidayLog
	h.ServeJSON()
}
