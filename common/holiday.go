package common

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/chanyipiaomiao/cal"
	"github.com/chanyipiaomiao/hltool"
	"github.com/tidwall/gjson"
)

const (
	holiworkdayTableName = "holiworkday"
	dateTemplate         = "2006-01-02"
)

// ReqHoliday 请求过来的节假日和工作日设置
type ReqHoliday struct {
	Holiday []struct {
		EndTime   string `json:"end_time"`
		Name      string `json:"name"`
		StartTime string `json:"start_time"`
		ZhName    string `json:"zh_name"`
	} `json:"holiday"`
	Workday []string `json:"workday"`
	Year    string   `json:"year"`
}

// HoliWorkday 节假日和工作日
type HoliWorkday struct{}

// 解析工作日节假日json字符串
func (h *HoliWorkday) parse(jsonstr []byte) (*cal.Calendar, error) {

	r := new(ReqHoliday)
	err := json.Unmarshal(jsonstr, r)
	if err != nil {
		return nil, err
	}

	calendar := cal.NewCalendar()

	for _, v := range r.Holiday {
		endTime, _ := time.Parse(dateTemplate, v.EndTime)
		startTime, _ := time.Parse(dateTemplate, v.StartTime)
		calendar.AddHoliday(cal.NewHolidayExact(startTime.Month(), startTime.Day(), startTime.Year()))
		sub := endTime.Sub(startTime).Hours() / 24
		for i := 1; i < int(sub); i++ {
			t := startTime.AddDate(0, 0, i)
			calendar.AddHoliday(cal.NewHolidayExact(t.Month(), t.Day(), t.Year()))
		}
		calendar.AddHoliday(cal.NewHolidayExact(endTime.Month(), endTime.Day(), endTime.Year()))
	}

	for _, v := range r.Workday {
		t, _ := time.Parse(dateTemplate, v)
		calendar.AddExtraWorkday(t)
	}

	calendar.Observed = cal.ObservedExact

	return calendar, nil
}

// Setting 保存节假日和工作日设置
func (h *HoliWorkday) Setting(reqBody []byte) error {

	db, err := hltool.NewBoltDB(DBPath, holiworkdayTableName)
	if err != nil {
		return err
	}

	year := gjson.Get(string(reqBody), "year").String()
	err = db.Set(map[string][]byte{
		year: reqBody,
	})

	if err != nil {
		return err
	}

	return nil
}

// IsHoliWorkday 检查给定的日期是工作日还是节假日
func (h *HoliWorkday) IsHoliWorkday(date string) (string, error) {

	t, err := time.Parse(dateTemplate, date)
	if err != nil {
		return "", err
	}

	db, err := hltool.NewBoltDB(DBPath, holiworkdayTableName)
	if err != nil {
		return "", err
	}

	year := strconv.Itoa(t.Year())

	result, err := db.Get([]string{year})
	if err != nil {
		return "", err
	}
	if result[year] == nil {
		return "", fmt.Errorf("%s year holiday setting not in db, please setting", year)
	}

	calendar, err := h.parse(result[year])
	if err != nil {
		return "", err
	}

	if calendar.IsExtraWorkday(t) {
		return "workday", nil
	}

	if calendar.IsWorkday(t) {
		return "workday", nil
	}

	if calendar.IsHoliday(t) {
		return "holiday", nil
	}

	return "weekend", nil
}
