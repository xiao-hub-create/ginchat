package models

import (
	"gorm.io/gorm"
	"time"
	"fmt"
	"ginchat/utils"
)
type UserBasic struct {
	gorm.Model
	Name  string
	Password string
	Phone string
	Email string
	ClientIp string
	Identity string
	ClientPort string
	LoginTime time.Time
	HeartbeatTime time.Time
	LoginOutTime time.Time
	IsLogout bool
	DeviceInfo string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic{
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
