package utils

import (
	"checkinWebsite/joshuaRequest"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

var APPTOKEN string

// 微信推送消息
func SendToWechat(uid string, content string, summary string) error {
	postUrl := "http://wxpusher.zjiecode.com/api/send/message"
	sendStruct := struct {
		Apptoken string   `json:"Apptoken"`
		Content  string   `json:"content"`
		Summary  string   `json:"summary"`
		Uids     []string `json:"uids"`
	}{APPTOKEN, content, summary, []string{uid}}

	jsonBytes, err := json.Marshal(&sendStruct)
	if err != nil {
		log.Printf("在给uid:%v用户推送微信消息content:%v,summary:%v时，转换json发生err:%v", uid, content, summary, err.Error())
		return err
	}

	response, err := joshuaRequest.Request(joshuaRequest.MethodPost, postUrl, joshuaRequest.SetData(string(jsonBytes)))
	if err != nil {
		log.Printf("在给uid:%v用户推送微信消息content:%v,summary:%v时，网络请求失败err:%v", uid, content, summary, err.Error())
		return err
	}

	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("在给uid:%v用户推送微信消息content:%v,summary:%v时，读取返回body时err:%v", uid, content, summary, err.Error())
		return err
	}
	if strings.Contains(string(all), "1000") {
		return nil
	} else {
		log.Printf("在给uid:%v用户推送微信消息content:%v,summary:%v时，返回body:%v未显示推送成功，疑似推送失败", uid, content, summary, string(all))
		return errors.New("成功返回，但是未检测到成功")
	}
}
