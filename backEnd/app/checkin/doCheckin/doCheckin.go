package doCheckin

import (
	"checkinWebsite/app/checkin/timer"
	"checkinWebsite/app/utils"
	"checkinWebsite/database/checkinDatabase"
	"checkinWebsite/database/userDatabase"
	"checkinWebsite/joshuaRequest"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type TasksTimeWheel struct {
	TimeWheel timer.Timer
	Nodes     map[string]timer.TimeNoder //用来记录所有node节点，以便快速查找方便停止特定任务  todo 有待优化空间
}

var Tasks *TasksTimeWheel

// CheckinController checkin核心管理程序
func CheckinController() {
	//初始化任务，将只启动还未到时间任务
	c := make(chan *TasksTimeWheel)
	go GenerateTimeWheel(time.Now(), c)
	Tasks = <-c
	if Tasks != nil {
		go Tasks.TimeWheel.Run()
	}

	//每天11:58分会进行第二天的队列生成，0:00会将生成好的队列启动并关闭上一日time wheel
	changeTime, generateTime := scheduleTime()
	ticket := time.NewTicker(1 * time.Second)

	for {
		rcTime := <-ticket.C
		if rcTime.Equal(generateTime) {
			log.Println("开始队列排序")
			go GenerateTimeWheel(time.Now(), c)
		}
		if rcTime.Equal(changeTime) {
			log.Println("开始每日更换队列")
			Tasks.TimeWheel.Stop()
			Tasks = <-c
			if Tasks != nil {
				go Tasks.TimeWheel.Run()
			}
			changeTime, generateTime = scheduleTime()
			log.Println("完成每日更换队列")
		}
	}
}

// DoCheckin 执行打卡
func DoCheckin(mail string) (err error, data string, uid string) {
	url := "https://we.cqupt.edu.cn/api/mrdk/get_mrdk_list_test.php"
	headers := map[string][]string{
		"Host":         []string{"we.cqupt.edu.cn"},
		"User-Agent":   []string{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat"},
		"content-type": []string{"application/json"},
	}

	checkinModel, err := checkinDatabase.GetCheckinModel(mail)
	userModel, err := userDatabase.GetCheckinModel(mail)

	if err != nil {
		log.Printf("用户%v, 到点执行打卡时读取数据库异常err:%v", mail, err.Error())
		return err, "", ""
	}

	ywjchblj := "无"
	xjzdywqzbl := "无"
	twsfzc := "是"
	ywytdzz := "无"
	if checkinModel.Ywjchblj == true {
		ywjchblj = "有"
	}
	if checkinModel.Xjzdywqzbl == true {
		xjzdywqzbl = "有"
	}
	if checkinModel.Twsfzc == false {
		twsfzc = "否"
	}
	if checkinModel.Ywytdzz == true {
		ywytdzz = "有"
	}
	rand.Seed(time.Now().UnixNano())
	latitude := checkinModel.Latitude[0:len(checkinModel.Latitude)-1] + strconv.Itoa(rand.Intn(10))
	longitude := checkinModel.Longitude[0:len(checkinModel.Longitude)-1] + strconv.Itoa(rand.Intn(10))

	var r = [...]string{"s9ZS", "jQkB", "RuQM", "O0_L", "Buxf", "LepV", "Ec6w", "zPLD", "eZry", "QjBF", "XPB0", "zlTr", "YDr2", "Mfdu", "HSoi", "frhT", "GOdB", "AEN0", "zX0T", "wJg1", "fCmn", "SM3z", "2U5I", "LI3u", "3rAY", "aoa4", "Jf9u", "M69T", "XCea", "63gc", "6_Kf"}
	var u = [...]string{"89KC",
		"pzTS",
		"wgte",
		"29_3",
		"GpdG",
		"FDYl",
		"vsE9",
		"SPJk",
		"_buC",
		"GPHN",
		"OKax",
		"_Kk4",
		"hYxa",
		"1BC5",
		"oBk_",
		"JgUW",
		"0CPR",
		"jlEh",
		"gBGg",
		"frS6",
		"4ads",
		"Iwfk",
		"TCgR",
		"wbjP"}

	timeNow := time.Now()
	mrdkey := r[timeNow.Day()%31] + u[timeNow.Hour()]

	jsonStruct := struct {
		Name          string `json:"name"`
		Xh            string `json:"xh"`
		Xb            string `json:"xb"`
		Openid        string `json:"openid"`
		Szdq          string `json:"szdq"`
		Xxdz          string `json:"xxdz"`
		LocationBig   string `json:"locationBig"`
		LocationSmall string `json:"locationSmall"`
		Latitude      string `json:"latitude"`
		Longitude     string `json:"longitude"`
		Ywjcqzbl      string `json:"ywjcqzbl"`
		Ywjchblj      string `json:"ywjchblj"`
		Xjzdywqzbl    string `json:"xjzdywqzbl"`
		Twsfzc        string `json:"twsfzc"`
		Ywytdzz       string `json:"ywytdzz"`
		Jkmresult     string `json:"jkmresult"`
		Beizhu        string `json:"beizhu"`
		Mrdkkey       string `json:"mrdkkey"` //校验码
		Timestamp     string `json:"timestamp"`
	}{userModel.Name, userModel.Xh, userModel.Sex, userModel.Openid, checkinModel.Szdq, checkinModel.Xxdz, checkinModel.LocationBig, checkinModel.LocationSmall, latitude, longitude, checkinModel.Ywjcqzbl, ywjchblj, xjzdywqzbl, twsfzc, ywytdzz, checkinModel.Jkmresult, checkinModel.Beizhu, mrdkey, strconv.FormatInt(timeNow.Unix(), 10)}

	uid = checkinModel.Uid

	jsonBytes, err := json.Marshal(&jsonStruct)
	if err != nil {
		log.Printf("在用户:%v, 构造打卡请求消息时json转换失败err:%v", mail, err.Error())
		return err, "", uid
	}
	payload := base64.StdEncoding.EncodeToString(jsonBytes)

	options := []joshuaRequest.FuncRequestOption{joshuaRequest.SetTimeout(2 * time.Second), joshuaRequest.SetHeader(headers), joshuaRequest.SetData(fmt.Sprintf("{\"key\":\"%s\"}", payload))}
	fmt.Printf(fmt.Sprintf("{\"key\":\"%s\"}", payload))

	response, err := joshuaRequest.Request(joshuaRequest.MethodPost, url, options...)

	if err != nil {
		log.Printf("在用户:%v, 发送打卡请求失败err:%v", mail, err.Error())
		return err, "", uid
	}

	all, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("在用户:%v, 读取返回body失败err:%v", mail, err.Error())
		return err, "", uid
	}

	stringAll := string(all)
	if strings.Contains(stringAll, "200") {
		return nil, stringAll, uid
	} else {
		log.Printf("在用户:%v, 打卡失败校服务器返回data:%v", mail, stringAll)
		return errors.New("校服服务器返回打卡失败"), stringAll, uid
	}

}

// GenerateFuncCallback 时间轮任务入口
func GenerateFuncCallback(mail string, timer timer.Timer, times int) func() {
	return func() {
		println("开始打卡啦")
		err, message, uid := DoCheckin(mail)

		var data string
		ok := true
		var d time.Duration

		if err != nil {
			println("dakachucuo")
			ok = false
			switch times {
			case 0:
				data = "第一次打卡失败,十分钟后自动重试"
				d = 10 * time.Minute
			case 1:
				data = "第二次打卡失败,十五分钟后自动重试"
				d = 15 * time.Minute
			case 2:
				data = "第三次打卡失败,半小时后自动重试"
				d = 30 * time.Minute
			case 3:
				data = "今日打卡失败,请手动打卡"
			}

			//除了这个错误之外，其它情况校服务器都是没有返回的
			if err.Error() != "校服服务器返回打卡失败" {
				message = err.Error()
			}
		} else {
			switch times {
			case 0:
				data = "打卡成功"
			case 1:
				data = "第二次打卡成功"
			case 2:
				data = "第三次打卡成功"
			case 3:
				data = "最后一次打卡成功"
			}
		}

		//发送至微信
		var wxErr error
		if uid != "" {
			println("wx")
			wxErr = utils.SendToWechat(uid, message, data)
		}

		if wxErr != nil {
			data += "(此消息微信推送失败)"
		}

		checkinDatabase.StoreCheckinLog(mail, data, message)

		if !ok && times != 3 && timer != nil {
			timer.AfterFunc(d, GenerateFuncCallback(mail, timer, times+1))
		}

	}
}

// GenerateTimeWheel 生成每日任务队列
func GenerateTimeWheel(t time.Time, c chan *TasksTimeWheel) {
	users := checkinDatabase.GetAllCheckUsers(t)
	if users == nil {
		c <- nil
		return
	}
	timeWheel := timer.NewTimer()
	tasks := &TasksTimeWheel{timeWheel, make(map[string]timer.TimeNoder)}

	for _, i := range users {
		fmt.Printf("%v\n", i.Time.Sub(time.Now()))
		tasks.Nodes[i.Mail] = timeWheel.AfterFunc(i.Time.Sub(time.Now()), GenerateFuncCallback(i.Mail, timeWheel, 0))
	}
	//队列排序完成，等待指定时间切换
	c <- tasks
	return
}

//生成每天产生队列以及更改队列的时间
func scheduleTime() (changeTime time.Time, generateTime time.Time) {
	tempTime := time.Now().Add(24 * time.Hour)
	changeTime = time.Date(tempTime.Year(), tempTime.Month(), tempTime.Day(), 0, 0, 0, 0, tempTime.Location())
	generateTime = changeTime.Add(-2 * time.Minute)
	return
}
