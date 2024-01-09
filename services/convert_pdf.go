package services

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type convertPdf struct {
	ctx context.Context
}

func NewConvertPdf(ctx context.Context) *convertPdf {
	return &convertPdf{
		ctx: ctx,
	}
}

//1下载文件
//2、下载的文件转pdf
//3、转化后的pdf上传oss

// ConvertToPdf 文件转pdf
func (c *convertPdf) ConvertToPdf(reqUrl, fileType string) (res string, err error) {

	// 对url中的 ASCII字符集 转义
	switchUrl, err := url.QueryUnescape(reqUrl)
	if err != nil {
		return "", fmt.Errorf("url转义失败" + err.Error())
	}
	//var deleteFiles []string
	osName := runtime.GOOS //获取系统类型
	nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	outPath := "./configs/" + nanoStr
	pdfFile := ""
	//获取下载文件的文件名称
	getFileName := path.Base(switchUrl)
	if err != nil {
		return "", fmt.Errorf("获取下载文件名称失败，" + err.Error())
	}
	// 1、下载文件
	fileSrcPath, err := c.FileDown(switchUrl, "./configs/"+nanoStr, getFileName)
	if err != nil {
		return "", fmt.Errorf("文件下载异常，" + err.Error())
	}
	// 删除文件下的文件
	defer func() {
		_ = os.RemoveAll(outPath)
	}()

	//2、下载的文件转pdf
	switch osName {
	case "darwin": //mac系统
		command := "/Applications/LibreOffice.app/Contents/MacOS/soffice"
		pdfFile, err = c.FuncToPdf(command, fileSrcPath, outPath, fileType)
		if err != nil {
			return "", fmt.Errorf("文件转化异常，" + err.Error())
		}
	case "linux":
		command := "libreoffice7.3"
		pdfFile, err = c.FuncToPdf(command, fileSrcPath, outPath, fileType)
		if err != nil {
			return "", fmt.Errorf("文件转化异常，" + err.Error())
		}
	default:
		return "", fmt.Errorf("暂时不支持的系统转化:" + runtime.GOOS)
	}
	//3、转化后的pdf上传oss
	filename := path.Base(pdfFile)
	ossUrl := ""
	file := outPath+"/"+filename
	fileName := filename
	filePath := "pdf/" + nanoStr
	ossUrl,err  = NewConvertOss(c.ctx).ConvertToOss(file,fileName,filePath)
	if err != nil {
		return  "",err
	}
	return ossUrl, nil
}

// mac 成功
//

/**
*  一定要注意：使用前mac和服务器均需先安装：LibreOffice
*
*@tips libreoffice 转换指令：
* libreoffice6.2 invisible --convert-to pdf csDoc.doc --outdir /home/[转出目录]
* mac指令：
* /Applications/LibreOffice.app/Contents/MacOS/soffice --convert-to pdf /Users/qinleixing/Desktop/123/test2.docx --outdir /Users/qinleixing/Desktop/123/
*
* @function 实现文档类型转换为pdf或html
* @param command:libreofficed的命令(具体以版本为准)；win：soffice； linux：libreoffice6.2
*     fileSrcPath:转换文件的路径
*     fileOutDir:转换后文件存储目录
*     converterType：转换的类型pdf/html
* @return fileOutPath 转换成功生成的文件的路径 error 转换错误
 */

// FuncToPdf 文件转化pdf
func (c *convertPdf) FuncToPdf(command string, fileSrcPath string, fileOutDir string, converterType string) (fileOutPath string, error error) {
	//校验fileSrcPath
	srcFile, erByOpenSrcFile := os.Open(fileSrcPath)
	if erByOpenSrcFile != nil && os.IsNotExist(erByOpenSrcFile) {
		return "", erByOpenSrcFile
	}
	//如文件输出目录fileOutDir不存在则自动创建
	outFileDir, erByOpenFileOutDir := os.Open(fileOutDir)
	if erByOpenFileOutDir != nil && os.IsNotExist(erByOpenFileOutDir) {
		erByCreateFileOutDir := os.MkdirAll(fileOutDir, os.ModePerm)
		if erByCreateFileOutDir != nil {
			//fmt.Println("File ouput dir create error.....", erByCreateFileOutDir.Error())
			return "", erByCreateFileOutDir
		}
	}
	//关闭流
	defer func() {
		_ = srcFile.Close()
		_ = outFileDir.Close()
	}()
	//convert
	cmd := exec.Command(command, "--invisible", "--language=zh-CN", "--convert-to", converterType,
		fileSrcPath, "--outdir", fileOutDir)
	_, errByCmdStart := cmd.Output()
	//命令调用转换失败
	if errByCmdStart != nil {
		return "", errByCmdStart
	}
	//success
	fileOutPath = fileOutDir + "/" + strings.Split(path.Base(fileSrcPath), ".")[0]
	if converterType == "html" {
		fileOutPath += ".html"
	} else {
		fileOutPath += ".pdf"
	}
	//fmt.Println("文件转换成功...", string(byteByStat))
	return fileOutPath, nil
}

// FileDown 下载文件
func (c *convertPdf) FileDown(downUrl string, savePath, fileName string) (localPath string, err error) {

	// Get the data
	resp, err := http.Get(downUrl)
	defer resp.Body.Close()

	//如文件输出目录fileOutDir不存在则自动创建
	_, erByOpenFilDir := os.Open(savePath)
	if erByOpenFilDir != nil && os.IsNotExist(erByOpenFilDir) {
		erByCreateFileDir := os.MkdirAll(savePath, os.ModePerm)
		if erByCreateFileDir != nil {
			return "", fmt.Errorf("创建文件夹失败，" + erByCreateFileDir.Error())
		}
	}
	pathFile := savePath + "/" + fileName
	// 创建一个文件用于保存
	out, err := os.Create(pathFile)
	if err != nil {
		return "", err
	}
	defer out.Close()
	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	return pathFile, nil
}

