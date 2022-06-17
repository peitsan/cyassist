package login

import (
	"checkinWebsite/database"
	"checkinWebsite/joshuaRequest"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func WebRegister(info *database.RegisterInfo) (int, string) {
	ok, code, message := checkMail(info.Mail)
	if !ok {
		return code, message
	}
	ok, code, message = checkPassword(info.Password)
	if !ok {
		return code, message
	}
	if ok {
		err := database.RegisterToDatabase(info.Mail, info.Password)
		if err != nil {
			return 500, "数据库错误"
		}

	}
	return 2200, "success"
}

func CquptBind(info *database.BindInfo) (code int, message string) {
	//检查邮箱是否存在
	ok, code := database.CheckMailExist(info.Mail)
	if code != 2200 {
		return code, "数据库错误"
	} else if !ok && code == 2200 {
		return 2400, "邮箱不存在"
	}

	return bindCqupt(info.Openid, info.Cqupt_username, info.Cqupt_password, info.Mail)
}

//检查密码是不是太短了或是太长了
func checkPassword(password string) (ok bool, code int, message string) {
	if len(password) > 20 {
		return false, 2400, "密码过长"
	}
	if len(password) < 6 {
		return false, 2400, "密码过短"
	}
	return true, 2200, ""
}

//检查输入是否是邮箱，以及数据库中查询邮箱是否已经被使用
func checkMail(mail string) (ok bool, code int, message string) {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(mail) {
		return false, 2400, "邮箱格式错误"
	}
	return database.RegisterCheckMailRepeat(mail)

}

//openId长度为28，前6位固定，后几位似乎是随机
func randomOpenid() string {
	fixHead := []byte("oIaII0") // we重邮小程序openid固定开头
	set := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXY0123456789_-"
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := strings.Builder{}
	result.Write(fixHead)

	for i := 0; i < 22; i++ {
		result.WriteByte(set[rand1.Intn(62)])
	}
	return result.String()
}

//与we重邮服务器进行绑定
func bindCqupt(sOpenid, username, password, mail string) (code int, message string) {
	post_url := "https://we.cqupt.edu.cn/api/users/bind.php"
	headers := map[string][]string{
		"Host":         {"we.cqupt.edu.cn"},
		"Connection":   {"keep-alive"},
		"User-Agent":   {"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat"},
		"content-type": {"application/json"},
	}

	type data struct {
		Name string `json:"name"`
		Xh   string `json:"xh"`
		Sex  string `json:"sex"`
	}

	type responseJsonStruct struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    data   `json:"data"`
	}

	var openid string

	if sOpenid == "" {
		openid = randomOpenid()
	} else {
		if len(sOpenid) != 28 {
			return 2400, "传入openid非法"
		}
		openid = sOpenid
	}

	for {
		information_dict := map[string]string{
			"openid":    openid,
			"yktid":     username,
			"passwd":    password,
			"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
		}
		data, err := json.Marshal(information_dict)
		payload, err := json.Marshal(map[string]string{"key": base64.StdEncoding.EncodeToString(data)})

		options := []joshuaRequest.FuncRequestOption{
			joshuaRequest.SetHeader(headers),
			joshuaRequest.SetData(string(payload)),
		}
		response, err := joshuaRequest.Request("POST", post_url, options...)

		if err != nil {
			log.Printf("请求绑定cqupt时错误,用户:%v,密码:%v,err:%v", username, password, err.Error())
			return 2504, "与重邮服务器连接失败"
		}

		if response.StatusCode == 200 {
			responseByte, err2 := io.ReadAll(response.Body)

			//openid重复，重新构造进行请求。 该返回不像其它错误返回标准json，而是直接返回字符串，如Duplicate entry 'oIaII0cEkAvDOGJC0KfjkdT-sT-D-2011213355' for key 'PRIMARY'
			if strings.Contains(string(responseByte), "Duplicate") {
				if sOpenid != "" {
					return 2400, "传入openid已被使用"
				}
				continue
			}

			if err2 != nil {
				return 503, err2.Error()
			}
			log.Println(string(responseByte))
			responseJson := &responseJsonStruct{}
			err2 = json.Unmarshal(responseByte, responseJson)

			if err2 != nil {
				return 503, err2.Error()
			}

			if responseJson.Status == 204 {
				return 2401, "cqupt账号或密码错误!"
			} else if responseJson.Status == 200 {
				if &responseJson.Data.Name == nil || &responseJson.Data.Xh == nil || &responseJson.Data.Sex == nil {
					return 2502, "无法获取身份信息"
				}
				ok, code, message := database.CquptBindToDatabase(username, password, openid, mail, responseJson.Data.Name, responseJson.Data.Xh, responseJson.Data.Sex)
				if !ok {
					return code, message
				}
				return 2200, "成功绑定"
			} else if responseJson.Status == 403 {
				return 2504, "向cqupt服务器请求失败"
			} else {
				log.Printf("请求绑定cqupt时错误，校服务器返回未知代码:%v", responseJson.Status)
				return 2503, "校服务器返回未知代码"
			}
		}
	}
}
