package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

func main() {

	println(runtime.GOMAXPROCS(runtime.NumCPU()))
	engine := gin.Default()

	engine.GET("/demo-service/_get", func(context *gin.Context) {

		context.JSON(200, gin.H{
			"code": "00",
			"msg":  "GET OK",
			"data": true,
		})
	})

	engine.POST("/demo-service/_post", func(context *gin.Context) {
		data, _ := context.GetRawData()
		context.JSON(200, gin.H{
			"code": "00",
			"msg":  "POST OK",
			"data": string(data),
		})
	})

	engine.MaxMultipartMemory = 8 << 20 // 8 MiB
	engine.POST("/demo-service/_upload", func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		basePath := "/download/files/temp/"
		filename := basePath + filepath.Base(file.Filename)
		if err := context.SaveUploadedFile(file, filename); err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		context.String(http.StatusOK, fmt.Sprintf("文件 %s 上传成功 ", file.Filename))
	})

	engine.GET("/demo-service/_download", func(context *gin.Context) {
		context.File("/download/files/temp/" + context.GetString("file"))
	})

	err := engine.Run(":7777")

	if err != nil {
		log.Writer()
		log.Fatalln("server stop", err)
	}
}
