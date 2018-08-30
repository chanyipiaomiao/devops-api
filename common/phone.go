package common

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/phonedata"
)

func QueryPhone(phone string) (map[string]string, error) {
	dbPath := beego.AppConfig.String("phone::dbPath")
	if dbPath == "" {
		return nil, fmt.Errorf("not found phone dat file")
	}
	p, err := phonedata.NewPhoneQuery(dbPath)
	if err != nil {
		return nil, err
	}
	m, err := p.Query(phone)
	if err != nil {
		return nil, err
	}
	return m, nil
}
