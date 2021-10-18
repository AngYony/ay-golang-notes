package forms

type PassWordLoginForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mymobile"` // 手机号码格式验证需要自定义
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	// 图形验证码
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mymobile"` // 手机号码格式验证需要自定义
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	// 短信验证码
	Code string `form:"code" json:"code" binding:"required,min=6,max=6"`
}
