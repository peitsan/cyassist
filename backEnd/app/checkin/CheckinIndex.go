package checkin

import "checkinWebsite/database/checkinDatabase"

func CheckinIndex(mail string) (code int, message string, data []checkinDatabase.CheckinLogModel) {
	exist, err := checkinDatabase.FindCheckinModel(mail)
	if err != nil {
		return 500, "数据库错误", nil
	}
	if exist {
		data, err := checkinDatabase.GetCheckinLogModel(mail)
		if err != nil {
			return 500, "数据库错误", nil
		}
		return 2200, "success", data
	} else {
		return 2200, "未录入", nil
	}
}
