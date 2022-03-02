/*
 * @date: 2021/12/17
 * @desc: ...
 */

package database

import (
	"DataCompare/global"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"time"
)

var redisPool *redis.Pool

// InitRedisPool 初始化 redis pool
func InitRedisPool() error {
	// 初始化redis
	redisPool = newRedisPool()
	// 检测 redis 是否可用
	_, err := GetRedisConn()
	return err
}

// GetRedisConn 获取redis连接
func GetRedisConn() (redis.Conn, error) {
	conn := redisPool.Get()
	if err := conn.Err(); err != nil {
		global.Logger.Error("redis 连接错误", zap.String("msg", err.Error()))
		return nil, err
	}
	return conn, nil
}

func newRedisPool() *redis.Pool {
	redisInfo := global.ServerConfig.DatabaseInfo.RedisInfo
	addr := fmt.Sprintf("%s:%d", redisInfo.Host, redisInfo.Port)
	db := redisInfo.DB
	username := redisInfo.Username
	password := redisInfo.Password
	connectTimeout := redisInfo.ConnectTimeout
	poolMaxIdleConns := redisInfo.PoolMaxIdleConns
	poolMaxOpenConns := redisInfo.PoolMaxOpenConns
	poolConnMaxLifetime := redisInfo.PoolConnMaxLifetime

	var dailOptions = []redis.DialOption{
		redis.DialDatabase(db),
		redis.DialUsername(username),
		redis.DialPassword(password),
		redis.DialConnectTimeout(time.Duration(connectTimeout) * time.Millisecond),
	}
	return &redis.Pool{
		MaxIdle:     poolMaxIdleConns,                                 //最大空闲连接数
		MaxActive:   poolMaxOpenConns,                                 //允许分配最大连接数
		IdleTimeout: time.Duration(poolConnMaxLifetime) * time.Second, // 空闲连接存活时间
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", addr, dailOptions...)
			if err != nil {
				global.Logger.Error("redis 连接错误", zap.String("msg", err.Error()))
				return nil, err
			}
			return conn, err
		},
	}
}
