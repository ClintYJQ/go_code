package dao

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/ClintYJQ/name_list_proj/entity"
)

type StudentDaoImpl struct {
}

var (
	ctx = context.Background()
)

// 插入新的数据
func (s *StudentDaoImpl) InsertNewData(sage, sex int32, sname, snum string) (err error) {
	student := entity.Student{
		Snum:  snum,
		Sage:  sage,
		Sname: sname,
		Sex:   sex,
	}
	// 更新数据库
	stat, err := mysqlDB.Prepare("INSERT INTO name_list(sage,sex,sname,snum) VALUES(?,?,?,?)")
	tx, err := mysqlDB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer stat.Close()
	_, err = stat.Exec(sage, sex, sname, snum)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
		return err
	}
	// 更新redis
	bytes, _ := json.Marshal(student)
	_ = rdb.Set(ctx, student.Snum, bytes, 0).Err()
	return nil
}
