package common

import (
	"crypto"
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
)

// TwoStepAuth 2步验证
type TwoStepAuth struct {

	// 用户名/账号/标识
	Username string

	// 发行者
	Issuer string
}

// NewTwoStepAuth 返回TwoStepAuth对象
func NewTwoStepAuth(username, issuer string) *TwoStepAuth {
	return &TwoStepAuth{
		Username: username,
		Issuer:   issuer,
	}
}

// Enable 启用2步验证
// return 生成的二维码图片路径 和 KEY，可以手动添加KEY，如果不支持扫描
func (t *TwoStepAuth) Enable() (map[string]string, error) {
	otp, err := twofactor.NewTOTP(t.Username, t.Issuer, crypto.SHA1, 8)
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

	return map[string]string{"key": otp.Secret(), "qrImage": path.Join("/", imgPath)}, nil
}

// Auth 验证用户输入的6位数字
func (t *TwoStepAuth) Auth(userCode string) (bool, error) {
	otp, err := t.GetOtp()
	if err != nil {
		return false, err
	}
	err = otp.Validate(userCode)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetOtp 从数据库中取出 otp对象
func (t *TwoStepAuth) GetOtp() (*twofactor.Totp, error) {
	var err error
	twoStepDb, err = hltool.NewBoltDB(DBPath, twoStepTable)
	if err != nil {
		return nil, err
	}
	m, err := twoStepDb.Get([]string{t.Username})
	if err != nil {
		return nil, err
	}
	otp, err := twofactor.TOTPFromBytes([]byte(m[t.Username]), t.Issuer)
	if err != nil {
		return nil, err
	}
	return otp, nil
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
