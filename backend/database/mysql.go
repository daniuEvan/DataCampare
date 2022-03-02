/*
 * @date: 2021/12/15
 * @desc: ...
 */

package database

import (
	"DataCompare/global"
	"DataCompare/handler/model"
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var MysqlDB *gorm.DB

func GetMysqlDB(ctx context.Context) (*gorm.DB, error) {
	// WithContext 实际是调用 db.Session(&Session{Context: ctx})，每次创建新 Session，各 db 操作之间互不影响
	dbManger, _ := MysqlDB.DB()
	err := dbManger.Ping()
	if err != nil {
		global.Logger.Warn("数据库连接异常,正在重新初始化数据库连接:", zap.String("error", err.Error()))
		err := dbManger.Close()
		if err != nil {
			global.Logger.Error("数据库连接关闭异常:", zap.String("error", err.Error()))
			return nil, err
		}
		err = InitMysqlDB()
		if err != nil {
			global.Logger.Error("初始化数据库连接异常:", zap.String("error", err.Error()))
			return nil, err
		}
	}
	return MysqlDB.WithContext(ctx), nil
}

// InitMysqlDB 初始化mysql
func InitMysqlDB() (err error) {
	mysqlDB, err := newMysqlConn()
	if err != nil {
		global.Logger.Error("初始化数据库连接异常:", zap.String("error", err.Error()))
		return err
	}
	err = mysqlDB.AutoMigrate(model.ModelsArr...)
	if err != nil {
		global.Logger.Error("数据库迁移失败:", zap.String("error", err.Error()))
		return err
	}
	MysqlDB = mysqlDB
	return nil
}

// newMysqlConn 获取数据库连接
func newMysqlConn() (db *gorm.DB, err error) {
	mysqlInfo := global.ServerConfig.DatabaseInfo.MysqlInfo
	databasePoolStatus := global.ServerConfig.OrmDatabasePoolInfo.Status
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Username,
		mysqlInfo.Password,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.DBName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s_", global.ServerConfig.OrmInfo.TablePrefix),
			SingularTable: true,
		},
	})
	if err != nil {
		global.Logger.Error("连接数据库异常:", zap.String("error", err.Error()))
		return nil, err
	}
	//db.
	if databasePoolStatus == "disable" {
		return db, err
	}
	// 构建数据库连接池
	if err := buildOrmDatabasePool(db); err != nil {
		return nil, err
	}
	return db, nil
}

// 构建 orm 数据库连接池
func buildOrmDatabasePool(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		global.Logger.Error("构建数据库连接池异常:", zap.String("error", err.Error()))
		return err
	}
	databasePoolInfo := global.ServerConfig.OrmDatabasePoolInfo
	sqlDB.SetMaxIdleConns(databasePoolInfo.MaxIdleConns)                                    // 空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(databasePoolInfo.MaxOpenConns)                                    // 数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(databasePoolInfo.ConnMaxLifetime)) // 连接可复用的最大时间。
	return nil
}
