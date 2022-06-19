package checkin

import (
	"checkinWebsite/app/checkin/doCheckin"
	"checkinWebsite/database/checkinDatabase"
	"log"
)

func DoImmediately(mail string) (code int, message string) {
	//先停止,再创建新任务
	result := doCheckin.Tasks.Nodes[mail]
	if result != nil {
		result.Stop()
	}
	ok, err := checkinDatabase.FindCheckinModel(mail)
	if err != nil {
		return 500, "数据库错误"
	}
	if ok {
		doCheckin.Tasks.TimeWheel.AfterFunc(0, doCheckin.GenerateFuncCallback(mail, doCheckin.Tasks.TimeWheel, 0))
		return 2200, "success"
	} else {
		log.Printf("DoImmediately执行时未在数据库发现用户%v信息，却触发此接口，疑似有人抓接口！", mail)
		return 404, ""
	}

}
