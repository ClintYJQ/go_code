package dao

type StudentDao interface {
	//增删改查
	InsertNewData(sage, sex int32, sname, snum string) error
	//删除记录
	//deleteData() (n int32, err error)
	//修改记录
	//updateDataById(id int32) error
	//查找记录
	//selectDataById(id int32) error
}

func GetStuDao() StudentDao {
	return &StudentDaoImpl{}
}
