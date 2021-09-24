package forms

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mymobile"` // 手机号码格式验证需要自定义
	Type   uint   `form:"type" json:"type" binding:"required,oneof=1 2"`    // oneof限定值只能取 1 或者 2

}
