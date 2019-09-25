package controllers

import (
	"go.uber.org/zap"
	"LianFaPhone/lfp-notify-api/api"
	"github.com/kataras/iris"
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-notify-api/models"
)

type (
	SmsTemplate struct {
		Controllers
	}
)

func (this *SmsTemplate) Add(ctx iris.Context) {
	param := new(api.SmsTemplateAdd)

	//参数检测
	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}
	uniFlag,err := new(models.SmsTemplate).UniqueByName(*param.Name)
	if err != nil {
		ZapLog().Error( "UniqueByName Add err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), err.Error())
		return
	}
	if !uniFlag {
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_EXISTS.Code(), "not unique")
		return
	}

	data,err := new(models.SmsTemplate).ParseAdd(param).Add()
	if err != nil {
		ZapLog().Error( "Activity Add err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), err.Error())
		return
	}

	this.Response(ctx, data)
}

func (this *SmsTemplate) Get(ctx iris.Context) {
	param := new(api.SmsTemplate)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	vip, err := new(models.SmsTemplate).Parse(param).Get()
	if err != nil {
		ZapLog().Error( "Activity GetForFront err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}

	this.Response(ctx, vip)
}

func (this *SmsTemplate) Gets(ctx iris.Context) {

	cmys, err := new(models.SmsTemplate).Gets()
	if err != nil {
		ZapLog().Error( "Activity GetForFront err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}

	this.Response(ctx, cmys)
}

func (this *SmsTemplate) Del(ctx iris.Context) {
	param := new(api.SmsTemplate)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}


	err = new(models.SmsTemplate).Parse(param).Del()
	if err != nil {
		ZapLog().Error( "Activity GetForFront err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}

	this.Response(ctx, nil)
}

func (this *SmsTemplate) Update(ctx iris.Context) {
	param := new(api.SmsTemplate)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	vip, err := new(models.SmsTemplate).Parse(param).Update()
	if err != nil {
		ZapLog().Error( "Activity GetForFront err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}

	this.Response(ctx, vip)
}

func (this *SmsTemplate) List(ctx iris.Context) {
	param := new(api.SmsTemplateList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	vip, err := new(models.SmsTemplate).ParseList(param).ListWithConds(param.Page, param.Size, nil)
	if err != nil {
		ZapLog().Error( "Activity GetForFront err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}

	this.Response(ctx, vip)
}