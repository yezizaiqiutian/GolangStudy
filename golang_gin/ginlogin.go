package main

import (
	"GolangStudy/golang_db"
	"GolangStudy/golang_gin/bean"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

var Trans ut.Translator

func main() {
	golang_db.DBinit()

	r := gin.Default()
	user := r.Group("user")
	{
		user.POST("/login", login)
		user.POST("/regist", regist)
	}
	news := r.Group("news")
	{
		news.GET("/list", newslist)
	}
	_ = InitTrans("zh")
	r.Run(":8888")

}

/**
新闻列表
*/
func newslist(context *gin.Context) {
	var newslist []bean.News

	//var newslist []bean.News = make([]int, 0)
	golang_db.DB.Find(&newslist)
	context.JSON(http.StatusOK, bean.Result{Msg: "成功", Code: "200", Data: newslist})
}

/**
注册
*/
func regist(context *gin.Context) {
	//username := context.PostForm("username")
	//password := context.PostForm("password")
	user := bean.User{}
	err := context.ShouldBind(&user)
	if err != nil {
		HandleValidatorError(context, err)
		//context.JSON(http.StatusOK, bean.Result{Msg: err.Error(), Code: "201"})
		return
	}
	//查询是否已有用户
	var userDB bean.User
	result := golang_db.DB.Where(bean.User{UserName: user.UserName}).Find(&userDB)
	if result.Error != nil {
		context.JSON(http.StatusOK, bean.Result{Msg: result.Error.Error(), Code: "201"})
		return
	}
	if result.RowsAffected != 0 {
		context.JSON(http.StatusOK, bean.Result{Msg: "用户已注册", Code: "201"})
		return
	}

	golang_db.DB.Create(&user)

	context.JSON(http.StatusOK, bean.Result{Msg: "成功", Code: "200", Data: user})

}

/**
登录
*/
func login(context *gin.Context) {
	user := bean.User{}
	err := context.ShouldBind(&user)
	if err != nil {
		HandleValidatorError(context, err)
		//context.JSON(http.StatusOK, bean.Result{Msg: err.Error(), Code: "201"})
		return
	}
	var userDB bean.User
	result := golang_db.DB.Where(bean.User{UserName: user.UserName}).Find(&userDB)
	if result.Error != nil {
		context.JSON(http.StatusOK, bean.Result{Msg: result.Error.Error(), Code: "201"})
		return
	}
	if result.RowsAffected == 0 {
		context.JSON(http.StatusOK, bean.Result{Msg: "用户未注册", Code: "201"})
		return
	}
	//password := context.PostForm("password")

	if context.PostForm("pass_word") != userDB.PassWord {
		context.JSON(http.StatusOK, bean.Result{Msg: "密码错误", Code: "201"})
		return
	}
	context.JSON(http.StatusOK, bean.Result{Msg: "成功", Code: "200", Data: userDB})

}

/**
校验错误
*/
func HandleValidatorError(context *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		context.JSON(http.StatusOK, bean.Result{Msg: err.Error(), Code: "201"})
	}
	context.JSON(http.StatusOK, bean.Result{Msg: removeTopStruct(errs.Translate(Trans)), Code: "201"})
	return

}

/**
校验错误文本拼接
*/
func removeTopStruct(fileds map[string]string) string {
	var rsp string
	for _, err := range fileds {
		rsp += err
	}
	return rsp
}

func InitTrans(locale string) (err error) {
	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, Trans)
		default:
			en_translations.RegisterDefaultTranslations(v, Trans)
		}
		return
	}

	return
}
