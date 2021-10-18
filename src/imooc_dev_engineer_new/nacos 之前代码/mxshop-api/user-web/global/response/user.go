package response

import (
	"fmt"
	"time"
)

type UserResponse struct {
	Id              int32     `json:"id"`
	NickName        string    `json:"name"`
	Birthday        time.Time `json:"birthday"`
	BirthdayString1 JsonTime  `json:"birthday_string1"` // 格式化时间格式字符串（方式一）
	BirthdayString2 string    `json:"birthday_string2"` // 格式化时间格式字符串（方式二）
	Gender          string    `json:"gender"`
	Mobile          string    `json:"mobile"`
}

/*
 如何格式化时间格式的字符串
*/
type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-01"))
	return []byte(stmp), nil
}
