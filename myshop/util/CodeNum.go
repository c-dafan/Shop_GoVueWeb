package util

type CodeNum struct {
	Code int32
	Msg string
	Data interface{}
}

func Success(data interface{}) CodeNum {
	return CodeNum{Code:200, Msg:"成功", Data:data}
}

func Error(err error) CodeNum {
	return CodeNum{Code:100, Msg:err.Error(), Data:nil}
}