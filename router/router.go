package router

import (
	"context"
	"convert.api/controllers"
	e "convert.api/libs/common/error_wrapper"
	"convert.api/libs/configs"
	"convert.api/router/middlewares/trace_id"
	"github.com/gin-gonic/gin"
	"github.com/outreach-golang/logger"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func InitRouter() {
	gin.SetMode(configs.GConfig.Mode)

	if configs.GConfig.Mode != "debug" {
		gin.DefaultWriter = ioutil.Discard
	}

	router := gin.Default()
	router.Use(
		trace_id.SetUp(),
	)

	router.GET("/ping", func(c *gin.Context) { c.String(200, "PONG") })
	api := router.Group("/api")
	{
		//保单管理列表
		api.POST("/convertPdf", e.WrapperError(controllers.ConvertPdf))
	}
	srv := &http.Server{
		Addr:    configs.GConfig.Server.Address + ":" + configs.GConfig.Server.Port,
		Handler: router,
	}
	go func() {
		// 服务连接
		logger.GLogger.Info("服务器启动成功!")
		f, err := os.OpenFile("pid.txt", os.O_CREATE|os.O_RDWR, 0666)
		if err == nil {
			_, err = f.WriteString(strconv.Itoa(syscall.Getpid()))
			if err != nil {
				defer f.Close()
			}
		} else {
			logger.GLogger.Fatal(err.Error())
		}

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.GLogger.Fatal(err.Error())
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.GLogger.Info("服务器关闭中 ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.GLogger.Fatal(err.Error())
	}
	logger.GLogger.Info("服务器关闭！")

}
