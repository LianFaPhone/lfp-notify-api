package models

import (
	"github.com/jinzhu/gorm"
	"LianFaPhone/lfp-notify-api/api"
	"LianFaPhone/lfp-notify-api/common"
	"LianFaPhone/lfp-notify-api/db"
	"time"
)

type (
	SmsRecord struct{
		Id         *int64       `json:"id,omitempty"         gorm:"column:id;primary_key;AUTO_INCREMENT:1;not null"` //加上type:int(11)后AUTO_INCREMENT无效
		TempId     *int64      `json:"temp_id,omitempty"   gorm:"column:temp_id;type:bigint(20)"`
		Phone      *string      `json:"phone,omitempty"   gorm:"column:phone;type:varchar(20)"`
		SuccFlag      *int      `json:"succ_flag,omitempty"   gorm:"column:succ_flag;type:tinyint(2)"`
		Param      *string     `json:"param,omitempty"   gorm:"column:param;type:varchar(100)"`
		PlayTp         *int      `json:"play_tp,omitempty"   gorm:"column:play_tp;type:tinyint(2)"` //0 sms;  1 voice
		Author       *string   `json:"author,omitempty"        gorm:"column:author;type:varchar(10)"`
		ReTry      *int          `json:"retry,omitempty"   gorm:"column:retry;type:tinyint(2)"`
		LastSendAt   *int64    `json:"last_send_at,omitempty"   gorm:"column:last_send_at;type:bigint(20)"`
		TryCount     *int      `json:"try_count,omitempty"   gorm:"column:try_count;type:tinyint(2)"`
		Table
	}
)

func (this *SmsRecord) TableName() string {
	return "notify_sms_record"
}

func (this *SmsRecord) ParseAdd(tempId *int64, Phone *string, SuccFlag, PlayTp, ReTry *int, Author  *string, paramStr string) *SmsRecord {
	c := &SmsRecord{
		TempId: tempId,
		Phone: Phone,
		SuccFlag: SuccFlag,
		Param: &paramStr,
		PlayTp: PlayTp,
		Author: Author,
		LastSendAt: new(int64),
		TryCount: new(int),
		ReTry: ReTry,
	}
	*c.TryCount = 1
	*c.LastSendAt = time.Now().Unix()
	if c.Valid == nil {
		c.Valid = new(int)
		*c.Valid = 1
	}
	return c
}

//func (this *SmsRecord) Parse(p *api.SmsRecord) *SmsRecord {
//	c := &SmsRecord{
//		Id:   p.Id,
//		Name: p.Name,
////		CompanyId: p.CompanyId,
////		EpeNum: p.EpeNum,
//	}
//	c.Valid = p.Vaild
//	return c
//}


func (this *SmsRecord) ParseList(p *api.SmsRecordList) *SmsRecord {
	c := &SmsRecord{
		Phone: p.Phone,
		SuccFlag: p.SuccFlag,
		Author: p.Author,
	}
	c.Valid = p.Vaild
	return c
}


func (this *SmsRecord) Add() ( error){
	err := db.GDbMgr.Get().Create(this).Error
	if err != nil {
		return  err
	}
	return nil
}

//func (this *Department) ExistByCompanyId(CompanyId int64) (bool, error){
//	count := 0
//	err := db.GDbMgr.Get().Where("company_id = ?", CompanyId).Count(&count).Error
//	if err!= nil {
//		return false, err
//	}
//	return count > 0,nil
//}


func (this *SmsRecord) Get() (*SmsRecord, error) {
	acty := new(SmsRecord)
	err := db.GDbMgr.Get().Where(this).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return acty,err
}

func (this *SmsRecord) Gets() ([]*SmsRecord, error) {
	var acty []*SmsRecord
	err := db.GDbMgr.Get().Where(this).Find(&acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return acty,err
}

func (this *SmsRecord) GetById(id int64) (*SmsRecord, error) {
	acty := new(SmsRecord)
	err := db.GDbMgr.Get().Where("id = ?", id).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return acty,err
}

func (this *SmsRecord) Del() (error) {
	err := db.GDbMgr.Get().Where("id = ?", *this.Id).Delete(&SmsRecord{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}


func (this *SmsRecord) Update() (*SmsRecord, error) {
	if err := db.GDbMgr.Get().Model(this).Updates(this).Error; err != nil {
		return nil,err
	}
	acty := new(SmsRecord)
	err := db.GDbMgr.Get().Where("id = ?", *this.Id).Last(acty).Error
	if err == gorm.ErrRecordNotFound {
		return nil,nil
	}
	return acty,err
}

func (this *SmsRecord) ListWithConds(page, size int64,  condPair []*SqlPairCondition) (*common.Result, error) {
	var list []*SmsRecord
	query := db.GDbMgr.Get().Where(this)
	for i:=0; i < len(condPair);i++ {
		if condPair[i] == nil {
			continue
		}
		query = query.Where(condPair[i].Key, condPair[i].Value)
	}
	query = query.Order("valid desc").Order("id desc")

	return new(common.Result).PageQuery(query, &SmsRecord{}, &list, page, size, nil, "")
}