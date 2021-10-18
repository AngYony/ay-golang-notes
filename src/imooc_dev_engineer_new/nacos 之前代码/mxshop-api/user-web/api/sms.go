package api

import (
	"context"
	"fmt"
	"math/rand"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/go-redis/redis/v8"
)

// GenerateSmsCode 生成width长度的短信验证码
func GenerateSmsCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano()) // 以纳秒作为因子

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// SendSms 通过阿里的短信服务发送短信验证码到手机端
func SendSms(ctx *gin.Context) {

	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBind(&sendSmsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	smsCode := GenerateSmsCode(6)
	// mobile := "18301219999"

	// client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", "xxxx", "xxx")
	// if err != nil {
	// 	panic(err)
	// }
	// request := requests.NewCommonRequest()
	// request.Method = "POST"
	// request.Scheme = "https" // https | http
	// request.Domain = "dysmsapi.aliyuncs.com"
	// request.Version = "2017-05-25"
	// request.ApiName = "SendSms"
	// request.QueryParams["RegionId"] = "cn-beijing"
	// request.QueryParams["PhoneNumbers"] = "接收者电话号码"                                // 手机号
	// request.QueryParams["SignName"] = "创建的签名"                                      // 阿里云验证过的签名 自己设置
	// request.QueryParams["TemplateCode"] = "xxx"                                    // 阿里云的短信模板号 自己设置
	// request.QueryParams["TemplateParam"] = "{\"code\":" + GenerateSmsCode(6) + "}" // 短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
	// response, err := client.ProcessCommonRequest(request)
	// fmt.Print(client.DoAction(request, response))
	// //  fmt.Print(response)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }

	// 将验证码保存起来
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	// 并设置过期时间
	rdb.Set(context.Background(), sendSmsForm.Mobile, smsCode, time.Second*time.Duration(global.ServerConfig.RedisInfo.Expire))

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})
	// json数据解析

}
