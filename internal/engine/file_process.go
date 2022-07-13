package engine

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/ClintYJQ/name_list_proj/dao"
	"github.com/ClintYJQ/name_list_proj/utils"
)

// 处理csv文件，解析成需要的数据，并写入数据库和redis
func ReadFileToDB(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("read local file failed msg: %v", err)
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("read file failed: %v", err)
			return
		}
		snum := utils.GenUniqueID()
		log.Printf("generate an unique student num is %v", snum)
		sname := record[0]
		sage, _ := strconv.Atoi(record[1])
		sex, _ := strconv.Atoi(record[2])
		insertErr := dao.GetStuDao().InsertNewData(int32(sage), int32(sex), sname, snum)
		if insertErr != nil {
			log.Fatalf("insert record failed msg: %v", err)
		}
	}
}
