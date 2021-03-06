package models

import (
	"LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-notify-api/config"
	"LianFaPhone/lfp-notify-api/db"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

func InitDbTable() {
	log.ZapLog().Info("start InitDbTable")
	if !config.GConfig.Db.Debug {
		log.ZapLog().Info("end InitDbTable")
		return
	}
	err := db.GDbMgr.Get().Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&DingTemplate{}, &SmsRecord{}, &SmsTemplate{}).Error
	if err != nil {
		log.ZapLog().Error("AutoMigrate err", zap.Error(err))
	}
	log.ZapLog().Info("end InitDbTable")
}
