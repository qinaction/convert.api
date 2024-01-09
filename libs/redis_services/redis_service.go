package redis_services

import (
	"convert.api/libs/configs"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

var (
	GRedis map[string]*redis.Client
)

func init() {
	GRedis = make(map[string]*redis.Client)
}

func InitRedis() (err error) {
	var confs = configs.GConfig.Database.Redis

	for _, conf := range confs {
		client := redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%d", conf.Address, conf.Port),
			DB:           conf.Db,
			Password:     conf.Password,
			PoolSize:     conf.PoolSize,
			MinIdleConns: conf.MinIdleConns,
			DialTimeout:  time.Duration(conf.DialTimeout) * time.Second,
		})

		if _, err = client.Ping().Result(); err != nil {
			return err
		}

		GRedis[conf.Asname] = client
	}

	return
}
