package checkinDatabase

import (
	"checkinWebsite/database"
	"log"
	"time"
)

type CheckinLogModel struct {
	Data string    `json:"data"`
	Time time.Time `json:"time"`
}

type CheckinLogResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Log     []CheckinLogModel `json:"log"`
}

// GetCheckinLogModel 获取该用户最近5次打卡信息，错误返回只有数据库异常
func GetCheckinLogModel(mail string) ([]CheckinLogModel, error) {
	var logs []CheckinLogModel
	sql := "select data,time from checkinlog where mail = ? order by time desc limit 5"
	err := database.Db.Select(&logs, sql, mail)
	if err != nil {
		log.Printf("获取用户%v最近打卡信息时，数据库错误err:%v", mail, err.Error())
		return nil, err
	}
	return logs, nil
}

func StoreCheckinLog(mail string, data string, logs string) {
	sql := "insert into checkinlog values (?,?,?,?)"
	_, err := database.Db.Exec(sql, mail, data, logs, time.Now())
	if err != nil {
		log.Printf("在储存用户%v，打卡信息时数据库异常,存入数据为data:%v,log:%v", mail, data, logs)
	}
}
