package api

import (
	"TodoList/serializer"
	"encoding/json"
	"fmt"
)

// 将错误序列化一下
func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "JSON 类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Status: 40001,
		Msg:    "错误",
		Error:  fmt.Sprint(err),
	}

}
