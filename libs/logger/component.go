package logger

import (
	"convert.api/libs/configs"
	sls "github.com/aliyun/aliyun-log-go-sdk"
)

var (
	GLoggerComponentClient sls.ClientInterface
)

func InitComponent() error {
	GLoggerComponentClient = sls.CreateNormalInterface(configs.GConfig.Log.Endpoint, configs.GConfig.Log.AccessKeyID, configs.GConfig.Log.AccessKeySecret, "")

	return nil
}
