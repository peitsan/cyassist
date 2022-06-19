package userDatabase

import (
	"checkinWebsite/database"
	"errors"
	"log"
)

type UserModel struct {
	mail   string //记录对应数据库mail，不可更改
	Openid string `json:"openid" db:"openid"` //微信唯一认证
	Name   string `json:"name" db:"name" `    //姓名
	Xh     string `json:"xh" db:"xh" `        //学号
	Sex    string `json:"sex" db:"sex" `      //性别
}

func GetCheckinModel(mail string) (*UserModel, error) {
	model := new(UserModel)
	err := database.Db.Get(model, "SELECT openid,name,xh,sex FROM user WHERE mail=? LIMIT 1", mail)
	if err != nil {
		log.Printf("获取用户%v的checkinModel错误err:%v", mail, err)
		return nil, errors.New("用户不存在")
	}
	model.mail = mail
	return model, nil
}
