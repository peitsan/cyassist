package checkin

import (
	"checkinWebsite/database/checkinDatabase"
	"checkinWebsite/joshuaRequest"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

func SubmitDatas(info checkinDatabase.CheckinModel) (code int, message string) {
	err := queryMapApi(&info)
	if err != nil && err.Error() == "out china" {
		return 2400, "打卡地区非法"
	} else if err != nil {
		return 500, "无法获取地图api信息"
	}
	ok, err := checkinDatabase.FindCheckinModel(info.Mail)
	if err != nil {
		return 500, "数据库错误"
	}
	if ok {
		m, err := checkinDatabase.GetCheckinModel(info.Mail)
		if err != nil {
			return 500, "数据库错误"
		}
		err = m.UpdateAllCheckinModel(info)
		if err != nil {
			return 500, "数据库错误"
		}
	} else {
		err := checkinDatabase.AddCheckinModel(info)
		if err != nil {
			return 500, "数据库错误"
		}
	}
	return 2200, "success"
}

func SubmitUID(mail, uid string) (code int, message string) {
	m, err := checkinDatabase.GetCheckinModel(mail)
	if err != nil && err.Error() == "用户不存在" {
		log.Printf("在SubmitUID中存入UID失败，原因数据库不存在该用户,mail:%v,uid:%v", mail, uid)
		return 2400, "用户不存在"
	} else if err != nil {
		log.Printf("在SubmitUID中存入UID失败，err:%v,mail:%v,uid:%v", err, mail, uid)
		return 500, "数据库错误"
	}
	m.Uid = uid
	err = m.Save()
	if err != nil {
		log.Printf("在SubmitUID中获取用户model后存入数据库失败UID失败err:%v,mail:%v,uid:%v", err, mail, uid)
		return 500, "数据库错误"
	}
	return 2200, "success"

}

func queryMapApi(info *checkinDatabase.CheckinModel) error {
	//返回json, 结构体构造
	type AdInfo struct {
		Name     string `json:"name"`
		Province string `json:"province"`
		City     string `json:"city"`
		District string `json:"district"`
	}

	type Result struct {
		Address string `json:"address"`
		AdInfo  AdInfo `json:"ad_info"`
	}

	type BackJson struct {
		Result Result `json:"result"`
	}

	response, err := joshuaRequest.Request("GET", fmt.Sprintf("https://apis.map.qq.com/ws/geocoder/v1?location=%v%%2C%v&key=7IMBZ-XWMWW-D4FR5-R3NAG-G7A7S-FMBFN", info.Latitude, info.Longitude))
	if err != nil {
		log.Printf("查询地图api请求response时出现错误,err:%v", err.Error())
		return err
	}
	backJson := &BackJson{}
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("查询地图api获取返回后转换调用readAll读取时出现错误,err:%v", err.Error())
		return err
	}
	err = json.Unmarshal(all, backJson)
	if err != nil {
		log.Printf("查询地图api获取返回后转换json时出现错误,err:%v", err.Error())
		return err
	}

	//根据实验，如果所选区域超出中国，则province和nation和district返回空，返回无效
	if &backJson.Result.AdInfo.Province == nil {
		log.Printf("查询地图api时传入经纬度非法, longtitude:%v,latitude:%v", info.Longitude, info.Latitude)
		return errors.New("out china")
	}

	info.LocationBig = backJson.Result.AdInfo.Name
	info.LocationSmall = backJson.Result.Address

	//we重邮打卡信息的sszq（所在地区）中，有比较特殊的地方，就是对重庆市非主城区地区，他提交时是(重庆市,县,xxx)，然而接口返回是(重庆市,重庆市,xxx)，所以要处理下
	temp := backJson.Result.AdInfo.District
	if temp == "城口县" || temp == "丰都县" || temp == "垫江县" || temp == "忠县" || temp == "云阳县" || temp == "奉节县" || temp == "巫山县" || temp == "巫溪县" || temp == "石柱土家族自治县" || temp == "秀山土家族苗族自治县" || temp == "酉阳土家族苗族自治县" || temp == "彭水苗族土家族自治县" {
		backJson.Result.AdInfo.City = "县"
	}
	info.Szdq = fmt.Sprintf("%v,%v,%v", backJson.Result.AdInfo.Province, backJson.Result.AdInfo.City, backJson.Result.AdInfo.District)
	return nil
}
