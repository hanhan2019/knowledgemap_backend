package http

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/httplib"
)

func main() {
	var obj interface{}
	addr := "http://127.0.0.1:8080/group1/upload"
	profile := os.Getenv("profile")
	if profile != "debug" {
		addr = "http://172.17.9.156:8080/group1/upload"
		//addr = "http://47.95.145.171:8080/group1/upload"

	}
	addr = "http://47.95.145.171:8080/list_dir"
	fmt.Println(addr)
	uploadReq := httplib.Post(addr)
	uploadReq.PostFile("file", "class.go") //注意不是全路径
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

}
