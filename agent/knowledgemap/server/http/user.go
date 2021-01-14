package http

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"knowledgemap_backend/agent/knowledgemap/server/http/comm"
	"knowledgemap_backend/microservices/common/middlewares"
	papi "knowledgemap_backend/microservices/knowledgemap/passport/api"
	uapi "knowledgemap_backend/microservices/knowledgemap/user/api"
	"net/http"
	"os"
	"pdd/store/model/cm"

	"github.com/astaxie/beego/httplib"
	"github.com/labstack/echo"
)

func userInfo(c echo.Context) error {
	clog := middlewares.Log(c)
	clog.Info("getUserInfo")

	// userId := c.Request().Header.Get("auth-uid")
	userId := c.Get("userId").(string)
	// fmt.Println(userId, c.Get("userType"))
	response, err := userSrv.UserInfo(context.TODO(), &uapi.UserReq{
		Userid:   userId,
		Identify: uapi.Identify(c.Get("userType").(int64)),
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(200, comm.Data(response))

}

func userRegister(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(papi.RegisterReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	//fmt.Println("req is", req.Rtype, req.Name, req.Major, req.College, req.Sex, req.Account, req.Password)
	if res, err := passportSrv.Register(context.TODO(), req); err != nil {
		clog.Errorf("call register %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
	// return c.JSON(http.StatusOK, comm.Data("response"))
}

func userLogin(c echo.Context) error {
	fmt.Println("进login")
	clog := middlewares.Log(c)
	req := new(papi.LoginReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	fmt.Println(req.Usertype, req.Password, req.Account)
	if res, err := passportSrv.Login(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		if res.User.Userid != "" {
			return c.JSON(http.StatusOK, comm.Data(res))
		}
		return c.JSON(http.StatusOK, comm.Err("need register", comm.STATUS_PHONE_NOT_REGISTER))
	}
}

func userChangePassword(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(papi.ChangePasswordReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	req.Userid = c.Get("userId").(string)
	req.Identify = uapi.Identify(c.Get("userType").(int64))
	if res, err := passportSrv.ChangePassword(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

func userChangeInfo(c echo.Context) error {
	clog := middlewares.Log(c)
	req := new(papi.ChangeUserInfoReq)
	if err := comm.VBind(c, req); err != nil {
		clog.Errorf("参数错误:%v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error(), comm.STATUS_INVALIDE_ARGS))
	}
	// req.Userid = c.Request().Header.Get("auth-uid")
	// userType, _ := strconv.ParseInt(c.Request().Header.Get("auth-type"), 10, 64)
	req.Userid = c.Get("userId").(string)
	req.Usertype = uapi.Identify(c.Get("userType").(int64))
	if res, err := passportSrv.ChangeUserInfo(context.TODO(), req); err != nil {
		clog.Error("error %v", err)
		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
	} else {
		return c.JSON(http.StatusOK, comm.Data(res))
	}
}

// func transport(c echo.Context) error {
// 	clog := middlewares.Log(c)
// 	req := new(papi.TransportReq)
// 	f, err := ioutil.ReadFile("/Users/firewinggames/Downloads/transport_origin.jpeg")
// 	if err != nil {
// 		fmt.Println("read fail", err)
// 	}
// 	req.Content = f
// 	if res, err := passportSrv.TestTransport(context.TODO(), req); err != nil {
// 		clog.Error("error %v", err)
// 		return c.JSON(http.StatusBadRequest, comm.Err(err.Error()))
// 	} else {
// 		return c.JSON(http.StatusOK, comm.Data(res))
// 	}
// }

func transport(c echo.Context) error {
	clog := middlewares.Log(c)
	mfroot, err := c.MultipartForm()
	if err != nil {
		clog.Error("%v\n", err)
		return c.JSON(http.StatusCreated, cm.Bad(err.Error()))
	}
	files := mfroot.File["file"]

	if len(files) != 1 {
		clog.Error("没有图片或文件\n")
		return c.JSON(http.StatusCreated, cm.Bad("请上传图片或文件"))
	}
	file := files[0]
	fmt.Printf("filename:%v\n", file.Filename)
	// ssss := strings.Split(file.Filename, ".")
	src, err := file.Open()
	if err != nil {
		clog.Error("%v", err)
		return c.JSON(http.StatusCreated, cm.Bad(err.Error()))
	}
	defer src.Close()
	data, err := ioutil.ReadAll(src)
	if err != nil {
		clog.Error("%v", err)
		return c.JSON(http.StatusCreated, cm.Bad(err.Error()))
	}

	logFile := "../" + file.Filename
	WriteFile(logFile, data, os.ModeAppend)

	var obj interface{}
	addr := "http://47.95.145.171:8080/upload"
	// profile := os.Getenv("profile")
	// if profile != "debug" {
	// 	addr = "http://172.17.9.156:8080/group1/upload"
	// 	//addr = "http://47.95.145.171:8080/group1/upload"

	// }
	fmt.Println(addr)
	uploadReq := httplib.Post(addr)
	uploadReq.PostFile("file", logFile) //注意不是全路径
	uploadReq.Param("output", "json")
	uploadReq.Param("scene", "")
	uploadReq.Param("path", "")
	a := uploadReq.GetRequest()
	fmt.Println(a)
	b, err := uploadReq.Response()
	fmt.Println(b, err)
	uploadReq.ToJSON(&obj)
	fmt.Println(obj)
	req := obj.(map[string]interface{})
	fmt.Println(req["md5"])
	fmt.Println(req["url"])

	//os.Remove(logFile)
	requst := FileRes{req["md5"].(string), req["url"].(string)}
	return c.JSON(http.StatusOK, comm.Data(requst))

}

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

type FileRes struct {
	FileId string `json:"fileid"`
	Url    string `json:"url"`
}
