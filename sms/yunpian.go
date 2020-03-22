package sms

import (
	"LianFaPhone/lfp-notify-api/api"
	"LianFaPhone/lfp-notify-api/common"
	"LianFaPhone/lfp-notify-api/config"
	"LianFaPhone/lfp-notify-api/models"
	. "LianFaPhone/lfp-base/log/zap"
	"fmt"
	"net/url"

	//qcloudsms "github.com/qichengzx/qcloudsms_go"
	"encoding/json"
	"strings"
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
)

func (this *SmsMgr) MutiYunPianSend(body string, param *api.SmsSend, temp *models.SmsTemplate) (int, error){
	paramStr, _ := json.Marshal(param.Params)
	var bigCode int
	var bigErr error
	body = strings.TrimSuffix(body, "退订回T")
	body = strings.TrimSuffix(body, "\r\n")
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

	//yParam := make(map[string]string)
	//yParam[ypclnt.MOBILE] = phone
	//yParam[ypclnt.TEXT] = smsBody
	//yParam["mobile_stat"] = "true"

	formBody := make(url.Values)
	formBody.Set(ypclnt.APIKEY, config.GConfig.YunPian.ApiKey)
	formBody.Set(ypclnt.MOBILE, phone)
	formBody.Set(ypclnt.TEXT, smsBody)
	formBody.Set("mobile_stat", "true")
	resBytes,err := common.HttpFormSend("https://sms.yunpian.com/v2/sms/single_send.json", formBody, "POST", nil)
	//result:=this.yunpian.Sms().Send(yParam)
	//if result.Code != 0 {
	//	return fmt.Errorf("%s", result.String())
	//}
	if err != nil {
		return err
	}

	res := new(YunPianResponse)

	if err = json.Unmarshal(resBytes, res); err != nil {
		return err
	}
	if res.Code != 0 {
		return fmt.Errorf("%d-%s", res.Code, res.Msg)
	}

	return nil

}

type YunPianResponse struct{
	Code	int   `json:"code"`
	Msg	    string `json:"msg"`
	//Count	int     `json:"count"`
	//Fee	    float64	  `json:"fee"`
	//Unit	string	`json:"unit"`
	//Mobile	string	`json:"mobile"`
	//Sid	    int64        `json:"sid"`
}