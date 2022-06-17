package main

import (
	"checkinWebsite/app/checkin"
	"checkinWebsite/app/login"
	"checkinWebsite/database"
	"checkinWebsite/database/checkinDatabase"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type BackJsonStruct struct {
	Code    int
	Message string
}

//把传入code，message，输出成json的string形式的{"code": 200,"message"="xxx"}的形式
func toJsonString(code int, message string) string {
	data, _ := json.Marshal(BackJsonStruct{code, message})
	return string(data)
}

//对传入code解析，将2000以上的code，取后三位后的code放置调用toJsonString中code的参数，backCode置200.用于告诉前端是他的问题
//小于2000的则直接backCode置输入code，message直接返回。用于告诉调用方是后端错误
func backFormat(code int, message string) (backCode int, payload string) {
	if code > 2000 {
		return 200, toJsonString(code-2000, message)
	} else {
		return code, message
	}
}

// AuthHandler 认证中间件，验证时间、用户
func AuthHandler(c *gin.Context) {
	for key, value := range c.Request.Cookies() {
		fmt.Printf("%v,%v\n", key, value)
	}
	for key, value := range c.Request.Header {
		fmt.Printf("%v,%v", key, value)
	}
	ok := login.Auth(c)
	if !ok {
		c.String(backFormat(2401, "认证失败"))
		c.Abort()
		return
	}
}

// RegisterHandler 注册
func RegisterHandler(c *gin.Context) {
	var form database.RegisterInfo
	if err := c.ShouldBind(&form); err != nil {
		log.Println("注册时转换Bind失败")
		c.String(backFormat(2400, "传入数据错误"))
		return
	}
	code, message := login.WebRegister(&form)
	if code == 2200 {
		token, err := login.GetLoginToken(form.Mail)
		if err != nil {
			c.String(500, "加解密错误")
			return
		}
		responseJson := &struct {
			Code    int    `json:"code"`
			Token   string `json:"token"`
			Message string `json:"message"`
		}{200, token, message}
		backString, _ := json.Marshal(responseJson)
		c.String(200, string(backString))
		return
	}

	c.String(backFormat(code, message))
}

//绑定openid
func BindCquptHandler(c *gin.Context) {
	var form database.BindInfo
	if err := c.ShouldBind(&form); err != nil {
		log.Println("绑定转换时转换Bind失败", c.Request.Body)
		c.String(backFormat(2400, "传入数据错误"))
		return
	}
	code, message := login.CquptBind(&form)
	c.String(backFormat(code, message))
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	var form database.LoginInfo
	if err := c.ShouldBind(&form); err != nil {
		log.Printf("登录转换时转换Bind失败err:%v", err.Error())
		c.String(backFormat(2400, "传入数据错误"))
		return
	}
	code, message, token := login.Login(form.Mail, form.Password, c)

	if code == 2200 {
		responseJson := &struct {
			Code    int    `json:"code"`
			Token   string `json:"token"`
			Message string `json:"message"`
		}{200, token, message}
		backString, _ := json.Marshal(responseJson)
		c.String(200, string(backString))
		return
	}

	c.String(backFormat(code, message))
}

// QueryDatasHandler 查询用户打卡录入的信息
func QueryDatasHandler(c *gin.Context) {
	value, ok := c.Get("mail")
	if !ok {
		log.Printf("在QueryDatasHandler中访问上下文mail时读取失败")
		c.String(500, "服务端获取上下文失败")
		return
	} else {
		c.String(checkin.QueryDatas(value.(string)))
	}
}

// SubmitdatasHandler 提交或更新打卡数据
func SubmitdatasHandler(c *gin.Context) {
	value, ok := c.Get("mail")
	if !ok {
		log.Printf("在SubmitdatasHandler中访问上下文mail时读取失败")
		c.String(500, "服务端获取上下文失败")
	}
	var form checkinDatabase.CheckinModel
	if err := c.ShouldBind(&form); err != nil {
		log.Printf("提交打卡数据时转换Bind失败err:%v", err.Error())
		c.String(backFormat(2400, "传入数据错误"))
		return
	}
	if &form.Offset == nil || &form.Ywjchblj == nil || &form.Xjzdywqzbl == nil || &form.Beizhu == nil || &form.Ywytdzz == nil || &form.Twsfzc == nil {
		log.Println("提交打卡数据阶段未传入必须要求字段")
		c.String(backFormat(2400, "传入数据错误"))
		return
	}
	if form.Offset > 864000 {
		log.Printf("提交打卡数据阶段发现过长offset:%v,hack！", form.Offset)
		c.String(backFormat(2400, "传入数据错误"))
		return
	}
	form.Mail = value.(string)
	c.String(backFormat(checkin.SubmitDatas(form)))
}

func CheckinIndexHandler(c *gin.Context) {
	value, ok := c.Get("mail")
	if !ok {
		log.Printf("在CheckIndexHandler中访问上下文mail时读取失败")
		c.String(500, "服务端获取上下文失败")
	} else {
		code, message, data := checkin.CheckinIndex(value.(string))
		if code > 2000 {
			response := checkinDatabase.CheckinLogResponse{code - 2000, message, data}
			marshal, err := json.Marshal(response)
			if err != nil {
				log.Printf("在CheckIndexHandler中marshal json 失败")
				c.String(500, "解析数据失败")
				return
			}
			c.String(200, string(marshal))
		} else {
			c.String(code, message)
		}
	}
}

func SubmitUIDHandler(c *gin.Context) {
	value, ok := c.Get("mail")
	if !ok {
		log.Printf("在SubmitUIDHandler中访问上下文mail时读取失败")
		c.String(500, "服务端获取上下文失败")
		return
	}

	postJson := &struct {
		UID string `json:"UID" form:"UID" binding:"required"`
	}{}

	if err := c.ShouldBind(postJson); err != nil {
		log.Printf("在SubmitUIDHandler中marshal json 失败,err:%v", err)
		c.String(500, "解析数据失败")
		return
	}

	c.String(backFormat(checkin.SubmitUID(value.(string), postJson.UID)))
	return
}

func AllowCorsHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

}

func OptionHandler(c *gin.Context) {
	/*for k, v := range c.Request.Header {
		fmt.Printf("%v-%v\n", k, v)
	}*/
	c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "authorization,content-type")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
}
