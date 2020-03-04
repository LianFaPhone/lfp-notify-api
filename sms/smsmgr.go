package sms

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-notify-api/api"
	"LianFaPhone/lfp-notify-api/config"
	"LianFaPhone/lfp-notify-api/models"

	"bytes"
	"fmt"
	"github.com/qichengzx/qcloudsms_go"
	"go.uber.org/zap"
	text "text/template"

)

var GSmsMgr SmsMgr

type SmsMgr struct {
	qcOpt      *qcloudsms.Options
	recordChan chan *models.SmsRecord
}

func (this *SmsMgr) Init() error {
	fmt.Println("smsMgr= ", config.GConfig.Qcloud.AppId, config.GConfig.Qcloud.AppKey, config.GConfig.Qcloud.Sign)
	this.qcOpt = qcloudsms.NewOptions(config.GConfig.Qcloud.AppId, config.GConfig.Qcloud.AppKey, config.GConfig.Qcloud.Sign)
	this.recordChan = make(chan *models.SmsRecord, 1024)

	//this.qcSms.SetDebug(true)
	this.Run()
	return nil
}

func (this *SmsMgr) record(phone, author, paramStr string, succFlag, palyTp, ReTry int, tempId *int64, remark *string) {
	p := new(models.SmsRecord).ParseAdd(tempId, &phone, &succFlag, &palyTp, &ReTry, &author, paramStr, remark)
	select {
	case this.recordChan <- p:
	default:

	}
}

func (this *SmsMgr) Run() {
	go func() {
		for {
			p, ok := <-this.recordChan
			if !ok {
				return
			}
			if err := p.Add(); err != nil {
				ZapLog().Error("smsrecord add err", zap.Error(err))
			}
		}
	}()
}

func (this *SmsMgr) MultiSend(param *api.SmsSend, temp *models.SmsTemplate) (int, error) {
	//检测平台
	if param.PlatformTp == nil {
		param.PlatformTp = new(int)
		*param.PlatformTp = api.CONST_PlatformTp_QQ
	}
	if temp.Tp == nil {
		temp.Tp = new(int)
		*temp.Tp = models.CONST_SMS_TP_Hyyx
	}

	//调用接口, 这里函数返回的code有点问题，以后修正
	switch *param.PlatformTp {
	case api.CONST_PlatformTp_QQ:
		return this.MutiQSend(param, temp)
	case api.CONST_PlatformTp_ChuangLan:
		smsBody := ""
		if temp.Content != nil {
			newParam :=make(map[string] interface{})
			for i:=0; i < len(param.Params);i++ {
				newParam[fmt.Sprintf("k%d", i+1)] = param.Params[i]

			}
			var err error
			smsBody, err = ParseTextTemplate(*temp.Content, newParam)
			if err != nil {
				//ZapLog().With(zap.Any("tempParam", param.Params), zap.Any("tempid", tmplate.Id), zap.Error(err)).Error("ParseTextTemplate err")
				//go RecordHistory(*tmplate.GroupId, Notify_Type_Sms, 0, len(this.Recipient), recordFlag)
				//return apibackend.BASERR_BASNOTIFY_TEMPLATE_PARSE_FAIL.Code(), errors.Annotate(err, "ParseTextTemplate")
				return 0,err
			}
		}
		return this.ChuanglanMutiSend(smsBody, param.Phone, param.Params, *temp.Tp)
	default:
		return this.MutiQSend(param, temp)
	}
	return 0,nil
}


func ParseTextTemplate(tmpBody string, params map[string]interface{}) (string, error) {
	tmp := text.New("smsTemp")
	_, err := tmp.Parse(tmpBody)
	if err != nil {
		return "", err
	}
	var tplBuf bytes.Buffer
	if err := tmp.Execute(&tplBuf, params); err != nil {
		return "", err
	}
	return tplBuf.String(), nil
}

//func ReplaceTemplate(tmpBody string, params []string) string{
//	lastIndex := -1
//	for i:=0; i<len(tmpBody); i++ {
//		if tmpBody[i] == '{' {
//			lastIndex = i
//			continue
//		}
//		if tmpBody[i] != '}' {
//			continue
//		}
//		if lastIndex < 0 {
//			continue
//		}
//
//	}
//}