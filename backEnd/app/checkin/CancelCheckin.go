package checkin

import (
	"checkinWebsite/app/checkin/doCheckin"
	"checkinWebsite/database/checkinDatabase"
	"log"
)

func CancelCheckin(mail string) (code int, message string) {
	model, err := checkinDatabase.GetCheckinModel(mail)
	if err != nil {
		if err.Error() == "用户不存在" {
			log.Printf("CancelCheckin执行时未在数据库发现用户%v信息，却触发此接口，疑似有人抓接口！", mail)
			return 404, ""
		}
		return 500, "数据库错误"
	}
	model.Enable = false
	err = model.Save()
	if doCheckin.Tasks != nil {
		doCheckin.Tasks.Nodes[mail].Stop() //取消当前任务队列中的任务
	}

	if err != nil {
		return 500, "数据库错误"
	}
	return 2200, "success"
}
