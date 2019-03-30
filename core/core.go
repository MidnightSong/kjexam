package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"kjexam/models"
	"strconv"
	"strings"
	"sync"

	resty "gopkg.in/resty.v1"
)

//Core .
type Core struct {
	Client           *resty.Client
	Mutex            sync.Mutex
	DoCount          int
	AttemptToken     string
	LearnerAttemptID string
}

//AddCount .
func (core *Core) AddCount() {
	core.Mutex.Lock()
	defer core.Mutex.Unlock()

	core.DoCount++
}

//DoOneCourse .
func (core *Core) DoOneCourse(classID int) {
	resp, err := core.Client.R().Get(fmt.Sprintf("http://jiangxi.e-nai.cn/api/learner/course/%d/outline/tree", classID))
	if err != nil {
		fmt.Println(err)
		return
	}

	res := models.Session{}
	json.Unmarshal(resp.Body(), &res)
	ss := getAllCourse(res)

	wg := sync.WaitGroup{}

	for index, val := range ss {
		wg.Add(1)
		go func(i int, v models.Session) {
			defer wg.Done()

			if strings.TrimSpace(v.Status) != "C" {
				if err := core.DoOneVideo(classID, v.ID); err != nil {
					fmt.Println("CourseID:", classID, "rco.id", v.ID, "error :", err.Error())
				}
				fmt.Println("ID:", v.ID, "title:", v.Title, "状态:", v.Status, "执行完成!")
			} else {
				fmt.Println("ID:", v.ID, "title:", v.Title, "状态:", v.Status)
			}
		}(index, val)
	}

	wg.Wait()
}

//DoOneVideo .
func (core *Core) DoOneVideo(courseID, rcoID int) error {
	if courseID == 0 {
		return errors.New("courseID must have a value")
	}

	c := resty.R()
	form := map[string]string{
		"rawStatus":        "completed",
		"credit":           "no-credit",
		"attemptToken":     "cc0b9eed-60fb-4fef-8141-a829cb3f40db",
		"learnerAttemptId": "1544607437896",
		"course.id":        strconv.Itoa(courseID),
		"classroom.id":     "330",
		"rco.id":           strconv.Itoa(rcoID),
		"sessionTime":      "00:00:02",
	}

	c.SetFormData(form)
	_, err := c.Post(fmt.Sprintf("http://jiangxi.e-nai.cn/api/learner/play/course/%d/save", courseID))
	if err != nil {
		return err
	}

	return nil
}

func getAllCourse(ss models.Session) []models.Session {
	ssA := []models.Session{}
	for _, val := range ss.Children {
		if val.Children != nil {
			ssA = append(ssA, getAllCourse(val)...)
		} else {
			ssA = append(ssA, val)
		}
	}
	return ssA
}
