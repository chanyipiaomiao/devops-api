# devops-api

提供一些运维常见操作的 http 接口，方便使用

# 目录
- [安装使用](#安装使用)
- [依赖](#依赖)
- [功能列表](#功能列表)
	- [发送邮件](#发送邮件)
		- [api接口](#发送邮件api接口)
		- [Curl 示例](#curl-示例)
		- [Python 文本邮件/html邮件 示例](#python-文本邮件html邮件-示例)
		- [Python 发送带附件邮件 示例](#python-发送带附件邮件-示例)
		- [Go 文本邮件/html邮件 示例](#go-文本邮件html邮件-示例)
		- [Go 发送带附件的邮件 示例](#go-发送带附件的邮件-示例)
	- [生成随机密码](#生成随机密码)
		- [api接口](#生成随机密码api接口)
		- [Curl 示例](#curl-示例)
		- [别名使用](#别名使用)
		- [Python示例](#python示例)
		- [Go示例](#go示例)
	- [获取字符串的MD5值](#获取字符串的md5值)
		- [api接口](#获取字符串的md5值api接口)
	- [生成验证密码](#生成验证密码)
		- [生成验证密码](#生成验证密码)
			- [api接口](#生成验证密码api接口)
		- [验证上面的密码是否正确](#验证上面的密码是否正确)
			- [api接口](#验证密码api接口)
	- [获取程序自身版本信息](#获取程序自身版本信息)
		- [api接口](#获取程序自身版本信息api接口)


# 安装使用

## 安装

直接下载二级制文件即可

## 使用

**注意: 如果配置文件app.conf中, security->enableToken 的值是 false, 可以跳过 步骤2 和 步骤3, 默认为true**


1. 自定义配置(**该步骤可选**)

	- 配置监听端口
	- 配置上传目录,确保有权限写入
	- 配置日志目录,默认情况当前目录下,确保有权限写入
	- 配置日志最大存放天数
	- 配置邮箱地址、端口、用户名、密码
	- 配置是否启用token验证
	- 配置jwt token签名字符串,请自行生成修改

2. 首先初始化, 会生成root token，该root token 管理其他的token(**该步骤可选**)

```shell
./devops-api init
```

```shell
注意: 忘记root token, 可以使用以下重新生成

./devops-api init --refresh-root-token
```

3. 使用root token 生成普通的token，用于验证请求(**该步骤可选**)

```shell
./devops-api token --create=名称 --root-token=上边的root token

注意：忘记token，重新生成即可
```

4. 启动服务

```shell
./devops-api server 
```

```shell
使用生成token就能愉快的访问API了

注意: token必须放到请求头里面,名称必须是: DEVOPS-API-TOKEN
```

5. 获取帮助

```shell
./devops-api --help
```

[返回到目录](#目录)

# 依赖

```go
go get github.com/astaxie/beego
go get github.com/robfig/cron
go get github.com/chanyipiaomiao/hltool
go get gopkg.in/alecthomas/kingpin.v2
go get -u github.com/satori/go.uuid
```

[返回到目录](#目录)

# API

## 发送邮件 

#### 发送邮件api接口

```shell
POST /api/v1/sendmail
```

- DEVOPS-API-TOKEN 指定API-TOKEN
- subject 主题
- content 内容
- type text/plain|text/html
- to 收件人多个用逗号分开
- cc 抄送人多个用逗号分开
- attach 指定附件的路径
- isattach 是否有附件 true|false

[返回到目录](#目录)

### Curl 示例

```shell
curl -X POST \
  http://127.0.0.1:8080/api/v1/sendmail \
  -H 'DEVOPS-API-TOKEN: 生成Token' \
  -H 'content-type: multipart/form-data' \
  -F subject=haha \
  -F 'content=测试一下哈' \
  -F type=text/plain \
  -F to=xxxx@qq.com,xxx@163.com \
  -F cc=xxxx@qq.com,xxx@163.com \
  -F 'attach=@C:\Users\41176\Desktop\1111.txt' \
  -F isattach=true
```
[返回到目录](#目录)

### Python 文本邮件/html邮件 示例

记得安装 requests

```python
pip install requests
```

```python
import requests

url = "http://127.0.0.1:8080/api/v1/sendmail"
token = "生成的Token"
headers = {'DEVOPS-API-TOKEN': token}
payload = {
    "subject": u"测试邮件",
    "content": u"测试邮件内容",  # u"<h1>测试邮件内容</h1>",
    "type":    "text/plain",    # "text/html"
    "to":      "xxxx@qq.com,xxx@163.com",
    "cc":      "xxxx@qq.com,xxx@163.com",
}
response = requests.post(url, data=payload, headers=headers)
print(response.json())
```
[返回到目录](#目录)

### Python 发送带附件邮件 示例

```python
import requests

url = "http://127.0.0.1:8080/api/v1/sendmail"
token = "生成的Token"
payload = {
    "subject": u"测试邮件",
    "content": u"测试邮件内容",
    "type":    "text/plain",
    "to":      "xxxx@qq.com",
    "cc":      "xxxx@163.com",
    "isattach": "true"
}
headers = {'DEVOPS-API-TOKEN': token}
files = {'attach': open('文件的路径', 'rb')} # 键名必须是 attach
response = requests.post(url, data=payload, headers=headers, files=files)
print(response.json())
```

[返回到目录](#目录)

### Go 文本邮件/html邮件 示例

记得安装 grequests

```go
go get github.com/levigross/grequests
```

```go
package main

import (
	"log"

	"github.com/levigross/grequests"
)

func main() {
	url := "http://127.0.0.1:8080/api/v1/sendmail"
	token := "生成的Token"
	o := &grequests.RequestOptions{
		Headers: map[string]string{
			"DEVOPS-API-TOKEN": token,
		},
		Data: map[string]string{
			"subject": "测试邮件",
			"content": "测试邮件内容", // <h1>测试邮件内容</h1>
			"type":    "text/plain", // text/html
            "to":      "xxxx@qq.com,xxx@163.com",
            "cc":      "xxxx@qq.com,xxx@163.com",
		},
	}
	resp, err := grequests.Post(url, o)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	log.Println(resp.String())
}

```

[返回到目录](#目录)

### Go 发送带附件的邮件 示例

```go
package main

import (
	"log"
	"os"

	"github.com/levigross/grequests"
)

func main() {
	url := "http://127.0.0.1:8080/api/v1/sendmail"
	token := "生成的token"

	fd, err := os.Open("文件路径")
	if err != nil {
		log.Fatalln("open file error: ", err)
	}

	o := &grequests.RequestOptions{
		Headers: map[string]string{
			"DEVOPS-API-TOKEN": token,
		},
		Data: map[string]string{
			"subject":  "测试邮件",
			"content":  "测试邮件内容", 
			"type":     "text/plain", 
			"to":       "xxx@qq.com",
			"cc":       "xxxx@163.com",
			"isattach": "true",
		},
		Files: []grequests.FileUpload{
			grequests.FileUpload{
				FileName:     "上传之后生成的文件名",
				FileContents: fd,
				FieldName:    "attach", // FieldName 必须是 attach
			},
		},
	}
	resp, err := grequests.Post(url, o)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	log.Println(resp.String())
}

```

[返回到目录](#目录)

## 生成随机密码

#### 生成随机密码api接口

```shell
GET /api/v1/password/generation
```

生成随机密码 默认32位, 可以添加如下参数

- length 指定长度
- name  指定多个名称，为多个名称生成密码
- specialChar 是否添加特殊字符(!@#%$*.=)到密码里面 specialChar=yes 添加，其他不添加

[返回到目录](#目录)

### Curl 示例

```shell
curl http://127.0.0.1:8080/api/v1/password/generation?length=64 \
  -H 'DEVOPS-API-TOKEN: 生成的Token'
```

[返回到目录](#目录)

### 别名使用

#### 首先设置一下环境变量

```shell
export DEVOPS_API_TOKEN=生成的Token
```

Linux

```shell
alias genpwd="curl -H \"DEVOPS-API-TOKEN: ${DEVOPS_API_TOKEN}\" http://127.0.0.1:8080/api/v1/password/generation?length=64;echo"
alias genpwdspecial="curl -H \"DEVOPS-API-TOKEN: ${DEVOPS_API_TOKEN}\" http://127.0.0.1:8080/api/v1/password/generation?length=64&specialChar=yes;echo"
```
Mac 可能需要把?=&转义一下

```shell
alias genpwd="curl -H \"DEVOPS-API-TOKEN: ${DEVOPS_API_TOKEN}\" http://127.0.0.1:8080/api/v1/password/generation\?length\=64;echo"
alias genpwdspecial="curl -H \"DEVOPS-API-TOKEN: ${DEVOPS_API_TOKEN}\" http://127.0.0.1:8080/api/v1/password/generation\?length\=64\&specialChar\=yes;echo"
```

[返回到目录](#目录)

### Python示例

```python
import requests
import pprint

url = "http://127.0.0.1:8080/api/v1/password/generation"
token = "生成的Token"
payload = {
    "length": 64,                   # 指定密码长度
    "name": "user1,user2,user3",    # 可以为多个名称生成密码
    "specialChar": "yes"            # 密码中是否带有特殊字符
}
headers = {'DEVOPS-API-TOKEN': token}
response = requests.get(url, params=payload, headers=headers)
pprint.pprint(response.json())
```

[返回到目录](#目录)

### Go示例

```go
package main

import (
	"log"

	"github.com/levigross/grequests"
)

func main() {
	url := "http://127.0.0.1:8080/api/v1/password/generation"
	token := "生成的Token"

	o := &grequests.RequestOptions{
		Headers: map[string]string{
			"DEVOPS-API-TOKEN": token,
		},
		Params: map[string]string{
			"length":      "64",                // 指定密码长度
			"name":        "user1,user2,user3", // 可以为多个名称生成密码
			"specialChar": "yes",               // 密码中是否带有特殊字符
		},
	}
	resp, err := grequests.Get(url, o)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	log.Println(resp.String())
}

```

[返回到目录](#目录)


## 获取字符串的MD5值

### 获取字符串的MD5值api接口

```shell
GET /api/v1/md5?string=123456
```

可以写一个shell脚本命令行传入字符串，返回MD5值

[返回到目录](#目录)


## 生成验证密码

### 生成验证密码

#### 生成验证密码api接口

```shell
GET /api/v1/password/manualGenAuthPassword
```

手动生成验证密码，该功能是 生成一个32位的密码，然后通过钉钉或者邮件通知，别人可以拿这个密码到本程序验证是否正确，这个密码也会定时的清除，不会一直生效 执行成功返回 {"manualGenAuthPassword": true}

### 验证上面的密码是否正确

#### 验证密码api接口

```shell
POST /api/v1/password/authPassword
```

验证密码是否正确，就上一个API生成的密码, 参数: password 执行成功返回 {"auth", true}

[返回到目录](#目录)

## 获取程序自身版本信息

#### 获取程序自身版本信息api接口

```shell
GET /version  
```

[返回到目录](#目录)