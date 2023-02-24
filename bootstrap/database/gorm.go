package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"skeleton/bootstrap/logger"
	"skeleton/config"
	"time"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.YAML.Database.Username,
		config.YAML.Database.Password,
		config.YAML.Database.Host,
		config.YAML.Database.Port,
		config.YAML.Database.DbName,
		config.YAML.Database.Encode,
	)

	zapLogger := zapgorm2.New(logger.Logger)
	zapLogger.SetAsDefault()

	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: zapLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: config.YAML.Database.Prefix,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败 -> %v\n", err))
		return
	}

	sqlDB, _ := DB.DB()
	// sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}
