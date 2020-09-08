package service

import (
	model "gateway/models"
	InterfaceEntity "gateway/models/InterfaceEntity"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "gateway.db")
	// defer DB.Close()
	if err = DB.AutoMigrate(model.Models...).Error; nil != err {
		log.Fatal("auto migrate tables failed: " + err.Error())
	}
	// 初始化汇总数据
	var sumInfo InterfaceEntity.SumInfo
	var chartInfo InterfaceEntity.ChartInfo
	if err := DB.Find(&sumInfo).Error; err != nil {
		sumInfo := InterfaceEntity.SumInfo{
			ServerSum:  0,
			WarningSum: 0,
			RequestSum: 0,
			FailSum:    0,
		}
		DB.Create(&sumInfo)
	}
	if err := DB.Find(&chartInfo).Error; err != nil {
		chartInfo := InterfaceEntity.ChartInfo{
			Time:    time.Now().Format("2006/01/02"),
			Total:   0,
			Success: 0,
			Fail:    0,
		}
		DB.Create(&chartInfo)
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(50)
	DB.DB().SetConnMaxLifetime(5 * time.Minute)
	DB.LogMode(true)
}
