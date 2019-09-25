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
	SmsRecord struct {
		Controllers
	}
)


func (this *SmsRecord) List(ctx iris.Context) {
	param := new(api.SmsRecordList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error( "param err", zap.Error(err))
		return
	}

	cond := make([]*models.SqlPairCondition, 0 )
	if param.Tids != nil && len(param.Tids) > 0 {
		temArr ,_ := new(models.SmsTemplate).GetsByTId(param.Tids)
		idArr := make([]*int64, 0, len(temArr))
		for i:=0; i < len(temArr); i++ {
			idArr = append(idArr, temArr[i].Id)
		}
		if len(idArr) > 0 {

		}
		cond  = append(cond, &models.SqlPairCondition{"temp_id in (?)", idArr})
	}


	vip, err := new(models.SmsRecord).ParseList(param).ListWithConds(param.Page, param.Size, cond)
	if err != nil {
		ZapLog().Error( "Activity GetForFront err", zap.Error(err))
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}

	this.Response(ctx, vip)
}

