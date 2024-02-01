package initGorm

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitGorm gorm初始化
func InitGorm(MysqlDataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlDataSource), &gorm.Config{})
	if err != nil {
		logx.Errorf("gorm.Open(mysql.Open(MysqlDataSource) , &gorm.Config{}) is failed : %v", err)
		panic(err)
	}
	return db
}
