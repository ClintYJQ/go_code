package entity

// 学生数据
type Student struct {
	Snum  string `json:"snum"`
	Sname string `json:"sname"`
	Sage  int32  `json:"sage"`
	Sex   int32  `json:"sex"`
}
