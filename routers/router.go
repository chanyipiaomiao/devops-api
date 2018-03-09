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
			beego.NSNamespace("/sendmail",
				beego.NSRouter("", &controllers.EmailController{}, "post:SendMail"),
			),
			beego.NSNamespace("/md5",
				beego.NSRouter("", &controllers.MD5Controller{}),
			),
			beego.NSNamespace("/twostepauth",
				beego.NSRouter("/enable", &controllers.TwoStepAuthController{}, "get:Enable"),
				beego.NSRouter("/auth", &controllers.TwoStepAuthController{}, "post:Auth"),
				beego.NSRouter("/disable", &controllers.TwoStepAuthController{}, "get:Disable"),
			),
			beego.NSNamespace("/managePassword",
				beego.NSRouter("/update", &controllers.ManagePassword{}, "post:Update"),
				beego.NSRouter("/delete", &controllers.ManagePassword{}, "post:Delete"),
			),
		),
	)
	beego.AddNamespace(apins)

	versions := beego.NewNamespace("/version",
		beego.NSRouter("", &controllers.VersionController{}),
	)

	beego.AddNamespace(versions)
}
