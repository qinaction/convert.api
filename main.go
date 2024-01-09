package main

import (
	"context"
	"convert.api/libs/common"
	"convert.api/libs/configs"
	"convert.api/libs/logger"
	"convert.api/router"
	"flag"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"github.com/outreach-golang/etcd"
	"log"
	"os"
	"strconv"
)

func main() {

	var (
		err             error
		envCl           string
		consulAddressCl string
		etcdAddressCl   = os.Getenv("etcd_address")
		etcdSSL, _      = strconv.Atoi(os.Getenv("etcd_ssl"))
	)
	/* 获取命令行参数 */
	flag.StringVar(&envCl, "e", "default", "环境变量默认是default")
	flag.StringVar(&consulAddressCl, "consul_address", "127.0.0.1:8500", "环境变量默认是127.0.0.1:8500")
	flag.Parse()

	common.CommandParameterAdd(
		"env", envCl,
		"consul_address", consulAddressCl,
		"etcd_address", etcdAddressCl,
		"etcd_ssl", etcdSSL,
	)
	/* 初始化配置 */
	if err = configs.InitConfig(); err != nil {
		log.Fatal(err.Error())
	}
	/* 初始化日志 */
	if err = logger.InitLogger(); err != nil {
		log.Fatal(err.Error())
	}
	/* 初始化日志存储组件 */
	if err = logger.InitComponent(); err != nil {
		log.Fatal(err.Error())
	}
	/* 初始化Etcd */
	if err = etcd.GEtcd.InitEtcd(
		etcd.Env(etcd.EnvVar(envCl)),
		etcd.Points([]string{etcdAddressCl}),
		etcd.NeedSSL(etcdSSL),
		etcd.DirPath("./configs/"+envCl+"/k8s_keys/"),
	); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(configs.GConfig.Server.AppName, configs.GConfig.Server.Port)
	/*注册服务到etcd*/
	etcdRegister := etcd.NewServiceRegister(etcd.GEtcd.GetCli())
	if err = etcd.ServiceRegisterHandler.InitServiceRegister(
		context.Background(),
		etcdRegister,
		configs.GConfig.Server.AppName,
		configs.GConfig.Server.Port); err != nil {
		log.Fatal(err.Error())
	}
	/* 初始化Mysql */
	//if err = dbs.InitMysql(); err != nil {
	//	log.Fatal(err.Error())
	//}
	/* 初始化redis */
	//if err = redis_services.InitRedis(); err != nil {
	//	logger.GLogger.Fatal(err.Error())
	//}
	/* 初始化服务器 */
	router.InitRouter()
}
