package comm

import (
	"net/http"

	"knowledgemap_backend/microservices/common/middlewares"

	"github.com/labstack/echo"
)

type COMMSTATUS int
type ApiRes struct {
	Msg  string      `json:"msg"`
	Code COMMSTATUS  `json:"code"`
	Data interface{} `json:"data"`
}

const (
	STATUS_OK                       COMMSTATUS = 0
	STATUS_BAD_REQUEST              COMMSTATUS = 1
	STATUS_NEED_LOGIN               COMMSTATUS = 2
	STATUS_WRONG_PHONE              COMMSTATUS = 3
	STATUS_WRONG_MCODE              COMMSTATUS = 4
	STATUS_WRONG_ICODE              COMMSTATUS = 5
	STATUS_PHONE_ALREADY_REGISTER   COMMSTATUS = 6
	STATUS_ICODE_MAX                COMMSTATUS = 7
	STATUS_PHONE_NOT_REGISTER       COMMSTATUS = 8
	STATUS_MCODE_TIMEOUT            COMMSTATUS = 9
	STATUS_NICHENG_LIMIT            COMMSTATUS = 10
	STATUS_SUBSCRIPTION_REPEAT      COMMSTATUS = 11
	STATUS_SUBSCRIPTION_ERROR       COMMSTATUS = 12
	STATUS_TRANSACTION_NOTCOM_ERROR COMMSTATUS = 13
	STATUS_TRANSACTION_OTHER_ERROR  COMMSTATUS = 14
	STATUS_EMAIL_ALREADY_BIND       COMMSTATUS = 15
	STATUS_TRANSACTION_NEWBIE       COMMSTATUS = 16
	STATUS_TRANSACTION_COOLDOWN     COMMSTATUS = 17
	STATUS_EMAIL_SEND_ERROR         COMMSTATUS = 18
	STATUS_COMMON                   COMMSTATUS = 10000
	STATUS_MAINTAIN                 COMMSTATUS = 10001
	STATUS_BLACKLIST                COMMSTATUS = 10002
	STATUS_VERIFY_ERROR             COMMSTATUS = 10003
	STATUS_TANABATA_NOT_START       COMMSTATUS = 10004
	STATUS_TANABATA_HAS_END         COMMSTATUS = 10005
	STATUS_MONEY_NOT_ENOUGH         COMMSTATUS = 10006
	STATUS_OVER_ENERGY_MAXLIMIT     COMMSTATUS = 10007
	STATUS_FIND_WRONG               COMMSTATUS = 10008
	STATUS_PROPOS_NOT_ENOUGH        COMMSTATUS = 10009
	STATUS_OVER_REBORN_LIMIT        COMMSTATUS = 10010
	STATUS_DROP_PROPS_ILLEGAL       COMMSTATUS = 10011
	STATUS_ACTIVITY_NOT_START       COMMSTATUS = 10012
	STATUS_OVER_TIME                COMMSTATUS = 10013
	STATUS_LEARDBOARD_WRONG         COMMSTATUS = 10014
	STATUS_OVER_LIMIT               COMMSTATUS = 10015
	STATUS_NO_WECHAT                COMMSTATUS = 10016
	STATUS_ACCOUNT_WRONG            COMMSTATUS = 10017
	STATUS_NEED_VERIFY              COMMSTATUS = 10018
	STATUS_VERIFY_WRONG             COMMSTATUS = 10019
	STATUS_NO_WECHATPUBLIC          COMMSTATUS = 10020
	STATUS_NO_WITHDRAW              COMMSTATUS = 10021
	STATUS_WRONG_INDEX              COMMSTATUS = 10022
	STATUS_INSUFFICIENT_KETH        COMMSTATUS = 10023
	STATUS_INVALIDE_ARGS            COMMSTATUS = 10024
	STATUS_NEED_ICODE               COMMSTATUS = 10025
	STATUS_INVALIDE_FILLCODE        COMMSTATUS = 10026
)

func OK(msg ...string) *ApiRes {
	var res ApiRes
	if len(msg) != 0 {
		res.Msg = msg[0]
	}
	res.Code = STATUS_OK
	return &res
}

func Data(data interface{}, msg ...string) *ApiRes {
	var res ApiRes
	res.Code = STATUS_OK
	res.Data = data
	if len(msg) != 0 {
		res.Msg = msg[0]
	}
	return &res
}

func DataWithCode(code COMMSTATUS, data interface{}, msg ...string) *ApiRes {
	var res ApiRes
	res.Code = code
	res.Data = data
	if len(msg) != 0 {
		res.Msg = msg[0]
	}
	return &res
}

func Err(errstr string, errcode ...COMMSTATUS) *ApiRes {
	var res ApiRes
	res.Msg = errstr
	res.Code = STATUS_BAD_REQUEST
	if len(errcode) != 0 {
		res.Code = errcode[0]
	}
	return &res
}

func VBind(c echo.Context, req interface{}) (err error) {

	clog := middlewares.Log(c)
	if err = c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, Err(err.Error(), STATUS_INVALIDE_ARGS))
	}
	if err = c.Validate(req); err != nil {
		clog.Error(err)
		c.JSON(http.StatusBadRequest, Err("", STATUS_INVALIDE_ARGS))
	}
	return
}
