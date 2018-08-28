package controllers

import (
	"fmt"
	"path"
	"strings"

	"devops-api/common"
)

var (
	mailEntryType = "SendMail"
)

// SendMail 发送邮件
func (e *EmailController) SendMail() {
	subject := e.GetString("subject")
	content := e.GetString("content")
	contentType := e.GetString("type")
	to := e.GetString("to")
	cc := e.GetString("cc")

	isattach, err := e.GetBool("isattach")
	if err != nil {
		isattach = false
	}

	var attachFilename string
	if isattach {
		f, h, err := e.GetFile("attach")
		if err != nil {
			errs := fmt.Sprintf("获取文件失败: %s", err)
			e.LogError(mailEntryType, StringMap{"errmsg": errs})
			e.JsonError(mailEntryType, errs, StringMap{}, false)
			return
		}
		defer f.Close()
		attachFilename = path.Join(common.UploadPath, h.Filename)
		e.SaveToFile("attach", attachFilename)
	}

	if subject == "" || content == "" {
		e.JsonError(mailEntryType, "发送失败: 主题或者内容不能为空", StringMap{"result": "send fail"}, false)
		return
	}

	if to == "" {
		e.JsonError(mailEntryType, "发送失败: 收件人不能为空", StringMap{"result": "send fail"}, false)
		return
	}

	if contentType == "" {
		contentType = "text/plain"
	}

	toMail := strings.Split(to, ",")

	var ccMail []string
	if cc == "" {
		ccMail = []string{}
	} else {
		ccMail = strings.Split(cc, ",")
	}

	_, err = common.SendByEmail(subject, content, contentType, attachFilename, toMail, ccMail)
	if err == nil {
		e.JsonOK(mailEntryType, StringMap{"result": "send ok"}, true)
		return
	}
	e.JsonError(mailEntryType, fmt.Sprintf("error: %s", err), StringMap{"result": "send fail"}, false)
}
