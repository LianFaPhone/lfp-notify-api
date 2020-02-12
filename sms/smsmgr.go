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

//type ReqChuanglanMsg struct {
//	Account string  `json:"account"`
//	Pwd     string  `json:"password"`
//	Msg     string  `json:"msg"`
//	Report  string  `json:"report"`
//	Phone   *string `json:"phone,omitempty"`
//	Params  *string `json:"params,omitempty"`
//}
//
//type ResChuanglanMsg struct {
//	Code       string `json:"code"`
//	ErrorMsg   string `json:"errorMsg"`
//	MsgId      string `json:"msgId"`
//	Time       string `json:"time"`
//	FailNum    string `json:"failNum"`
//	successNum string `json:"successNum"`
//}
//
////成功数量
//func (this *SmsMgr) ChuanglanMutiSend(body string, phones []string, params []string) (num int, err error) {
//	if len(phones) == 0 {
//		return 0, nil
//	}
//	for i:=0; i < len(phones); i++ {
//		phones[i] = strings.TrimLeft(phones[i], "0086")
//		phones[i] = strings.TrimLeft(phones[i], "+86")
//	}
//	if len(params) == 0 {
//		phonesStr := strings.Join(phones, ",")
//		num, err = this.sendToChuanglan(body, phonesStr, "")
//	} else {
//		newParams := ""
//		for i := 0; i < len(phones); i++ {
//			newParams += phones[i]
//			for j := 0; j < len(params); j++ {
//				newParams += "," + params[j]
//			}
//			newParams += ";"
//		}
//		phonesStr := strings.Join(phones, ",")
//		newParams = strings.TrimRight(newParams, ";")
//		num, err = this.sendToChuanglan(body, phonesStr, newParams)
//	}
//	if (num == 0) && (err == nil) {
//		return len(phones), err
//	}
//	return num, err
//}
//
//func (this *SmsMgr) sendToChuanglan(body, phone string, params string) (int, error) {
//	req := new(ReqChuanglanMsg)
//	req.Account = config.GConfig.ChuangLan.Account
//	req.Pwd = config.GConfig.ChuangLan.Pwd
//	req.Report = "true"
//	req.Msg = url.QueryEscape(body)
//	req.Msg = body
//	if len(params) != 0 {
//		req.Params = new(string)
//		*req.Params = params
//	}
//	req.Phone = new(string)
//	*req.Phone = phone
//
//	bytesData, err := json.Marshal(req)
//	if err != nil {
//		return 0, err
//	}
//	//ZapLog().Info("sendToChuanglan", zap.String("json", string(bytesData)))
//	reader := bytes.NewReader(bytesData)
//	url := config.GConfig.ChuangLan.Url
//	request, err := http.NewRequest("POST", url, reader)
//	if err != nil {
//		return 0, err
//	}
//	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
//	client := http.Client{}
//	resp, err := client.Do(request)
//	if err != nil {
//		return 0, err
//	}
//	defer resp.Body.Close()
//	if resp.StatusCode != 200 {
//		return 0, fmt.Errorf("%d_%s", resp.StatusCode, resp.Status)
//	}
//	respBytes, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return 0, err
//	}
//
//	resMsg := new(ResChuanglanMsg)
//	if err = json.Unmarshal(respBytes, resMsg); err != nil {
//		return 0, err
//	}
//	if resMsg.Code != "0" {
//		return 0, fmt.Errorf("%s_%s_%s", resMsg.Code, resMsg.ErrorMsg, resMsg.MsgId)
//	}
//	//ZapLog().Info("sendToChuanglan", zap.String("res", resMsg.Code+" "+resMsg.ErrorMsg+" "+resMsg.successNum+" "+resMsg.FailNum))
//	num, _ := strconv.Atoi(resMsg.successNum)
//	return num, nil
//}
//
//func (this *SmsMgr) MutiQSend(param *api.SmsSend, temp *models.SmsTemplate) (int, error){
//	paramStr, _ := json.Marshal(param.Params)
//	var bigCode int
//	var bigErr error
//	for i := 0; i < len(param.Phone); i++ {
//		phone := strings.TrimSpace(param.Phone[i])
//		if len(phone) <= 0 {
//			continue
//		}
//		code, err := this.QSend(param.PlayTp, phone, param.Params, temp)
//		if code != 0 {
//			bigCode = code
//			bigErr = err
//			ZapLog().Sugar().Errorf("phone[%s] playTp[%d] param[%s] tempName[%s] send Fail", phone, param.PlayTp, string(paramStr), temp.Name)
//		} else {
//			ZapLog().Sugar().Infof("phone[%s] playTp[%d] param[%s] tempName[%s] send success", phone, param.PlayTp, string(paramStr), temp.Name)
//		}
//		//记录
//
//		if param.IsRecord == 1 {
//			succFlag := 0
//			if code == 0 {
//				succFlag = 1
//			}
//			author := ""
//			if param.Author != nil {
//				author = *param.Author
//			}
//			this.record(phone, author, string(paramStr), succFlag, param.PlayTp, param.ReTry, temp.Id, param.Remark)
//		}
//	}
//	return bigCode, bigErr
//}
//
//func (this *SmsMgr) QSend(playTp int, phone string, params []string, temp *models.SmsTemplate) (int, error) {
//	if playTp == 0 { //短信
//		client := qcloudsms.NewClient(this.qcOpt)
//		t := &qcloudsms.SMSSingleReq{
//			Tel:    qcloudsms.SMSTel{"86", phone},
//			Type:   0,
//			TplID:  int(*temp.QcloudTid),
//			Params: params,
//		}
//		if ok, err := client.SendSMSSingle(*t); !ok {
//			return 1, err
//		}
//		return 0, nil
//	} else {
//		client := qcloudsms.NewClient(this.qcOpt)
//		t := &qcloudsms.SMSVoiceTemplate{
//			Tel:   qcloudsms.SMSTel{"86", phone},
//			TplId: int(*temp.QcloudTid),
//			// 播放次数
//			Playtimes: 2,
//			Params:    params,
//		}
//		if ok, err := client.VoiceTemplateSend(*t); !ok {
//			return 1, err
//		}
//		return 0, nil
//	}
//	return 0, nil
//}

func (this *SmsMgr) MultiSend(param *api.SmsSend, temp *models.SmsTemplate) (int, error) {
	//检测平台
	if param.PlatformTp == nil {
		param.PlatformTp = new(int)
	}

	//调用接口
	switch *param.PlatformTp {
	case api.CONST_PlatformTp_QQ:
		return this.MutiQSend(param, temp)
	case api.CONST_PlatformTp_ChuangLan:
		smsBody := ""
		if temp.Content != nil {
			newParam :=make(map[string] interface{})
			for i:=0; i < len(param.Params);i++ {
				newParam[fmt.Sprintf("%d", i+1)] = param.Params[i]
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
		return this.ChuanglanMutiSend(smsBody, param.Phone, param.Params)
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