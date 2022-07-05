package utils

import (
	"Oenone/common/base"
	"github.com/cloopen/go-sms-sdk/cloopen"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
	"log"
	"strconv"
)

func SendCaptcha(telephone string) (int, error) {
	captchaClient := cloopen.NewJsonClient((base.GLOBAL_RESOURCE[base.CaptchaClientConfig]).(*cloopen.Config)).SMS()
	// 随机4位字符
	captchaCode := funk.RandomInt(1000, 9999)

	// 发送短信
	input := &cloopen.SendRequest{
		// 应用的APPID
		AppId: viper.GetString("messageInfo.AppID"),
		// 手机号码
		To: telephone,
		// 模版ID
		TemplateId: viper.GetString("messageInfo.TemplateID"),
		// 模版变量内容 非必
		Datas: []string{strconv.Itoa(captchaCode), "5"},
	}

	_, err := captchaClient.Send(input)

	if err != nil {
		log.Println(err)
		return 0, err
	}
	return captchaCode, nil
}
