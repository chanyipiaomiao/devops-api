package routers

import (
	"devops-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	apins := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSNamespace("/password",
				beego.NSRouter("/generation", &controllers.PasswordController{}, "get:GenPassword"),
				beego.NSRouter("/authPassword", &controllers.PasswordController{}, "post:AuthGenPassword"),
				beego.NSRouter("/manualGenAuthPassword", &controllers.PasswordController{}, "get:ManualGenAuthPassword"),
			),
			beego.NSNamespace("/sendmsg",
				beego.NSNamespace("/mail",
					beego.NSRouter("", &controllers.EmailController{}, "post:SendMail"),
				),
				beego.NSNamespace("/weixin",
					beego.NSRouter("", &controllers.WeixinController{}, "post:SendMessage"),
				),
				beego.NSNamespace("/dingding",
					beego.NSRouter("", &controllers.DingdingController{}, "post:SendMessage"),
				),
			),
			beego.NSNamespace("/md5",
				beego.NSRouter("", &controllers.MD5Controller{}),
			),
			beego.NSNamespace("/twostepauth",
				beego.NSRouter("/enable", &controllers.TwoStepAuthController{}, "get:Enable"),
				beego.NSRouter("/auth", &controllers.TwoStepAuthController{}, "post:Auth"),
				beego.NSRouter("/disable", &controllers.TwoStepAuthController{}, "get:Disable"),
			),
			beego.NSNamespace("/storepass",
				beego.NSRouter("/?:id(.+)", &controllers.StorePasswordController{}),
			),
			beego.NSNamespace("/holiworkday",
				beego.NSRouter("", &controllers.HolidayController{}),
			),
			beego.NSNamespace("/queryip",
				beego.NSRouter("", &controllers.QueryIPController{}),
			),
			beego.NSNamespace("/queryphone",
				beego.NSRouter("", &controllers.PhoneController{}),
			),
			beego.NSNamespace("/version",
				beego.NSRouter("", &controllers.VersionController{}),
			),
		),
	)
	beego.AddNamespace(apins)

}
