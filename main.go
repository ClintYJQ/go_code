package main

import (
	"context"
	"log"

	"github.com/ClintYJQ/name_list_proj/config"
	"github.com/ClintYJQ/name_list_proj/service"
)

var ctx = context.Background()

func main() {
	conf := &config.Config{}
	var err error
	if conf, err = config.LoadConfig(ctx); err != nil {
		log.Fatalf("load config failed,err msg:%v", err)
	}
	log.Printf("load config success,config:&v", conf)
	globalHandler := service.NewGlobalHandler(conf)
	globalHandler.Run()
}
