package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"recruit/settings"
)

var (
	DB *gorm.DB
)

// Init 初始化MySQL连接r
func Init(cfg *settings.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	DB, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		return
	} else {
		sqlDB, _ := DB.DB()
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	}

	//自动建表
	//DB.AutoMigrate(&models.Arrange{})
	//DB.AutoMigrate(&models.InterviewRecord{})
	//DB.AutoMigrate(&models.Message{})
	//DB.AutoMigrate(&models.ReadMessage{})
	//DB.AutoMigrate(&models.StudentArrange{})
	//DB.AutoMigrate(&models.Student{})
	//DB.AutoMigrate(&models.TimeArrange{})
	//DB.AutoMigrate(&models.User{})

	return
}

// 自动建表方法
func createTable(dst interface{}) {
	if !DB.Migrator().HasTable(dst) {
		err := DB.AutoMigrate(dst)
		if err != nil {
			return
		}
		if DB.Migrator().HasTable(dst) {
			fmt.Println("表创建成功")
		} else {
			fmt.Println("表创建失败")
		}
	} else {
		fmt.Println("表已存在")
	}
}
