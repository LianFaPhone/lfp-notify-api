package sms

import (
	"github.com/qichengzx/qcloudsms_go"
	"LianFaPhone/lfp-notify-api/models"
	"LianFaPhone/lfp-notify-api/api"
	"strings"
	"github.com/gin-gonic/gin/json"
	"LianFaPhone/lfp-notify-api/config"
	"go.uber.org/zap"
	. "LianFaPhone/lfp-base/log/zap"
)

var GSmsMgr SmsMgr

type SmsMgr struct{
	qcOpt *qcloudsms.Options
	recordChan chan *models.SmsRecord
}

func (this *SmsMgr)Init() error {
	this.qcOpt = qcloudsms.NewOptions(config.GConfig.Qcloud.AppId,config.GConfig.Qcloud.AppKey,config.GConfig.Qcloud.Sign)
	this.recordChan = make(chan *models.SmsRecord, 1024)

	//this.qcSms.SetDebug(true)
	this.Run()
	return nil
}

func (this *SmsMgr) record(phone,author,paramStr string,succFlag, palyTp,ReTry int, tempId *int64) {
	p := new(models.SmsRecord).ParseAdd(tempId, &phone, &succFlag, &palyTp,&ReTry, &author, paramStr)
	select{
		case this.recordChan <- p:
	default:

	}
}

func (this *SmsMgr) Run() {
	go func(){
		for{
			p,ok := <- this.recordChan
			if !ok {
				return
			}
			if err := p.Add(); err != nil {
				ZapLog().Error( "smsrecord add err", zap.Error(err))
			}
		}
	}()
}

func (this *SmsMgr) Send(playTp int, phone string,params []string, temp *models.SmsTemplate)(int,error){
	return 0, nil
	if playTp == 0 { //短信
		client := qcloudsms.NewClient(this.qcOpt)
		t := &qcloudsms.SMSSingleReq {
			Tel :   qcloudsms.SMSTel{"86", phone},
			Type  : 0,
			TplID  : int(*temp.QcloudTid),
			Params: params,
		}
		if ok,err := client.SendSMSSingle(*t); !ok {
			return 1,err
		}
		return 0, nil
	}else{
		client := qcloudsms.NewClient(this.qcOpt)
		t := &qcloudsms.SMSVoiceTemplate {
			Tel :   qcloudsms.SMSTel{"86", phone},
			TplId  : int(*temp.QcloudTid),
			// 播放次数
			Playtimes : 2,
			Params: params,
		}
		if ok, err := client.VoiceTemplateSend(*t); !ok {
			return 1,err
		}
		return 0, nil
	}
	return 0, nil
}

func (this *SmsMgr) MultiSend(param *api.SmsSend, temp *models.SmsTemplate)(int,error) {
	//检测平台

	//调用接口

	paramStr, _ := json.Marshal(param.Params)
	var bigCode int
	var bigErr error
	for i:=0; i < len(param.Phone); i++ {
		phone := strings.TrimSpace(param.Phone[i])
		if len(phone) <= 0  {
			continue
		}
		code, err := this.Send(param.PlayTp, phone,param.Params, temp)
		if code != 0 {
			bigCode = code
			bigErr = err
		}
		//记录

		if param.IsRecord == 1 {
			succFlag := 0
			if code == 0 {
				succFlag = 1
			}
			author := ""
			if param.Author != nil {
				author = *param.Author
			}
			this.record(phone, author, string(paramStr), succFlag, param.PlayTp, param.ReTry,temp.Id)
		}
	}
	return bigCode, bigErr
}