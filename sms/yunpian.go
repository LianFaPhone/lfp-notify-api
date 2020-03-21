package sms

import (
	"LianFaPhone/lfp-notify-api/api"
	"LianFaPhone/lfp-notify-api/models"
	. "LianFaPhone/lfp-base/log/zap"
	"fmt"
	//qcloudsms "github.com/qichengzx/qcloudsms_go"
	"encoding/json"
	"strings"
)

func (this *SmsMgr) MutiYunPianSend(body string, param *api.SmsSend, temp *models.SmsTemplate) (int, error){
	paramStr, _ := json.Marshal(param.Params)
	var bigCode int
	var bigErr error
	for i := 0; i < len(param.Phone); i++ {
		phone := strings.TrimSpace(param.Phone[i])
		if len(phone) <= 0 {
			continue
		}
		err := this.YunPianSend(param.PlayTp, phone, body)
		if err != nil {
			bigErr = err
			ZapLog().Sugar().Errorf("phone[%s] playTp[%d] param[%s] tempName[%s] send Fail", phone, param.PlayTp, string(paramStr), temp.Name)
		} else {
			ZapLog().Sugar().Infof("phone[%s] playTp[%d] param[%s] tempName[%s] send success", phone, param.PlayTp, string(paramStr), temp.Name)
		}
		//记录

		if param.IsRecord == 1 {
			succFlag := 0
			if err == nil {
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

func (this *SmsMgr) YunPianSend(playTp int, phone string,  smsBody string) ( error) {

	yParam := make(map[string]string)
	yParam["mobile"] = phone
	yParam["text"] = smsBody
	yParam["mobile_stat"] = "true"

	result:=this.yunpian.Sms().Send(yParam)
	if result.Code != 0 {
		return fmt.Errorf("%s", result.String())
	}
	
	return nil

}
