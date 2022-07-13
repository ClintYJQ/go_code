package service

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ClintYJQ/name_list_proj/internal/engine"
	"github.com/gin-gonic/gin"
)

func addFileRoute(rg *gin.RouterGroup) {
	rg.POST("/upload", func(ctx *gin.Context) {
		//使用curl自测接口
		form, _ := ctx.MultipartForm()
		files := form.File["upload[]"]
		// 用户上传文件暂存位置
		dst := "/home/sheep/go_code_wsl/go_namelist_project/dst/"
		for _, file := range files {
			log.Println(file.Filename)
			savePath := dst + file.Filename
			// 防止重复上传
			_, err := os.Stat(savePath)
			if err == nil {
				log.Println("There is a file with same name!!!!")
				ctx.String(http.StatusOK, fmt.Sprintln("file is exist!!!")) // 替换成json消息
				return
			}
			ctx.SaveUploadedFile(file, savePath)
			log.Println("save file to " + savePath)
			// 替换成csv解析方法, 让后调用dao层服务，讲数据持久化
			engine.ReadFileToDB(savePath)
		}
		ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded ||records saved succes\n", len(files))) // 替换成json消息
	})
}
