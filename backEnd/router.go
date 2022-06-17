package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitialRouter(key string) *gin.Engine {
	store := cookie.NewStore([]byte(key)) //设置session加密密钥
	r := gin.Default()
	r.Use(sessions.Sessions("GOSESSION", store))
	r.Use(AllowCorsHandler)
	PostRouter(r)
	OptionRouter(r)
	GetRouter(r)
	return r
}

func GetRouter(r *gin.Engine) {
	userGroup := r.Group("/user", AuthHandler)
	userGroup.GET("/checkin", CheckinIndexHandler)
	userGroup.GET("/checkin/querydatas", QueryDatasHandler)
}

func PostRouter(r *gin.Engine) {
	userGroup := r.Group("/user", AuthHandler)
	r.POST("/register", RegisterHandler)
	userGroup.POST("/bindCqupt", BindCquptHandler)
	r.POST("/login", LoginHandler)
	userGroup.POST("/checkin/submitdatas", SubmitdatasHandler)
	userGroup.POST("/checkin/submitUID", SubmitUIDHandler)
}

func OptionRouter(r *gin.Engine) {
	r.OPTIONS("/register", OptionHandler)
	r.OPTIONS("/user/bindCqupt", OptionHandler)
	r.OPTIONS("/login", OptionHandler)
	r.OPTIONS("/user/checkin/submitdatas", OptionHandler)
	r.OPTIONS("/user/checkin", OptionHandler)
	r.OPTIONS("/user/checkin/querydatas", OptionHandler)
	r.OPTIONS("/user/checkin/submitUID", OptionHandler)
}
