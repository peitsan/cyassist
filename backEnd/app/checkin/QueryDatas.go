package checkin

import (
	"checkinWebsite/database/checkinDatabase"
	"encoding/json"
	"log"
	"time"
)

func QueryDatas(mail string) (code int, responseString string) {
	model, err := checkinDatabase.GetCheckinModel(mail)

	type responseJson struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}

	if err != nil {
		marshal, err := json.Marshal(responseJson{404, "未查到用户信息", nil})
		if err != nil {
			log.Printf("在QueryDatas完且结果为用户不存在，生成返回json时err:%v", err.Error())
			return 500, "生成json失败"
		}
		return 200, string(marshal)
	}

	data := struct {
		Longitude  string    `json:"longitude"`
		Latitude   string    `json:"latitude"`
		Xxdz       string    `json:"xxdz"`
		Ywjchblj   bool      `json:"ywjchblj"`
		Xjzdywqzbl bool      `json:"xjzdywqzbl"`
		Twsfzc     bool      `json:"twsfzc"`
		Ywytdzz    bool      `json:"ywytdzz"`
		Ywjcqzbl   string    `json:"ywjcqzbl"`
		Beizhu     string    `json:"beizhu"`
		Jkmresult  string    `json:"jkmresult"`
		Time       time.Time `json:"time"`
		Offset     int       `json:"offset"`
	}{model.Longitude, model.Latitude, model.Xxdz, model.Ywjchblj, model.Xjzdywqzbl, model.Twsfzc, model.Ywytdzz, model.Ywjcqzbl, model.Beizhu, model.Jkmresult, model.Time, model.Offset}

	marshal, err := json.Marshal(responseJson{200, "success", data})
	if err != nil {
		log.Printf("在QueryDatas完且结果为用户存在，生成返回json时err:%v", err.Error())
		return 500, "生成json失败"
	}
	return 200, string(marshal)

}
