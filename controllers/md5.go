package controllers

import (
	"github.com/chanyipiaomiao/hltool"
)

// Get 方法
func (m *MD5Controller) Get() {
	rawString := m.GetString("rawstr")
	rawStringMD5 := hltool.GetMD5(rawString)
	data := map[string]string{
		"rawString":    rawString,
		"rawStringMD5": rawStringMD5,
	}
	m.JsonOK("Get String MD5", data, true)
}
