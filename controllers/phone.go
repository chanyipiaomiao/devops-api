package controllers

import (
	"devops-api/common"
	"fmt"
)

var (
	queryPhoneEntryType = "Query Phone Location"
)

func (p *PhoneController) Get() {
	phone := p.GetString("phone")
	m, err := common.QueryPhone(phone)
	if err != nil {
		p.JsonError(queryPhoneEntryType, fmt.Sprintf("%s", err), StringMap{}, true)
		return
	}
	p.JsonOK(queryPhoneEntryType, m, true)
}
