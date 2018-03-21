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
				beego.NSRouter("/update", &controllers.StorePasswordController{}, "post:Update"),
				beego.NSRouter("/delete", &controllers.StorePasswordController{}, "get:Delete"),
				beego.NSRouter("/get", &controllers.StorePasswordController{}, "get:Get"),
			),
		),
	)
	beego.AddNamespace(apins)

	versions := beego.NewNamespace("/version",
		beego.NSRouter("", &controllers.VersionController{}),
	)

	beego.AddNamespace(versions)
}
