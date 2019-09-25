package api

type(
	SmsTemplateAdd struct{
		Name             *string     `valid:"required"  json:"name,omitempty"`
		Title            *string     `valid:"optional"  json:"title,omitempty"`
		QcloudTid        *int64        `valid:"optional"  json:"qcloud_tid,omitempty"`
		AliyunTid        *int64        `valid:"optional"  json:"aliyun_tid,omitempty"`
		Content          *string     `valid:"optional"  json:"content,omitempty"`
		Tp               *int        `valid:"optional"  json:"tp,omitempty"`
		Vaild      		 *int        `valid:"optional"  json:"valid,omitempty" `
	}

	SmsTemplate struct{
		Id               *int64      `valid:"required" json:"id,omitempty"    ` //加上type:int(11)后AUTO_INCREMENT无效
		Name             *string     `valid:"optional" json:"name,omitempty" `
		Title            *string     `valid:"optional" json:"title,omitempty"`
		QcloudTid        *int64        `valid:"optional" json:"qcloud_tid,omitempty"`
		AliyunTid        *int64        `valid:"optional" json:"aliyun_tid,omitempty"`
		Content          *string     `valid:"optional" json:"content,omitempty"`
		Tp               *int        `valid:"optional" json:"tp,omitempty"`
		Vaild      		 *int        `valid:"optional" json:"valid,omitempty" `
	}

	SmsTemplateList struct{
		Id               *int64      `valid:"optional" json:"id,omitempty"` //加上type:int(11)后AUTO_INCREMENT无效
		Name             *string     `valid:"optional" json:"name,omitempty"  `
		Title            *string     `valid:"optional" json:"title,omitempty" `

		Vaild      		*int       	 `valid:"optional" json:"valid,omitempty" `

		Page       		int64       `valid:"required" json:"page,omitempty"`
		Size       		int64       `valid:"optional" json:"size,omitempty"`
	}

	DingTemplateAdd struct{
		Name             *string     `valid:"required"  json:"name,omitempty"`
		Title            *string     `valid:"optional"  json:"title,omitempty"`
		QunName          *string        `valid:"optional"  json:"qun_name,omitempty"`
		Content          *string     `valid:"optional"  json:"content,omitempty"`
		Tp               *int        `valid:"optional"  json:"tp,omitempty"`
		Vaild      		 *int        `valid:"optional"  json:"valid,omitempty" `
	}

	DingTemplate struct{
		Id               *int64      `valid:"required" json:"id,omitempty"    ` //加上type:int(11)后AUTO_INCREMENT无效
		Name             *string     `valid:"optional" json:"name,omitempty" `
		Title            *string     `valid:"optional" json:"title,omitempty"`
		QunName          *string     `valid:"optional"  json:"qun_name,omitempty"`
		Content          *string     `valid:"optional" json:"content,omitempty"`
		Tp               *int        `valid:"optional" json:"tp,omitempty"`
		Vaild      		 *int        `valid:"optional" json:"valid,omitempty" `
	}

	DingTemplateList struct{
		Id               *int64      `valid:"optional" json:"id,omitempty"` //加上type:int(11)后AUTO_INCREMENT无效
		Name             *string     `valid:"optional" json:"name,omitempty"  `
		Title            *string     `valid:"optional" json:"title,omitempty" `

		Vaild      		*int       	 `valid:"optional" json:"valid,omitempty" `

		Page       		int64       `valid:"required" json:"page,omitempty"`
		Size       		int64       `valid:"optional" json:"size,omitempty"`
	}

	SmsRecordList struct{
		Phone      *string      `valid:"optional" json:"phone,omitempty" `
		SuccFlag      *int      `valid:"optional" json:"succ_flag,omitempty" `
		Vaild      		 *int        `valid:"optional" json:"valid,omitempty" `
		Author       *string     `valid:"optional" json:"author,omitempty"`

		Tids            []*int64        `valid:"optional" json:"tids,omitempty"`

		Page       		int64       `valid:"required" json:"page,omitempty"`
		Size       		int64       `valid:"optional" json:"size,omitempty"`
	}


)