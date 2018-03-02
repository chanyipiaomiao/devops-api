package common

import (
	"crypto"
	"os"
	"path"

	"github.com/chanyipiaomiao/hltool"
	"github.com/sec51/twofactor"
)

var (

	// twoStepDb db操作对象
	twoStepDb *hltool.BoltDB
)

const (
	twoStepTable      = "twostep"
	twoStepTableCount = "twostep_count"

	// TwoStepAuthDigits 验证码的位数
	TwoStepAuthDigits = 6
)

// TwoStepAuth 2步验证
type TwoStepAuth struct {

	// 用户名/账号/标识
	Username string

	// 发行者
	Issuer string

	// 身份验证器上显示的位数 6|7|8 一般是6位
	Digits int
}

// NewTwoStepAuth 返回TwoStepAuth对象
func NewTwoStepAuth(username string) *TwoStepAuth {
	return &TwoStepAuth{
		Username: username,
	}
}

// SaveOtp 保存 2步验证的 对象到数据库
func (t *TwoStepAuth) SaveOtp(otp *twofactor.Totp) error {
	var err error
	twoStepDb, err = hltool.NewBoltDB(DBPath, twoStepTable)
	if err != nil {
		return err
	}
	otpBytes, err := otp.ToBytes()
	if err != nil {
		return err
	}

	twoStepDb.Set(map[string][]byte{
		t.Username: otpBytes,
	})

	return nil
}

// get 根据用户名(键)获取 从表中获取 值
func (t *TwoStepAuth) get() (map[string][]byte, error) {
	var err error
	twoStepDb, err = hltool.NewBoltDB(DBPath, twoStepTable)
	if err != nil {
		return nil, err
	}
	m, err := twoStepDb.Get([]string{t.Username})
	if err != nil {
		return nil, err
	}
	return m, nil
}

// GetOtp 从数据库中取出 otp对象
func (t *TwoStepAuth) GetOtp() (*twofactor.Totp, error) {
	m, err := t.get()
	if err != nil {
		return nil, err
	}
	if m[t.Username] != nil {
		otp, err := twofactor.TOTPFromBytes(m[t.Username], t.Issuer)
		if err != nil {
			return nil, err
		}
		return otp, nil
	}
	return nil, nil
}

// Enable 启用2步验证
// return 生成的二维码图片路径 和 KEY，可以手动添加KEY，如果不支持扫描
func (t *TwoStepAuth) Enable() (map[string]interface{}, error) {
	otp, err := twofactor.NewTOTP(t.Username, t.Issuer, crypto.SHA1, TwoStepAuthDigits)
	if err != nil {
		return nil, err
	}
	qrBytes, err := otp.QR()
	if err != nil {
		return nil, err
	}
	imgPath := path.Join(QrImageDir, t.Username+".png")
	err = hltool.BytesToImage(qrBytes, imgPath)
	if err != nil {
		return nil, err
	}

	// 保存otp对象到数据库，到验证的时候取出来再验证
	t.SaveOtp(otp)

	return map[string]interface{}{"key": otp.Secret(), "qrImage": path.Join("/", imgPath)}, nil
}

// Disable 禁用2步验证 实际就是从数据库删除记录
func (t *TwoStepAuth) Disable() error {
	twoStepDb, err := hltool.NewBoltDB(DBPath, twoStepTable)
	if err != nil {
		return err
	}

	// 从数据库中删除该用户名
	err = twoStepDb.Delete([]string{t.Username})
	if err != nil {
		return err
	}

	// 从磁盘上删除生成的用户对应的二维码图片文件
	imgPath := path.Join(QrImageDir, t.Username+".png")
	if hltool.IsExist(imgPath) {
		err = os.Remove(imgPath)
		if err != nil {
			return err
		}
	}

	return nil
}

// Auth 验证用户输入的6位数字
func (t *TwoStepAuth) Auth(userCode string) (bool, error) {
	otp, err := t.GetOtp()
	if err != nil {
		return false, err
	}

	if otp == nil {
		return false, nil
	}

	err = otp.Validate(userCode)
	if err != nil {
		return false, err
	}

	return true, nil
}
