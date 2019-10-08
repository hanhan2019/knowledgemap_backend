package model

import "gopkg.in/mgo.v2/bson"

const USER_COLLECTION_NAME = "users"

const (
	EMAILSTATUS_NOTSET    = "NOTSET"
	EMAILSTATUS_NOTVERIFY = "NOTVERIFY"
	EMAILSTATUS_GOOD      = "good"
)

type Email struct {
	Addr   string `bson:"addr" json:"addr"`
	Status string `bson:"status" json:"status"`
}
type FillCode struct {
	Wechat bool `bson:"wechat" json:"wechat"`
	QQ     bool `bson:"qq" json:"qq"`
	Kabao  bool `bson:"kabao" json:"kakao"`
}
type User struct {
	ID             bson.ObjectId `bson:"_id" json:"_id"`
	UserName       int64         `bson:"username" json:"username"`
	NiCheng        string        `bson:"nicheng" json:"nicheng"`
	Index          int64         `bson:"index" json:"index"` // 第几个人物
	Telephone      string        `bson:"telephone" json:"telephone"`
	Icode          string        `bson:"icode" json:"icode"`   // myself icode
	FIcode         string        `bson:"ficode" json:"ficode"` // who invites me
	MyAvailableSub int64         `bson:"mysub" json:"mysub"`   // 我的一级下线人数还有几个空位
	CreateTime     int64         `bson:"createtime" json:"createtime"`
	Suanli         int64         `bson:"suanli" json:"suanli"`
	B              int64         `bson:"b" json:"b"`
	LastLoginAt    int64         `bson:"lla" json:"lla"` // 上次上线时间
	Email          *Email        `bson:"email" json:"email"`
	FillCode       *FillCode     `bson:"fillcode" json:"fillcode"` // 微信和qq验证码是否已经填过了
	WechatID       string        `bson:"wechatid" json:"wechatid"`
	Country        string        `bson:"country" json:"country"`
	Device         string        `bson:"device" json:"device"` //上一次登陆的设备号
	ChaosticID     string        `bson:"chaosticid" json:"chaosticid"`
	Channel        string        `bson:"channel" json:"channel"`
	updater        bson.M
}
