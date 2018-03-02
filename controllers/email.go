package controllers

import (
	"fmt"
	"path"
	"strings"

	"devops-api/common"
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
			errs := fmt.Sprintf("从POST中获取文件失败: %s", err)
			getAttachFields := map[string]interface{}{
				"entryType": "Get Attach From Post",
				"requestID": e.Data["RequestID"].(string),
			}
			common.GetLogger().Error(getAttachFields, errs)
			e.Data["json"] = map[string]string{"result": "发送失败: " + errs, "statuscode": "1", "requestID": e.Data["RequestID"].(string)}
			e.ServeJSON()
			return
		}
		defer f.Close()
		attachFilename = path.Join(common.UploadPath, h.Filename)
		e.SaveToFile("attach", attachFilename)
	}

	if subject == "" || content == "" {
		e.Data["json"] = map[string]string{"result": "发送失败: 主题或者内容不能为空", "statuscode": "1", "requestID": e.Data["RequestID"].(string)}
		e.ServeJSON()
		return
	}

	if to == "" {
		e.Data["json"] = map[string]string{"result": "发送失败: 收件人不能为空", "statuscode": "1", "requestID": e.Data["RequestID"].(string)}
		e.ServeJSON()
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
		e.Data["json"] = map[string]string{"result": "send ok", "statuscode": "0", "requestID": e.Data["RequestID"].(string)}
		e.ServeJSON()
		return
	}
	mess := fmt.Sprintf("error: %s", err)
	e.Data["json"] = map[string]string{"result": mess, "statuscode": "1", "requestID": e.Data["RequestID"].(string)}
	e.ServeJSON()

}
