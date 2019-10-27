package sms

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-notify-api/api"
	"LianFaPhone/lfp-notify-api/config"
	"LianFaPhone/lfp-notify-api/models"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/qichengzx/qcloudsms_go"
	"go.uber.org/zap"
	"strings"
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

func (this *SmsMgr) Send(playTp int, phone string, params []string, temp *models.SmsTemplate) (int, error) {
	if playTp == 0 { //短信
		client := qcloudsms.NewClient(this.qcOpt)
		t := &qcloudsms.SMSSingleReq{
			Tel:    qcloudsms.SMSTel{"86", phone},
			Type:   0,
			TplID:  int(*temp.QcloudTid),
			Params: params,
		}
		if ok, err := client.SendSMSSingle(*t); !ok {
			return 1, err
		}
		return 0, nil
	} else {
		client := qcloudsms.NewClient(this.qcOpt)
		t := &qcloudsms.SMSVoiceTemplate{
			Tel:   qcloudsms.SMSTel{"86", phone},
			TplId: int(*temp.QcloudTid),
			// 播放次数
			Playtimes: 2,
			Params:    params,
		}
		if ok, err := client.VoiceTemplateSend(*t); !ok {
			return 1, err
		}
		return 0, nil
	}
	return 0, nil
}

func (this *SmsMgr) MultiSend(param *api.SmsSend, temp *models.SmsTemplate) (int, error) {
	//检测平台

	//调用接口

	paramStr, _ := json.Marshal(param.Params)
	var bigCode int
	var bigErr error
	for i := 0; i < len(param.Phone); i++ {
		phone := strings.TrimSpace(param.Phone[i])
		if len(phone) <= 0 {
			continue
		}
		code, err := this.Send(param.PlayTp, phone, param.Params, temp)
		if code != 0 {
			bigCode = code
			bigErr = err
			ZapLog().Sugar().Errorf("phone[%s] playTp[%d] param[%s] tempName[%s] send Fail", phone, param.PlayTp, string(paramStr), temp.Name)
		} else {
			ZapLog().Sugar().Infof("phone[%s] playTp[%d] param[%s] tempName[%s] send success", phone, param.PlayTp, string(paramStr), temp.Name)
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
			this.record(phone, author, string(paramStr), succFlag, param.PlayTp, param.ReTry, temp.Id, param.Remark)
		}
	}
	return bigCode, bigErr
}
