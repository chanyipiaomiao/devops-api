package common

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/hltool"
)

const (

	// 存放密码表
	manPassTable = "password"
)

// JSONManPassFields 通过POST传递过来的json字符串结构
type JSONManPassFields struct {
	Password []struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	} `json:"password"`
	UniqueID string `json:"uniqueId"`
}

// ManagePassword 密码管理
type ManagePassword struct {
	Jwt *hltool.JWToken
	DB  *hltool.BoltDB
}

// NewManagePassword 返回对象
func NewManagePassword() (*ManagePassword, error) {
	passwordDb, err := hltool.NewBoltDB(DBPath, manPassTable)
	if err != nil {
		return nil, err
	}
	crytoString := beego.AppConfig.String("security::jwtokenSignString")
	if crytoString == "" {
		return nil, fmt.Errorf("warning: in conf file jwtokenSignString must not null")
	}
	return &ManagePassword{
		Jwt: hltool.NewJWToken(crytoString),
		DB:  passwordDb,
	}, nil
}

// save 根据k,v存储
func (m *ManagePassword) save(key string, v map[string]interface{}) error {

	// 密码存储使用 jwt 的方式进行加密存储，方便解密
	crytoAfterString, err := m.Jwt.GenJWToken(v)
	if err != nil {
		return err
	}

	// 存入到数据库中
	return m.DB.Set(map[string][]byte{
		key: []byte(crytoAfterString),
	})
}

// Update 保存/更新密码
func (m *ManagePassword) Update(reqbody []byte) error {

	// 把请求体转换为结构体
	passwordFields := new(JSONManPassFields)
	err := json.Unmarshal(reqbody, passwordFields)
	if err != nil {
		return err
	}

	// 加密前的密码map，用来使用Jwt加密存储
	inReqPassword := make(map[string]interface{})
	for _, v := range passwordFields.Password {
		inReqPassword[v.Name] = v.Password
	}

	// 先根据唯一标识从数据库获取一下，看看其值是否存在
	result, err := m.get([]string{passwordFields.UniqueID})
	if err != nil {
		return err
	}

	// 如果该标识在数据库中不存在,直接保存
	if len(result) == 0 {
		return m.save(passwordFields.UniqueID, inReqPassword)
	}

	// 如果存在,则进行请求体中的和数据库中的比对进行更新
	inDBpassword, err := m.Jwt.ParseJWToken(string(result[passwordFields.UniqueID]))
	if err != nil {
		return err
	}
	for k, v := range inReqPassword {
		if inDBpassword[k].(string) != v {
			inDBpassword[k] = v
		}
	}

	return m.save(passwordFields.UniqueID, inDBpassword)
}

// get 数据库中根据id查询
func (m *ManagePassword) get(uniqueIDs []string) (map[string][]byte, error) {
	r, err := m.DB.Get(uniqueIDs)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Get 根据唯一标识获取值,如果数组中所有的标识在数据库都没有值,就返回一个空的map
func (m *ManagePassword) Get(uniqueIDs []string) (map[string]interface{}, error) {
	result, err := m.get(uniqueIDs)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return map[string]interface{}{}, nil
	}
	r := make(map[string]interface{})
	for _, id := range uniqueIDs {
		if _, ok := result[id]; !ok {
			r[id] = map[string]interface{}{}
			continue
		}
		inDBpassword, err := m.Jwt.ParseJWToken(string(result[id]))
		if err != nil {
			continue
		}
		r[id] = inDBpassword
	}
	return r, nil
}

// Delete 根据唯一标识 删除
func (m *ManagePassword) Delete(uniqueIDs []string) error {
	return m.DB.Delete(uniqueIDs)
}
