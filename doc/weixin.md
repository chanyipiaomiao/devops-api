# 微信报警流程

以前是通过微信企业号做报警,现在企业微信替代了微信企业号,所以现在只能通过企业微信来进行报警

## 1. [注册](https://work.weixin.qq.com/wework_admin/register_wx?from=myhome)企业微信

![注册](/doc/image/weixin/reg.png)

## 2. [登录](https://work.weixin.qq.com/wework_admin/loginpage_wx)企业微信后台

使用管理员的微信扫描登录管理后台,登录之后如下图:

![主页](/doc/image/weixin/home.png)

## 3. 邀请成员

通过 微信扫描二维码的方式, 把所有要接收报警的人邀请进来,确保都在通讯录里面,才能接收到报警

![邀请成员](/doc/image/weixin/add_p.png)

可以不用下载企业微信客户端

## 4. 添加应用

在企业应用里面创建一个应用

![创建应用](/doc/image/weixin/add_app.png)

![填写应用信息](/doc/image/weixin/add_app_person.png)

注意: 可以选择 部门/用户/标签 做为接收人,这一步很重要.

创建之后,点开应用

![应用详情](/doc/image/weixin/app_detail.png)

**记录 AgentId 和 Secret 这个会在调用API时用到**

## 5. 接收报警的人都要使用微信扫码关注微信插件的二维码

这一步至关重要,不关注的话接收不到报警信息

![微信插件](/doc/image/weixin/weixin_plugin.png)

## 6. 查看企业信息

**记录 CorpID,这个会在调用API时用到**

![企业信息](/doc/image/weixin/mycorp.png)

## 7. 查看通讯录,记录部门/用户/标签的 ID

部门ID

![部门ID](/doc/image/weixin/dep_id.png)

用户ID

![用户ID](/doc/image/weixin/person_click.png)
![用户ID](/doc/image/weixin/person_id.png)

标签ID

![标签ID](/doc/image/weixin/tag_click.png)
![标签ID](/doc/image/weixin/tag_id.png)

记录这3种ID,然后调用API就可以发送了.