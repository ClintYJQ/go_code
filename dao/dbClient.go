package dao

import (
	"database/sql"
	"log"

	"github.com/ClintYJQ/name_list_proj/config"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlDB *sql.DB
	rdb     *redis.Client
)

func InitClient(conf *config.Config) error {
	var err error
	mysqlDB, err = sql.Open("mysql", conf.MySQLConf.URL)
	if err != nil {
		log.Printf("init mysql client failed,err msg:%+v", err)
		return err
	}
	log.Printf("init mysql client success!")
	stat, err := mysqlDB.Prepare("SELECT * FROM name_list")
	if err != nil {
		log.Fatalf("DB failed %v", err)
	}
	stat.Exec()
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.RedisConf.Addr,
		Password: conf.RedisConf.PassWd,
		DB:       conf.RedisConf.DB,
	})
	return nil
}
