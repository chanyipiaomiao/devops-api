package common

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/chanyipiaomiao/cal"
	"github.com/chanyipiaomiao/hltool"
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

// Setting 保存节假日和工作日设置
func (h *HoliWorkday) Setting(reqBody []byte) error {

	r := new(ReqHoliday)
	err := json.Unmarshal(reqBody, r)
	if err != nil {
		return err
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

	db, err := hltool.NewBoltDB(DBPath, holiworkdayTableName)
	if err != nil {
		return err
	}

	o, err := hltool.StructToBytes(calendar)
	if err != nil {
		return err
	}
	err = db.Set(map[string][]byte{
		r.Year: o,
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
		fmt.Println("1111")
		return "", err
	}
	db, err := hltool.NewBoltDB(DBPath, holiworkdayTableName)
	if err != nil {
		fmt.Println("2222")
		return "", err
	}
	year := string(t.Year())
	r, err := db.Get([]string{year})
	if err != nil {
		fmt.Println("3333")
		return "", err
	}
	calendar := &cal.Calendar{}
	err = hltool.BytesToStruct(r[year], calendar)
	if err != nil {
		fmt.Println("4444")
		return "", err
	}
	if calendar.IsWorkday(t) {
		return "workday", nil
	}
	return "holiday", nil
}
