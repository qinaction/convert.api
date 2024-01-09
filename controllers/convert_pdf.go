package controllers

import (
	"convert.api/libs/common/error_wrapper"
	"convert.api/models/request"
	"convert.api/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var ch = make(chan int, 1)

func ConvertPdf(c *gin.Context) error {
	var (
		req request.ConvertPdf
		err error
	)
	if err = c.ShouldBindWith(&req, binding.JSON); err != nil {
		return error_wrapper.WitheError(err.Error())
	}
	ch <- 1
	url, err := services.NewConvertPdf(c).ConvertToPdf(req.Url, req.FileType)
	if err != nil {
		<-ch
		return error_wrapper.WitheError("转化文件失败:" + err.Error())
	}
	<-ch
	return error_wrapper.WithSingle(url)
}
