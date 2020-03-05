package sms

import (
	//. "LianFaPhone/lfp-base/log/zap"
	//"LianFaPhone/lfp-notify-api/api"
	"LianFaPhone/lfp-notify-api/config"
	"LianFaPhone/lfp-notify-api/models"
	"LianFaPhone/lfp-notify-api/api"

	//"LianFaPhone/lfp-notify-api/models"
	"bytes"
	"fmt"
	"encoding/json"
	//text "text/template"
	//"github.com/qichengzx/qcloudsms_go"
	//"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ReqChuanglanMsg struct {
	Account string  `json:"account"`
	Pwd     string  `json:"password"`
	Msg     string  `json:"msg"`
	Report  string  `json:"report"`
	Phone   *string `json:"phone,omitempty"`
	Params  *string `json:"params,omitempty"`
}

type ResChuanglanMsg struct {
	Code       string `json:"code"`
	ErrorMsg   string `json:"errorMsg"`
	MsgId      string `json:"msgId"`
	Time       string `json:"time"`
	FailNum    string `json:"failNum"`
	successNum string `json:"successNum"`
}

//成功数量
func (this *SmsMgr) ChuanglanMutiSend(body string, param *api.SmsSend, temp *models.SmsTemplate) (num int, err error) {
	if len(param.Phone) == 0 {
		return 0, nil
	}
	for i:=0; i < len(param.Phone); i++ {
		param.Phone[i] = strings.TrimLeft(param.Phone[i], "0086")
		param.Phone[i] = strings.TrimLeft(param.Phone[i], "+86")
	}
	phonesStr := strings.Join(param.Phone, ",")
	if len(param.Params) == 0 {
	//	phonesStr := strings.Join(param.Phone, ",")
		num, err = this.sendToChuanglan(body, phonesStr, "", *temp.Tp)
	} else {
		newParams := ""
		for i := 0; i < len(param.Phone); i++ {
			newParams += param.Phone[i]
			for j := 0; j < len(param.Params); j++ {
				newParams += "," + param.Params[j]
			}
			newParams += ";"
		}
		//phonesStr := strings.Join(param.Phone, ",")
		newParams = strings.TrimRight(newParams, ";")
		num, err = this.sendToChuanglan(body, phonesStr, newParams, *temp.Tp)
	}
	//if param.IsRecord == 1 {
	//	succFlag := 0
	//	if err == nil {
	//		succFlag = 1
	//	}
	//	author := ""
	//	if param.Author != nil {
	//		author = *param.Author
	//	}
	//	this.record(phonesStr, author, string(paramStr), succFlag, param.PlayTp, param.ReTry, temp.Id, param.Remark)
	//}
	if (num == 0) && (err == nil) {
		return len(param.Phone), err
	}
	return num, err
}

func (this *SmsMgr) sendToChuanglan(body, phone string, params string, tp int) (int, error) {
	req := new(ReqChuanglanMsg)
	if tp == models.CONST_SMS_TP_Hyyx {
		req.Account = config.GConfig.ChuangLan.HyyxAccount
		req.Pwd = config.GConfig.ChuangLan.HyyxPwd
	}else if tp == models.CONST_SMS_TP_Yzm {
		req.Account = config.GConfig.ChuangLan.YzmAccount
		req.Pwd = config.GConfig.ChuangLan.YzmPwd
	}

	req.Report = "true"
	req.Msg = url.QueryEscape(body)
	req.Msg = body
	if len(params) != 0 {
		req.Params = new(string)
		*req.Params = params
	}
	req.Phone = new(string)
	*req.Phone = phone

	bytesData, err := json.Marshal(req)
	if err != nil {
		return 0, err
	}
	//ZapLog().Info("sendToChuanglan", zap.String("json", string(bytesData)))
	reader := bytes.NewReader(bytesData)
	url := config.GConfig.ChuangLan.Url
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return 0, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("%d_%s", resp.StatusCode, resp.Status)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	resMsg := new(ResChuanglanMsg)
	if err = json.Unmarshal(respBytes, resMsg); err != nil {
		return 0, err
	}
	if resMsg.Code != "0" {
		return 0, fmt.Errorf("%s_%s_%s", resMsg.Code, resMsg.ErrorMsg, resMsg.MsgId)
	}
	//ZapLog().Info("sendToChuanglan", zap.String("res", resMsg.Code+" "+resMsg.ErrorMsg+" "+resMsg.successNum+" "+resMsg.FailNum))
	num, _ := strconv.Atoi(resMsg.successNum)
	return num, nil
}
