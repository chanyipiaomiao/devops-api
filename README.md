# devops-api

Golang + Beego编写, 提供一些开发/运维常见操作的HTTP API接口，方便使用

# 主要功能

- 微信报警
- 2步验证(Google Authenticator验证)
- 密码存储
- 发送邮件
- 生成随机密码
- 获取字符串的MD5值
- 生成验证密码

更多功能,正在思考中...

# 目录
- [安装使用](#安装使用)
- [依赖](#依赖)
- [功能列表](#功能列表)
	- [微信报警](#微信报警)
		- [发送消息](#发送消息)
	- [2步验证](#2步验证)
		- [启用2步验证](#启用2步验证)
		- [验证6位数字](#验证google-authenticator或是其他的类似的app生成的6位数字)
		- [禁用2步验证](#禁用2步验证)
	- [密码存储](#密码存储)
		- [存储](#存储)
		- [查询](#查询)
		- [删除](#删除)
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

直接[**下载**](https://github.com/chanyipiaomiao/devops-api/releases)二级制文件即可

## 使用

**注意: 如果配置文件app.conf中, security->enableToken 的值是 false, 可以跳过 步骤2 和 步骤3, 默认为true, 如果是false 可以不用在请求头里面添加 DEVOPS-API-TOKEN 头**


1. 自定义配置(**该步骤可选**)

	- 配置监听端口
	- 配置上传目录,确保有权限写入
	- 配置日志目录,默认情况当前目录下,确保有权限写入
	- 配置日志最大存放天数
	- 配置邮箱地址、端口、用户名、密码
	- 配置是否启用token验证
	- 配置jwt token签名字符串,请自行生成修改
	- 配置微信报警的配置, corpID、warningAppAgentID、warningAppSecret,可参考文档[设置微信报警流程](/doc/weixin.md)

2. 首先初始化, 会生成root token，该root token 管理其他的token(**该步骤可选**)

```sh
./devops-api init
```

```sh
注意: 忘记root token, 可以使用以下重新生成

./devops-api init --refresh-root-token
```

3. 使用root token 生成普通的token，用于验证请求(**该步骤可选**)

```sh
./devops-api token --create=名称 --root-token=上边的root token

注意：忘记token，重新生成即可
```

4. 启动服务

```sh
./devops-api server --mode=prod

--mode 指定程序运行模式, prod: 生产模式 dev: 开发模式, 默认是开发模式

也可以在配置文件app.conf中修改 runmode=prod 也可以设置为生产模式
```

```sh
使用生成token就能愉快的访问API了

注意: token必须放到请求头里面,名称必须是: DEVOPS-API-TOKEN
```

5. 获取帮助

```sh
./devops-api --help
```

6. 备份数据库文件

```sh
./devops-api backup --filepath=备份文件路径
```

[返回到目录](#目录)

# 依赖

```go
go get github.com/astaxie/beego
go get github.com/robfig/cron
go get github.com/chanyipiaomiao/hltool
go get gopkg.in/alecthomas/kingpin.v2
go get github.com/satori/go.uuid
go get github.com/sec51/twofactor
go get github.com/chanyipiaomiao/weixin-kit
```

[返回到目录](#目录)

# API

## 微信报警

点击查看设置[微信报警流程](/doc/weixin.md)

还需要在 conf/app.conf 中配置已下3个参数
- corpID 				公司或者组织的ID
- warningAppAgentID		报警应用的ID
- warningAppSecret    	报警应用的密钥

这3个参数可以在企业微信后台管理页面可以看到，详情可以查看上面的文档。


#### 发送消息

```sh
POST /api/v1/sendmsg/weixin

msgType: text 		消息类型,目前只支持文本消息
toTag:   标签ID 		在企业微信后台可以查看到
toUser： 用户ID 		在企业微信后台可以查看到
toParty: 部门ID 		在企业微信后台可以查看到
text：   要发送的文本消息
```

[返回到目录](#目录)

## 2步验证

2步验证,可以方便的集成到系统中,只需要调用3个API即可.

### 2步验证api接口

#### 启用2步验证

```sh
GET /api/v1/twostepauth/enable?username=用户名&issuer=发行者

username email或者是用户名
issuer 可以是比如 公司的域名/公司的代号等
```
返回结果:

```sh
{
    "enable": true,
    "key": "656C7AAU556TAMNONWZXLPEYTCXR3QE2",
    "qrImage": "/static/download/qr/lei.he.png",
    "requestId": "ee3145bf-c329-4830-947b-69ef74a269f5",
    "statuscode": 0
}

enable 		启用成功
qrImage 	2步验证 二维码图片路径
key     	没办法扫描二维码时可以手动添加
statuscode  	返回0,代表成功,其他失败 
```

[返回到目录](#目录)

#### 验证Google Authenticator或是其他的类似的APP生成的6位数字

```sh
POST /api/v1/twostepauth/auth

username: 用户名
issuer:   发行者
token:    6位数字的验证码
```
返回结果:

```sh
{
    "auth": true,
    "requestId": "2f9aa9b5-2c02-4c7f-af4e-3c1d931eb7aa",
    "statuscode": 0,
    "username": "lei.he"
}

auth: 验证成功true, 不成功false
```

[返回到目录](#目录)

#### 禁用2步验证

```sh
GET /api/v1/twostepauth/disable?username=用户名
```

返回结果:

```sh
{
    "disable": true,
    "requestId": "4f73c93c-ae99-4582-81b1-81ce75133599",
    "statuscode": 0,
    "username": "lei.he"
}

disable 禁用成功,删除二维码图片,从数据库中删除该用户
```

[返回到目录](#目录)


## 密码存储

### 密码存储API

#### 存储

```sh
POST /api/v1/storepass/update

设置头部: Content-Type: application/json

请求体内容格式:
{
  "uniqueId": "10.10.1.2",
  "password": [
    {"name": "root", "password": "444"},
    {"name": "user1", "password": "333"}
  ]
}
uniqueId 唯一标识

注意字段名称必须是以上格式
```
返回结果:
```sh
{
    "requestID": "13b4dc78-7f28-4477-a7c6-e319cb1c00ea",
    "statuscode": 0,
    "update": true
}
```

[返回到目录](#目录)

#### 查询
```sh
GET /api/v1/storepass/get?id=10.10.1.2,1.1.1.1

多个标识逗号分开
```

返回结果:
```sh
{
    "get": true,
    "password": {
        "1.1.1.1": {},
        "10.10.1.2": {
            "root": "444",
            "user1": "333"
        }
    },
    "requestID": "35ce7716-2721-4c0b-82ee-f258fdb9a6c1",
    "statuscode": 0
}

1.1.1.1 在数据库中不存在则返回的是空
```
[返回到目录](#目录)

#### 删除

```sh
GET /api/v1/storepass/delete?id=10.10.1.2,1.1.1.1

如果请求的标识在数据库中不存从则返回空
```

返回结果:
```sh
{
    "delete": true,
    "id": "10.10.1.2,1.1.1.1",
    "requestID": "562ca3e3-52bd-43ff-8ba2-09b20b928baa",
    "statuscode": 0
}
```

[返回到目录](#目录)


## 发送邮件 

#### 发送邮件api接口

```sh
POST /api/v1/sendmsg/mail
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

```sh
curl -X POST \
  http://127.0.0.1:8080/api/v1/sendmsg/mail \
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

url = "http://127.0.0.1:8080/api/v1/sendmsg/mail"
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

url = "http://127.0.0.1:8080/api/v1/sendmsg/mail"
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
	url := "http://127.0.0.1:8080/api/v1/sendmsg/mail"
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
	url := "http://127.0.0.1:8080/api/v1/sendmsg/mail"
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

```sh
GET /api/v1/password/generation
```

生成随机密码 默认32位, 可以添加如下参数

- length 指定长度
- name  指定多个名称，为多个名称生成密码
- specialChar 是否添加特殊字符(!@#%$*.=)到密码里面 specialChar=yes 添加，其他不添加

[返回到目录](#目录)

### Curl 示例

```sh
curl http://127.0.0.1:8080/api/v1/password/generation?length=64 \
  -H 'DEVOPS-API-TOKEN: 生成的Token'
```

[返回到目录](#目录)

### 别名使用

#### 首先设置一下环境变量

```sh
export DEVOPS_API_TOKEN=生成的Token
```

Linux

```sh
alias genpwd="curl -H \"DEVOPS-API-TOKEN: ${DEVOPS_API_TOKEN}\" http://127.0.0.1:8080/api/v1/password/generation?length=64;echo"
alias genpwdspecial="curl -H \"DEVOPS-API-TOKEN: ${DEVOPS_API_TOKEN}\" http://127.0.0.1:8080/api/v1/password/generation?length=64&specialChar=yes;echo"
```
Mac 可能需要把?=&转义一下

```sh
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

```sh
GET /api/v1/md5?string=123456
```

可以写一个shell脚本命令行传入字符串，返回MD5值

[返回到目录](#目录)


## 生成验证密码

### 生成验证密码

#### 生成验证密码api接口

```sh
GET /api/v1/password/manualGenAuthPassword
```

手动生成验证密码，该功能是 生成一个32位的密码，然后通过钉钉或者邮件通知，别人可以拿这个密码到本程序验证是否正确，这个密码也会定时的清除，不会一直生效 执行成功返回 {"manualGenAuthPassword": true}

### 验证上面的密码是否正确

#### 验证密码api接口

```sh
POST /api/v1/password/authPassword
```

验证密码是否正确，就上一个API生成的密码, 参数: password 执行成功返回 {"auth", true}

[返回到目录](#目录)

## 获取程序自身版本信息

#### 获取程序自身版本信息api接口

```sh
GET /version  
```

[返回到目录](#目录)