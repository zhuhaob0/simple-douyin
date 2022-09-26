package dao

import (
	"fmt"
	"log"
)

// ResetOnlineStatus 重置数据库中用户在线状态为下线
func ResetOnlineStatus() {
	err := Mysql.Model(&User{}).Where("id>?", 0).Update("online", 0).Error
	if err != nil {
		log.Println("ResetOnlineStatus: ", err)
	}
}

// 导入数据库配置参数
func loadConfigParameters() string {
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	DBName := "douyin"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, DBName)
	return dsn
}
