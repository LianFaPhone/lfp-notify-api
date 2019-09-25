package controllers

import (
	"go.uber.org/zap"
	"LianFaPhone/lfp-notify-api/api"
	"github.com/kataras/iris"
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-notify-api/models"
	"LianFaPhone/lfp-notify-api/sms"
)

type SmsCtrler struct{
	Controllers
}

func (this * SmsCtrler) Send(ctx iris.Context) {
	param := new(api.SmsSend)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), "param err")
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	//获取模板，
	var temp *models.SmsTemplate
	if param.TempId != nil {
		temp, err = new(models.SmsTemplate).GetById(*param.TempId)
		if err != nil {
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), "SmsTemplate.GetById err")
			ZapLog().Error( "SmsTemplate.GetById err", zap.Error(err))
			return
		}
	}else if  param.TempName != nil {
		temp, err = new(models.SmsTemplate).GetByName(*param.TempName)
		if err != nil {
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), "SmsTemplate.GetByName err")
			ZapLog().Error( "SmsTemplate.GetByName err", zap.Error(err))
			return
		}
	}
	if temp == nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_NOT_FOUND.Code(), "SmsTemplate nofind err")
		ZapLog().Error( "SmsTemplate nofind err")
		return
	}

	errCode, err := sms.GSmsMgr.MultiSend(param, temp)
	if errCode != 0 {
		this.ExceptionSerive(ctx, errCode, err.Error())
		ZapLog().Error( "Send err", zap.Error(err), zap.Any("param", *param))
		return
	}
	this.Response(ctx, nil)
}

