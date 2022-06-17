package checkinDatabase

import (
	"checkinWebsite/database"
	"errors"
	"log"
	"time"
)

type CheckinModel struct {
	mail          string    //记录对应数据库mail，不可更改
	Time          time.Time `json:"time" form:"time" binding:"required"`                    //每日打开设定时间
	Offset        int       `json:"offset" form:"offset" `                                  //每日打卡时间偏移范围
	Mail          string    `json:"mail" form:"mail" db:"mail"`                             //邮箱
	Szdq          string    `json:"szdq" form:"szdq" `                                      //所在地区
	Xxdz          string    `json:"xxdz" form:"xxdz" binding:"required"`                    //详细地址
	LocationBig   string    `json:"locationBig" form:"locationBig" db:"locationBig" `       //中国,重庆市,重庆市,南岸区
	LocationSmall string    `json:"locationSmall" form:"locationSmall" db:"locationSmall" ` //重庆市南岸区崇文路
	Latitude      string    `json:"latitude" form:"latitude" binding:"required"`            //29.53xx  维度
	Longitude     string    `json:"longitude" form:"longitude" binding:"required"`          //106.60xx 经度
	Ywjcqzbl      string    `json:"ywjcqzbl" form:"ywjcqzbl" binding:"required"`            //新冠肺炎风险：低风险、中风险、高风险
	Ywjchblj      bool      `json:"ywjchblj" form:"ywjchblj" `                              //14天高风险旅居史
	Xjzdywqzbl    bool      `json:"xjzdywqzbl" form:"xjzdywqzbl" `                          //14天接触高风险人员
	Twsfzc        bool      `form:"twsfzc" json:"twsfzc" `                                  //体温是否正常
	Ywytdzz       bool      `json:"ywytdzz" form:"ywytdzz" `                                //有无感染有关症状
	Jkmresult     string    `json:"jkmresult" form:"jkmresult" binding:"required"`          //渝康码颜色：绿色、黄色、红色、其它
	Beizhu        string    `json:"beizhu" form:"beizhu" `                                  //备注
	Uid           string    `json:"uid" form:"uid"`                                         //微信推送
}

func GetCheckinModel(mail string) (*CheckinModel, error) {
	model := new(CheckinModel)
	err := database.Db.Get(model, "SELECT * FROM checkin WHERE mail=? LIMIT 1", mail)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	model.mail = model.Mail
	return model, nil
}

// FindCheckinModel 查询用户是否录入过信息
func FindCheckinModel(mail string) (bool, error) {
	var count int
	err := database.Db.Get(&count, "select count(*) from checkin where mail = ?", mail)
	if err != nil {
		log.Printf("查询用户是否录入过信息时，发生错误err:%v", err.Error())
		return false, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func AddCheckinModel(model CheckinModel) error {
	if &model.Uid == nil {
		model.Uid = ""
	}
	sql := "INSERT INTO checkin values (:mail,:time,:offset,:szdq,:xxdz,:locationBig,:locationSmall,:latitude,:longitude,:ywjcqzbl,:ywjchblj,:xjzdywqzbl,:twsfzc,:ywytdzz,:jkmresult,:beizhu,:uid)"
	_, err := database.Db.NamedExec(sql, &model)
	if err != nil {
		log.Printf("新建打卡信息时发生错误，user:%v,err:%v", model.Mail, err.Error())
		return err
	}
	return nil
}

func (c *CheckinModel) UpdateAllCheckinModel(model CheckinModel) error {
	if &model.Uid == nil {
		model.Uid = ""
	}
	sql := "update checkin set mail=?,szdq=?,xxdz=?,locationBig=?,locationSmall=?,latitude=?,longitude=?,ywjcqzbl=?,ywjchblj=?,xjzdywqzbl=?,twsfzc=?,ywytdzz=?,jkmresult=?,beizhu=?,uid=? WHERE mail=?"
	_, err := database.Db.Exec(sql, model.Mail, model.Szdq, model.Xxdz, model.LocationBig, model.LocationSmall, model.Latitude, model.Longitude, model.Ywjcqzbl, model.Ywjchblj, model.Xjzdywqzbl, model.Twsfzc, model.Ywytdzz, model.Jkmresult, model.Beizhu, model.Uid, c.mail)
	if err != nil {
		log.Printf("更新打卡信息时发生错误，user:%v,err:%v", model.Mail, err.Error())
		return err
	}
	return nil
}

func (c *CheckinModel) Save() error {
	sql := "update checkin set mail=?,szdq=?,xxdz=?,locationBig=?,locationSmall=?,latitude=?,longitude=?,ywjcqzbl=?,ywjchblj=?,xjzdywqzbl=?,twsfzc=?,ywytdzz=?,jkmresult=?,beizhu=?,uid=? WHERE mail=?"
	_, err := database.Db.Exec(sql, c.Mail, c.Szdq, c.Xxdz, c.LocationBig, c.LocationSmall, c.Latitude, c.Longitude, c.Ywjcqzbl, c.Ywjchblj, c.Xjzdywqzbl, c.Twsfzc, c.Ywytdzz, c.Jkmresult, c.Beizhu, c.Uid, c.mail)
	if err != nil {
		print(err.Error())
		return err
	}
	return nil
}
