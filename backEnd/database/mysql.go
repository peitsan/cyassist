package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type RegisterInfo struct {
	//binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Mail     string `form:"mail" json:"mail" uri:"mail" xml:"mail" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

type BindInfo struct {
	Cqupt_username string `form:"cqupt_username" json:"cqupt_username" uri:"cqupt_username" xml:"cqupt_username" binding:"required"`
	Cqupt_password string `form:"cqupt_password" json:"cqupt_password" uri:"cqupt_password" xml:"cqupt_password" binding:"required"`
	Openid         string `form:"openid" json:"openid" uri:"openid" xml:"openid" `
	Mail           string `form:"mail" json:"mail" uri:"mail" xml:"mail" binding:"required"`
}

type LoginInfo struct {
	Mail     string `form:"mail" json:"mail" uri:"mail" xml:"mail" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

var Db *sqlx.DB

func InitDatabase(user, password, host, port, dbname string) {
	//开启了时间解析功能
	Db = sqlx.MustConnect("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", user, password, host, port, dbname))
	return
}

// RegisterToDatabase 用户注册数据注入数据库
func RegisterToDatabase(mail, password string) error {
	sql := "insert into user(mail,password)values (?,?)"

	_, err := Db.Exec(sql, mail, password)

	if err != nil {
		log.Printf("注册时，注册信息注入数据库发生错误，err:", err.Error())
		return err
	} else {
		return nil
	}
}

// CquptBindToDatabase 用户注册数据注入数据库
func CquptBindToDatabase(cqupt_user_id, cqupt_password, openid, mail, name, xh, sex string) (ok bool, code int, message string) {
	sql := "UPDATE user SET cqupt_user_id=?,cqupt_password=?,openid=?,name=?,xh=?,sex=? WHERE mail=?"

	result, err := Db.Exec(sql, cqupt_user_id, cqupt_password, openid, name, xh, sex, mail)
	if err != nil {
		log.Printf("绑定cqupt阶段，数据库插入后，获取返回result时发生err:%v", err.Error())
		return false, 500, err.Error()
	}
	affectedRows, err := result.RowsAffected()

	if err != nil {
		log.Printf("绑定cqupt阶段，数据库插入后，获取返回生效的affectRows时发生err:%v", err.Error())
		return false, 500, err.Error()
	} else if affectedRows == 0 {
		return false, 2400, "该邮箱未被注册"
	} else {
		return true, 2200, "success"
	}

}

func CheckMailExist(mail string) (bool, int) {
	sql := "select count(*) from user where mail = ?"
	result, err := Db.Query(sql, mail)

	if result == nil {
		log.Printf("注册查询数据库时错误,查询邮箱:%v,error:%v", mail, err)
	}

	defer result.Close()

	if result.Next() {
		var count int
		err := result.Scan(&count)
		if err != nil {
			log.Printf("注册查询数据库是否已存在邮箱时错误,查询邮箱:%v,error:%v", mail, err)
			return false, 500
		}
		if count > 0 {
			return true, 2200
		} else {
			return false, 2200
		}
	} else {
		log.Printf("注册查询数据库是否已存在邮箱时错误,查询邮箱:%v,error:%v", mail, "数据库未返回任何数据")
		return false, 500
	}
}

// RegisterCheckMailRepeat 检查用户注册是否已存在邮箱
func RegisterCheckMailRepeat(mail string) (ok bool, code int, message string) {
	ok, code = CheckMailExist(mail)
	if ok == true && code == 2200 {
		return false, 2400, "邮箱已存在"
	} else if code == 2200 {
		return true, 2200, ""

	} else {
		return false, code, ""
	}

}

func Login(user, password string) (code int, message string) {
	sql := "select password from user where mail = ?"
	result, err := Db.Query(sql, user)
	if err != nil {
		return 500, "查询数据库失败"
	}
	if result.Next() {
		var rPassword string
		err := result.Scan(&rPassword)
		if err != nil {
			log.Printf("登录查询数据库是否已存在邮箱时错误,查询邮箱:%v,error:%v", user, err)
			return 500, "查询数据库失败"
		}
		if rPassword == password {
			return 2200, "success"
		} else {
			return 2401, "邮箱或密码错误"
		}
	} else {
		return 2401, "不存在该邮箱"
	}
}
