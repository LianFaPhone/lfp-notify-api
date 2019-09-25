package main

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
	"LianFaPhone/lfp-notify-api/controllers"
)

func (this *WebServer) routes()  {
	app := this.mIris
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "X-Requested-With", "X_Requested_With", "Content-Type", "Access-Token", "Accept-Language", "Api-Key", "Req-Real-Ip"},
		AllowCredentials: true,
	})

	app.Any("/", func(ctx iris.Context) {
		ctx.JSON(
			map[string]interface{}{
			"code": 0,
		})
	})


	v1 := app.Party("/v1/bk/notify", crs)
	{
			v1.Any("/", func(ctx iris.Context){
				ctx.JSON(
					map[string]interface{}{
						"code": 0,
					})
			})
			//活动，添加，list，更新，警用 全是后端的
			acParty := v1.Party("/smstemplate")
			{
				ac := new(controllers.SmsTemplate)

				acParty.Post("/add", ac.Add)
				acParty.Post("/get", ac.Get)
				acParty.Post("/gets", ac.Gets)
				acParty.Post("/del", ac.Del)
				acParty.Post("/update", ac.Update)
				acParty.Post("/list", ac.List)
			}
			//大红包，创建，list（back）
			redParty := v1.Party("/dingtemplate")
			{
				viplv := new(controllers.DingTemplate)

				redParty.Post("/add", viplv.Add)
				redParty.Post("/get", viplv.Get)
				redParty.Post("/gets", viplv.Gets)
				redParty.Post("/del", viplv.Del)
				redParty.Post("/update", viplv.Update)
				redParty.Post("/list", viplv.List)
			}
			//小红包，创建，list（back）
			emParty := v1.Party("/smsrecord")
			{
				viplv := new(controllers.SmsRecord)

				emParty.Post("/list", viplv.List)
			}

			deParty := v1.Party("/sms") //参数确定短信还是语音
			{
				shareInfo := new(controllers.SmsCtrler)

				deParty.Post("/send", shareInfo.Send)
			}

	}

	v1bk := app.Party("/v1/ft/notify", crs)
	{

		deParty := v1bk.Party("/sms") //参数确定短信还是语音
		{
			shareInfo := new(controllers.SmsCtrler)

			deParty.Post("/send", shareInfo.Send)
		}

	}


}

