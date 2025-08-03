package controller

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var trans ut.Translator

func InitTrans(local string) (err error) {
	// 这里可以初始化翻译
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 创建中文翻译器
		enT := en.New() // 创建英文翻译器

		//第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(enT, zhT, enT)
		uni := ut.New(enT, zhT, enT)

		// local 通常取决于 http 请求的 Accept-Language 头部
		var ok bool
		// 也可以使用 uni.findTranslator(local) 传入多个local进行查找
		trans, ok = uni.GetTranslator(local)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", local)
		}

		// 注册翻译器
		switch local {
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func removeTopStruct(fields map[string]string) map[string]string {
	// 去掉结构体的顶层字段名
	res := make(map[string]string, len(fields))
	for k, v := range fields {
		// 去掉结构体的顶层字段名
		res[k[strings.Index(k, ".")+1:]] = v
	}
	return res
}
