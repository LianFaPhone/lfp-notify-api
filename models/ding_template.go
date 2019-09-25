package models

import (
	"LianFaPhone/lfp-notify-api/db"
	"github.com/jinzhu/gorm"
	"LianFaPhone/lfp-notify-api/api"
	"LianFaPhone/lfp-notify-api/common"
)

type(
	DingTemplate struct{
		Id               *int64      `json:"id,omitempty"              gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
		Name             *string     `json:"name,omitempty"            gorm:"column:name;type:varchar(50);not null;unique;index"`
		Title            *string     `json:"title,omitempty"           gorm:"column:title;type:varchar(50)"`
		QunName          *string     `json:"qun_name,omitempty"     gorm:"column:qun_name;type:varchar(40);"`
		Content          *string     `json:"content,omitempty"         gorm:"column:content;type:varchar(500)"`
		Tp               *int        `json:"tp,omitempty"              gorm:"column:tp;type:int(11)""`
		Table
	}
)

func (this *DingTemplate) TableName() string {
	return "notify_ding_tempalte"
}

func (this *DingTemplate) ParseAdd(p *api.DingTemplateAdd) *DingTemplate {
	c := &DingTemplate{
		Name      : p.Name,
		Title   : p.Title,
		QunName : p.QunName,
		Content :p.Content,
		Tp : p.Tp,
	}
	c.Valid = p.Vaild
	if c.Valid == nil {
		c.Valid = new(int)
		*c.Valid = 1
	}
	return c
}

func (this *DingTemplate) Parse(p *api.DingTemplate) *DingTemplate {
	c := &DingTemplate{
		Id:       p.Id,
		Name      : p.Name,
		Title   : p.Title,
		QunName : p.QunName,
		Content :p.Content,
		Tp : p.Tp,
	}
	c.Valid = p.Vaild
	return c
}

func (this *DingTemplate) ParseList(p *api.DingTemplateList) *DingTemplate {
	c := &DingTemplate{
		Id:       p.Id,
		Name      : p.Name,
		Title   : p.Title,
	}
	c.Valid = p.Vaild
	return c
}

func (this *DingTemplate) Add() (*DingTemplate, error){
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return nil, err
	}
	acty := new(DingTemplate)
	err = db.GDbMgr.Get().Where("id = ?", *this.Id).Last(acty).Error
	if err!= nil {
		return nil, err
	}
	return acty,nil
}


func (this *DingTemplate) Get() (*DingTemplate, error) {
	acty := new(DingTemplate)
	err := db.GDbMgr.Get().Where(this).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return acty,err
}

func (this *DingTemplate)Gets() ([]*DingTemplate, error) {
	var acty [] *DingTemplate
	err := db.GDbMgr.Get().Find(&acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return acty,err
}

func (this *DingTemplate) GetById( id int64) (*DingTemplate, error) {
	acty := new(DingTemplate)
	err := db.GDbMgr.Get().Where("id = ?", id).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return acty,err
}

func (this *DingTemplate) Del() (error) {
	err := db.GDbMgr.Get().Where("id = ?", *this.Id).Delete(&DingTemplate{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return  err
}

func (this *DingTemplate) Update() (*DingTemplate, error) {
	if err := db.GDbMgr.Get().Model(this).Updates(this).Error; err != nil {
		return nil,err
	}
	acty := new(DingTemplate)
	err := db.GDbMgr.Get().Where("id = ?", *this.Id).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return acty,err
}

func (this *DingTemplate) ListWithConds(page, size int64,  condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*DingTemplate
	query := db.GDbMgr.Get().Where(this)
	for i:=0; i < len(condPair);i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}
	query = query.Order("valid desc").Order("id")

	return new(common.Result).PageQuery(query, &DingTemplate{}, &list, page, size, nil, "")
}