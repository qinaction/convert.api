package dbs

import (
	"convert.api/libs/configs"
	"fmt"
	"github.com/outreach-golang/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var GMysql map[string]*gorm.DB

func init() {
	GMysql = make(map[string]*gorm.DB)
}

func InitMysql() error {
	var confs = configs.GConfig.Database.Mysql

	for _, conf := range confs {
		cli, err := gorm.Open(
			mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&timeout=5s&loc=Asia%%2FShanghai", conf.Username, conf.Password, conf.Address, conf.Dbname)),
			&gorm.Config{
				QueryFields: true,
			},
		)
		if err != nil {
			return err
		}

		err = cli.Use(&logger.TracePlugin{})
		if err != nil {
			return err
		}

		db, _ := cli.DB()

		db.SetMaxOpenConns(conf.MaxOpenConns)
		db.SetMaxIdleConns(10)
		db.SetConnMaxLifetime(time.Hour)

		GMysql[conf.Asname] = cli
	}

	return nil
}
