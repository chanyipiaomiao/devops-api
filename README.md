# devops-api

Golang + Beego编写, 提供一些开发/运维常见操作的HTTP API接口，提供开发/运维常用操作的HTTP API接口: 手机归属地查询、IP地址查询、工作日节假日判断、微信报警、钉钉报警、2步验证、密码存储、发送邮件、生成随机密码等功能

# 主要功能

- 手机归属地查询
- IP地址查询
- 工作日节假日判断
- 微信报警
- 钉钉报警
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
    - [手机归属地查询](#手机归属地查询)
    - [IP地址查询](#ip地址查询)
	- [工作日节假日判断](#工作日节假日判断)
		- [设置节假日和工作日](#设置节假日和工作日)
		- [判断给定的日期是节假日/工作日/周末](#判断给定的日期是节假日工作日周末)
	- [微信报警](#微信报警)
		- [发送微信消息](#发送微信消息)
	- [钉钉报警](#钉钉报警)
		- [发送钉钉消息](#发送钉钉消息)
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

配置文件说明:

- app.conf          主配置文件
- dev.conf          开发时的配置文件
- prod.conf         线上生产运行时的配置文件
- authpassword.conf 验证密码配置文件
- db.conf           数据库配置文件
- email.conf        邮箱用户名和密码配置
- log.conf          日志相关配置
- security.conf     安全相关的配置
- twostep.conf      2步验证相关
- weixin.conf       微信报警相关
- phone.conf        手机归属地查询配置

主配置文件 app.conf 通过include的方式加载其他的配置文件


**注意: 如果配置文件security.conf中, security->enableToken 的值是 false, 可以跳过 步骤2 和 步骤3, 默认为true, 如果是false 可以不用在请求头里面添加 DEVOPS-API-TOKEN 头**


1. 自定义配置(**该步骤可选**)

	- 配置监听端口
	- 配置上传目录,确保有权限写入
	- 配置日志目录,默认情况当前目录下,确保有权限写入
	- 配置日志最大存放天数
	- 配置邮箱地址、端口、用户名、密码
	- 配置是否启用token验证
	- 配置jwt token签名字符串,请自行生成修改
	- 配置微信报警的配置, corpID、warningAppAgentID、warningAppSecret,可参考文档[设置微信报警流程](/doc/weixin.md)
	- 配置钉钉机器人的URL(**可选**)

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

# 开发依赖

```go
go get github.com/astaxie/beego
go get github.com/robfig/cron
go get gopkg.in/alecthomas/kingpin.v2
go get github.com/satori/go.uuid
go get github.com/sec51/twofactor
go get github.com/tidwall/gjson
go get github.com/chanyipiaomiao/hltool
go get github.com/chanyipiaomiao/weixin-kit
go get github.com/chanyipiaomiao/cal
go get github.com/chanyipiaomiao/ip2region/binding/golang
go get github.com/chanyipiaomiao/phonedata
```

[返回到目录](#目录)

# API

## 手机归属地查询

本功能使用了 [xluohome](https://github.com/xluohome/phonedata) 项目提供的手机归属地数据库

首先进入到script目录，执行 get_phone_dat.sh 来下载数据文件，可以定期执行脚本获取最新的.

```sh
/api/v1/queryphone?phone=手机号

phone 要查询的手机号
```

返回:

```sh
{
    "data": {
        "AreaZone": "021",
        "CardType": "中国移动",
        "City": "上海",
        "PhoneNum": "xxxxxxxxxx",
        "Province": "上海",
        "ZipCode": "200000"
    },
    "entryType": "Query Phone Location",
    "errmsg": "",
    "requestId": "0860edaa-db7f-46ee-ac89-d41eeb2ed80d",
    "statuscode": 0
}
```

## IP地址查询

本功能使用了 [狮子的魂](https://gitee.com/lionsoul/ip2region) 项目提供的IP地址数据库.

首先进入到script目录, 执行 gen_ip_region.sh 脚本, 来下载IP地址数据库, 可以定期执行脚本获取最新的.


```sh
GET /api/v1/queryip?ip=xxx.xxx.xxx.xxx

ip  要查询的IP地址
```

返回: 

```sh
{
    "data": {
        "ip": "xxx.xxx.xxx.xxx",
        "ipInfo": {
            "CityId": 995,
            "Country": "中国",
            "Region": "0",
            "Province": "上海",
            "City": "上海市",
            "ISP": "电信"
        }
    },
    "entryType": "Query IP",
    "errmsg": "",
    "requestId": "6aae483e-5c72-4cb7-bbb7-50089e2da4d3",
    "statuscode": 0
}
```

[返回到目录](#目录)

## 工作日节假日判断

#### 设置节假日和工作日

```sh
POST /api/v1/holiworkday

设置头部: Content-Type: application/json

内容: 参考下面的模板
```
节假日和工作日模板json, 每年都要根据国内的放假安排提前做一下设置,毕竟国内放假安排不是固定的,日期不足2位必须补0

workday 是指放假安排中的调整上班的日期

```sh
{
    "year": "2018",
    "holiday": [
        {
            "name": "yuandan",
            "zh_name": "元旦",
            "start_time": "2018-01-01",
            "end_time": "2018-01-01"
        },
        {
            "name": "chunjie",
            "zh_name": "春节",
            "start_time": "2018-02-15",
            "end_time": "2018-02-21"
        },
        {
            "name": "qingming",
            "zh_name": "清明节",
            "start_time": "2018-04-05",
            "end_time": "2018-04-07"
        },
        {
            "name": "laodong",
            "zh_name": "劳动节",
            "start_time": "2018-04-29",
            "end_time": "2018-05-01"
        },
        {
            "name": "duanwu",
            "zh_name": "端午节",
            "start_time": "2018-06-16",
            "end_time": "2018-06-18"
        },
        {
            "name": "zhongqiu",
            "zh_name": "中秋节",
            "start_time": "2018-09-22",
            "end_time": "2018-09-24"
        },
        {
            "name": "guoqing",
            "zh_name": "国庆节",
            "start_time": "2018-10-01",
            "end_time": "2018-10-07"
        }
    ],
    "workday": [
        "2018-02-11",
        "2018-02-24",
        "2018-04-08",
        "2018-04-28",
        "2018-09-29",
        "2018-09-30"
    ]
}
```

[返回到目录](#目录)

#### 判断给定的日期是节假日工作日周末

```sh
GET /api/v1/holiworkday?date=2018-08-25

date: 判断的日期, 日期格式不足2位必须补0
```

返回:
```sh
{
    "data": {
        "date": "2018-08-25",
        "dateType": "weekend"
    },
    "entryType": "Get Holiday/Workday",
    "errmsg": "",
    "requestId": "562444c2-1a48-4c69-9ed1-d2553dea3cba",
    "statuscode": 0
}

dateType: 有3种 workday: 工作日  holiday: 节假日  weekend: 周末
```

[返回到目录](#目录)

## 微信报警

点击查看设置[微信报警流程](/doc/weixin.md)

还需要在 conf/app.conf 中配置已下3个参数
- corpID 				公司或者组织的ID
- warningAppAgentID		报警应用的ID
- warningAppSecret    	报警应用的密钥

这3个参数可以在企业微信后台管理页面可以看到，详情可以查看上面的文档。


#### 发送微信消息

```sh
POST /api/v1/sendmsg/weixin

msgType: text 		消息类型,目前只支持文本消息
toTag:   标签ID 		在企业微信后台可以查看到
toUser： 用户ID 		在企业微信后台可以查看到
toParty: 部门ID 		在企业微信后台可以查看到
msg：    要发送的消息
```

[返回到目录](#目录)

## 钉钉报警

需要在conf/app.conf 配置钉钉机器人的URL，或者是在POST中传递url参数指定钉钉机器人地址

#### 发送钉钉消息

```sh
POST /api/v1/sendmsg/dingding

msgType： text|markdown
msg:      要发送的消息
title:    发送markdown消息时需要指定此参数，指定标题
url:      可以指定钉钉机器人的URL，这样就不用在conf/app.conf 配置钉钉机器人的URL
```

注意: 在发送markdown消息时，markdown的语法可以查看[官方文档](https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.xuaZtG&treeId=257&articleId=105735&docType=1#s2)

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
    "data": {
        "key": "xxxxxxxxxxxx",
        "qrImage": "/static/download/qr/xxxx.png"
    },
    "entryType": "TwoStepAuth",
    "errmsg": "",
    "requestId": "e55fc2ea-4465-4a4f-aba7-f73272900b03",
    "statuscode": 0
}

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
    "data": {
        "auth": false,
        "issuer": "xxxxx",
        "username": "xxxxxxx"
    },
    "entryType": "TwoStepAuth",
    "errmsg": "Tokens mismatch.",
    "requestId": "5529567b-1c5a-4e04-aaa0-5a86ac19ca94",
    "statuscode": 1
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
    "data": {
        "disable": "yes",
        "username": "xxxxx"
    },
    "entryType": "TwoStepAuth",
    "errmsg": "",
    "requestId": "4451ef02-f245-466a-8bb4-172238f47c50",
    "statuscode": 0
}

disable: yes 禁用成功,删除二维码图片,从数据库中删除该用户
```

[返回到目录](#目录)


## 密码存储

### 密码存储API

#### 存储

```sh
POST /api/v1/storepass

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
    "data": {
        "update": "ok"
    },
    "entryType": "Store Password",
    "errmsg": "",
    "requestId": "2494ad20-ca52-4d3e-8e8e-6dd0d6289f4a",
    "statuscode": 0
}
```

[返回到目录](#目录)

#### 查询
```sh
GET /api/v1/storepass/10.10.1.2,1.1.1.1

多个标识逗号分开
```

返回结果:
```sh
{
    "data": {
        "password": {
            "1.1.1.1": {},
            "10.10.1.2": {
                "root": "444",
                "user1": "333"
            }
        }
    },
    "entryType": "Store Password",
    "errmsg": "",
    "requestId": "d5f61efe-8c22-4e4e-9d97-c343eb1e7f58",
    "statuscode": 0
}

1.1.1.1 在数据库中不存在则返回的是空
```
[返回到目录](#目录)

#### 删除

```sh
DELETE /api/v1/storepass/10.10.1.2,1.1.1.1

如果请求的标识在数据库中不存从则返回空
```

返回结果:
```sh
{
    "data": {
        "delete": "ok",
        "id": "10.10.1.2,1.1.1.1"
    },
    "entryType": "Store Password",
    "errmsg": "",
    "requestId": "29ed3301-319f-45b6-8cbd-934becc7c5cb",
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
GET /api/v1/md5?rawstr=123456
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
GET /api/v1/version  
```

[返回到目录](#目录)
