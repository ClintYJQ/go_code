package service

import (
	"context"
	"log"

	"github.com/ClintYJQ/name_list_proj/config"
	"github.com/ClintYJQ/name_list_proj/dao"
	"github.com/gin-gonic/gin"
)

type GlobalHandlerImpl struct {
	ctx    context.Context
	router *gin.Engine
	conf   *config.Config
}

// 添加路由
func (g GlobalHandlerImpl) setupRouter() {
	//学生数据操作组
	//stuRoute := g.router.Group("/stu")

	//文件操作组
	fileRoute := g.router.Group("/file")
	addFileRoute(fileRoute)
}

func (g GlobalHandlerImpl) Run() {
	if err := dao.InitClient(g.conf); err != nil {
		log.Fatalf("init mysql client failed,err msg:%+v", err)
		panic(err)
	}
	g.setupRouter()
	g.router.Run()
}
