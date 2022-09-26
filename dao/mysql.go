package dao

// database access object
import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Mysql *gorm.DB

// Init 初始化数据库的一些操作
func Init() (err error) {
	dsn := loadConfigParameters()
	Mysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed in gorm.Open: %s\n", err.Error())
		return
	}
	err = Mysql.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(&User{}, &Video{}, &Comment{})
	ResetOnlineStatus()
	return
}
