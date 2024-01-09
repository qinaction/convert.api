package logger

import (
	"convert.api/libs/common"
	"convert.api/libs/configs"
	"github.com/outreach-golang/logger"
	"go.uber.org/zap"
)

var (
	GLogger         *zap.Logger
	GSentinelLogger *SentinelLogger
)

type SentinelLogger struct {
	GLogger *zap.Logger
}

func (s *SentinelLogger) Debug(msg string, keysAndValues ...interface{}) {
	s.GLogger.Debug(msg, zap.Any("sentinel", keysAndValues))
}

func (s *SentinelLogger) DebugEnabled() bool {
	return true
}

func (s *SentinelLogger) Info(msg string, keysAndValues ...interface{}) {
	s.GLogger.Info(msg, zap.Any("sentinel", keysAndValues))
}

func (s *SentinelLogger) InfoEnabled() bool {
	return true
}

func (s *SentinelLogger) Warn(msg string, keysAndValues ...interface{}) {
	s.GLogger.Warn(msg, zap.Any("sentinel", keysAndValues))
}

func (s *SentinelLogger) WarnEnabled() bool {
	return true
}

func (s *SentinelLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	s.GLogger.Error(err.Error(), zap.String("sentinel-error", msg), zap.Any("sentinel", keysAndValues))
}

func (s *SentinelLogger) ErrorEnabled() bool {
	return true
}

func InitLogger() error {
	var (
		err  error
		conf = configs.GConfig.Log
	)

	GLogger, err = logger.NewLogger(
		logger.ServerName(configs.GConfig.Server.AppName),
		logger.SaveLogAddr(logger.AliLog),

		logger.AccessKeyID(conf.AccessKeyID),
		logger.AccessKeySecret(conf.AccessKeySecret),
		logger.LogStore(conf.LogStore),
		logger.Endpoint(conf.Endpoint),
		logger.Project(conf.Project),
		logger.Source(common.LocalIP()),
		logger.LooKAddr(conf.LookAddr),

		logger.DingHost(configs.GConfig.Dingtalk.Host),
		logger.DingWebhook(configs.GConfig.Dingtalk.Webhook),
	)

	if err != nil {
		return err
	}

	GSentinelLogger = &SentinelLogger{}
	GSentinelLogger.GLogger = GLogger

	return nil
}
