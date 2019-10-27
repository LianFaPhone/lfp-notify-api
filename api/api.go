package api

type (
	SmsSend struct {
		TempName *string  `valid:"optional" json:"temp_name,omitempty"` //groupname+lang 联合使用
		TempId   *int64   `valid:"optional" json:"temp_id,omitempty"`   //groupid+lang联合使用
		Params   []string `valid:"-" json:"params,omitempty"`           //optional
		Phone    []string `valid:"optional" json:"phone,omitempty"`     //require
		Author   *string  `valid:"optional" json:"author,omitempty"`
		PlayTp   int      `valid:"optional" json:"play_tp,omitempty"`   // 0短信，1语音
		IsRecord int      `valid:"optional" json:"is_record,omitempty"` // 0不记录，1记录
		ReTry    int      `valid:"optional" json:"retry,omitempty"`     // 0不，1是
		Remark   *string  `valid:"optional" json:"remark,omitempty"`
	}

	DingSend struct {
		TempName *string                `valid:"optional" json:"group_name,omitempty"` //groupname+lang 联合使用
		TempId   *int64                 `valid:"optional" json:"group_id,omitempty"`   //groupid+lang联合使用
		Params   map[string]interface{} `valid:"-" json:"params,omitempty"`            //optional
		QunName  *string                `valid:"optional" json:"qun_name,omitempty"`
	}
)
