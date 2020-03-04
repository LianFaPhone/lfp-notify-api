package sms

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-notify-api/api"
	"fmt"

	//"LianFaPhone/lfp-notify-api/config"
	"LianFaPhone/lfp-notify-api/models"
	//"bytes"
	//"fmt"
	"encoding/json"
	//text "text/template"
	"github.com/qichengzx/qcloudsms_go"
	//"go.uber.org/zap"
	//"io/ioutil"
	//"net/http"
	//"net/url"
	//"strconv"
	"strings"
)

func (this *SmsMgr) MutiQSend(param *api.SmsSend, temp *models.SmsTemplate) (int, error){
	if temp.QcloudTid == nil {
		return 1, fmt.Errorf("qCloudId is nil")
	}

	paramStr, _ := json.Marshal(param.Params)
	var bigCode int
	var bigErr error
	for i := 0; i < len(param.Phone); i++ {
		phone := strings.TrimSpace(param.Phone[i])
		if len(phone) <= 0 {
			continue
		}
		code, err := this.QSend(param.PlayTp, phone, param.Params, temp)
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

func (this *SmsMgr) QSend(playTp int, phone string, params []string, temp *models.SmsTemplate) (int, error) {
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
