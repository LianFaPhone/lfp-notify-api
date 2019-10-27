package models

import (
	"LianFaPhone/lfp-notify-api/api"
	"LianFaPhone/lfp-notify-api/common"
	"LianFaPhone/lfp-notify-api/db"
	"github.com/jinzhu/gorm"
)

type (
	SmsTemplate struct {
		Id        *int64  `json:"id,omitempty"              gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
		Name      *string `json:"name,omitempty"            gorm:"column:name;type:varchar(50);not null;unique;index"`
		Title     *string `json:"title,omitempty"           gorm:"column:title;type:varchar(50)"`
		QcloudTid *int64  `json:"qcloud_tid,omitempty"      gorm:"column:qcloud_tid;type:bigint(20);"`
		AliyunTid *int64  `json:"aliyun_tid,omitempty"       gorm:"column:aliyun_tid;type:bigint(20);"`
		Content   *string `json:"content,omitempty"         gorm:"column:content;type:varchar(500)"`
		Tp        *int    `json:"tp,omitempty"              gorm:"column:tp;type:int(11)""`
		Table
	}
)

func (this *SmsTemplate) TableName() string {
	return "notify_sms_tempalte"
}

func (this *SmsTemplate) ParseAdd(p *api.SmsTemplateAdd) *SmsTemplate {
	c := &SmsTemplate{
		Name:      p.Name,
		Title:     p.Title,
		QcloudTid: p.QcloudTid,
		AliyunTid: p.AliyunTid,
		Content:   p.Content,
		Tp:        p.Tp,
	}
	c.Valid = p.Vaild
	if c.Valid == nil {
		c.Valid = new(int)
		*c.Valid = 1
	}
	return c
}

func (this *SmsTemplate) Parse(p *api.SmsTemplate) *SmsTemplate {
	c := &SmsTemplate{
		Id:        p.Id,
		Name:      p.Name,
		Title:     p.Title,
		QcloudTid: p.QcloudTid,
		AliyunTid: p.AliyunTid,
		Content:   p.Content,
		Tp:        p.Tp,
	}
	c.Valid = p.Vaild
	return c
}

func (this *SmsTemplate) ParseList(p *api.SmsTemplateList) *SmsTemplate {
	c := &SmsTemplate{
		Id:    p.Id,
		Name:  p.Name,
		Title: p.Title,
	}
	c.Valid = p.Vaild
	return c
}

func (this *SmsTemplate) Add() (*SmsTemplate, error) {
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return nil, err
	}
	acty := new(SmsTemplate)
	err = db.GDbMgr.Get().Where("id = ?", *this.Id).Last(acty).Error
	if err != nil {
		return nil, err
	}
	return acty, nil
}

func (this *SmsTemplate) UniqueByName(name string) (bool, error) {
	count := 0
	err := db.GDbMgr.Get().Model(this).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (this *SmsTemplate) Get() (*SmsTemplate, error) {
	acty := new(SmsTemplate)
	err := db.GDbMgr.Get().Where("id = ?", *this.Id).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *SmsTemplate) GetByName(name string) (*SmsTemplate, error) {
	acty := new(SmsTemplate)
	err := db.GDbMgr.Get().Where("name = ?", name).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *SmsTemplate) Gets() ([]*SmsTemplate, error) {
	var acty []*SmsTemplate
	err := db.GDbMgr.Get().Find(&acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *SmsTemplate) GetById(id int64) (*SmsTemplate, error) {
	acty := new(SmsTemplate)
	err := db.GDbMgr.Get().Where("id = ?", id).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *SmsTemplate) GetsByTId(id []*int64) ([]*SmsTemplate, error) {
	var acty []*SmsTemplate
	err := db.GDbMgr.Get().Where("qcloud_tid in (?) or aliyun_tid in (?)", id, id).Select("id").Find(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *SmsTemplate) Del() error {
	err := db.GDbMgr.Get().Where("id = ?", *this.Id).Delete(&SmsTemplate{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}

func (this *SmsTemplate) Update() (*SmsTemplate, error) {
	if err := db.GDbMgr.Get().Model(this).Updates(this).Error; err != nil {
		return nil, err
	}
	acty := new(SmsTemplate)
	err := db.GDbMgr.Get().Where("id = ?", *this.Id).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return acty, err
}

func (this *SmsTemplate) ListWithConds(page, size int64, condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*SmsTemplate
	query := db.GDbMgr.Get().Where(this)
	for i := 0; i < len(condPair); i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}
	query = query.Order("valid desc").Order("id desc")

	return new(common.Result).PageQuery(query, &SmsTemplate{}, &list, page, size, nil, "")
}
