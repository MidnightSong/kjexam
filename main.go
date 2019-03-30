package main

import (
	"encoding/json"
	"fmt"
	"kjexam/core"
	"kjexam/models"
	"net/http"
	"sync"

	resty "gopkg.in/resty.v1"
)

var cc core.Core

func init() {
	// resty.SetDebug(true)
	resty.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36")

	cc = core.Core{Mutex: sync.Mutex{}}
}

func main() {
	cc.Client = resty.SetCookie(&http.Cookie{Name: "SESSION", Value: "ad4d3220-d825-49e2-9d3c-050bbacc4ed6"})

	classResp, err := cc.Client.R().Get("http://jiangxi.e-nai.cn/api/learner/trainingClasses/330/courses?sort=displayOrder,asc&page=0&size=100")
	if err != nil {
		fmt.Println(err)
		return
	}

	classRes := models.Class{}
	json.Unmarshal(classResp.Body(), &classRes)
	for index, val := range classRes.Content {
		fmt.Println("Class Name:", val.Name, "Class ID:", val.ID)
		cc.DoOneCourse(val.ID)
		fmt.Println("当前:", (index + 1), "总课程:", len(classRes.Content), "进度: ", ((index + 1) * 100 / len(classRes.Content)), "%")
	}

	fmt.Println("程序结束，刷课数:", cc.DoCount)
}
