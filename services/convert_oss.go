package services

import (
	"bytes"
	"context"
	"convert.api/libs/common/error_wrapper"
	"convert.api/libs/configs"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type convertOss struct {
	ctx context.Context
}

type OssRes struct {
	Errno    string `json:"errno" `
	Error    string `json:"error"`
	DataType string `json:"dataType"`
	Data     string `json:"data"`
}

func NewConvertOss(ctx context.Context) *convertOss {
	return &convertOss{
		ctx: ctx,
	}
}

// ConvertToOss 文件转oss
func (c *convertOss) ConvertToOss(file, fileName, ossFilePath string) (res string, err error) {

	postData := make(map[string]string)
	postData["fileName"] = fileName
	postData["filePath"] = ossFilePath
	hostUrl := configs.GConfig.Host.OssUpload
	hostUrl += "/api/oss/uploadByByte"
	oss, err := c.PostFile(file, hostUrl, postData)
	if err != nil {
		return "", err
	}
	return oss, nil
}

// PostFile 文件上传
func (c *convertOss) PostFile(filename string, targetUrl string, params map[string]string) (ossUrl string, err error) {

	var res OssRes
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return "", err
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file" + err.Error())
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return "", err
	}
	//添加其他参数
	if len(params) != 0 {
		//param是一个一维的map结构
		for k, v := range params {
			bodyWriter.WriteField(k, v)
		}
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(respBody, &res)
	if err != nil {
		return "", err
	}
	if res.Errno != "0" {
		return "", error_wrapper.WitheError(res.Error)
	}
	return res.Data, nil
}
