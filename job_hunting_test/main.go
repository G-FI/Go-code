package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"job_hunting_test/middleware"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.POST("/getjobsinfo", getJobsInfo)
	r.POST("/gethistoryinfos", getHistoryInfos)
	r.POST("/getroundempinfos", getRoundEmpInfos)
	r.POST("/getroundempdetail", getRoundEmpDetail)
	r.POST("/getcurrentempinfo", getCurrentEmpInfo)
	r.POST("/deleteempinfo", deleteEmpInfo)
	r.POST("/getcompanyinfo", getCompanyInfo)
	r.POST("/createevaluation", createEvaluation)
	r.POST("/finishevaluation", finishEvaluation)
	r.POST("/postevaluation", postEvaluation)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func postEvaluation(c *gin.Context) {
	c.GetPostForm("string")
	type Resp struct {
		Msg string `json:"msg"`
	}
	c.IndentedJSON(http.StatusOK, Resp{Msg: "ok"})
}

func finishEvaluation(c *gin.Context) {
	type Resp struct {
		Msg string `json:"msg"`
	}
	isCreated = false
	c.IndentedJSON(http.StatusOK, Resp{"ok"})
}

var isCreated = false

func createEvaluation(c *gin.Context) {
	startDate, _ := c.GetPostForm("startDate")
	endDate, _ := c.GetPostForm("endDate")
	note, _ := c.GetPostForm("note")
	fmt.Println(startDate, endDate, note)
	type Resp struct {
		Msg string `json:"msg"`
	}
	//isCreated = true
	c.IndentedJSON(http.StatusOK, Resp{"ok"})
}
func getCompanyInfo(c *gin.Context) {
	type data struct {
		IsCreated bool `json:"isCreated"`
	}
	type companyResp struct {
		Msg  string `json:"msg"`
		Data data   `json:"data"`
	}

	c.IndentedJSON(http.StatusOK, companyResp{"ok", data{IsCreated: isCreated}})
}

func deleteEmpInfo(c *gin.Context) {
	empId, _ := c.GetPostForm("id")
	fmt.Printf("离职员工的ID: %s", empId)
	type Resp struct {
		Msg string `json:"msg"`
	}
	c.IndentedJSON(http.StatusOK, Resp{"ok"})
}

func getCurrentEmpInfo(c *gin.Context) {

	type data struct {
		EmpInfos []empInfo `json:"empInfos"`
	}
	type empResponse struct {
		Msg  string `json:"msg"`
		Data data   `json:"data"`
	}
	c.IndentedJSON(http.StatusOK, empResponse{"ok", data{EmpInfos: empInfos}})
}

func getRoundEmpDetail(c *gin.Context) {
	roundIndex, _ := strconv.Atoi("roundIndex")
	empId, _ := strconv.Atoi("empId")
	fmt.Printf("请求第 %d 轮评价，员工号为 %d的员工信息", roundIndex, empId)

	eval := evaluation{[]int{10, 10, 20, 100, 10, 12, 80, 100, 19, 89}, "这个员工很认真"}

	type empDetail struct {
		Id         string     `json:"id"`
		Name       string     `json:"name"`
		Type       string     `json:"type"`
		Department string     `json:"department"`
		EvInfo     evaluation `json:"evInfo"`
	}

	type data struct {
		EmpInfo empDetail `json:"empInfo"`
	}
	type evalResponse struct {
		Msg  string `json:"msg"`
		Data data   `json:"data"`
	}
	detail := empDetail{Id: "1111", Name: "小黑", Type: "主管", Department: "市场部", EvInfo: eval}
	d := data{EmpInfo: detail}
	c.IndentedJSON(http.StatusOK, evalResponse{Msg: "ok", Data: d})
}

func getRoundEmpInfos(c *gin.Context) {
	type data struct {
		EmpInfos []empInfo `json:"empInfos"`
	}
	type empResponse struct {
		Msg  string `json:"msg"`
		Data data   `json:"data"`
	}
	//获取指定的那一轮员工信息
	roundIndex, _ := strconv.Atoi(c.GetHeader("roundIndex"))
	fmt.Printf("请求了第 %d 轮员工历史评价数据\n", roundIndex)
	response := empResponse{"ok", data{allRounds[1]}}
	c.IndentedJSON(http.StatusOK, response)
}

func getJobsInfo(c *gin.Context) {

	type data struct {
		JobInfos []jobInfo `json:"job_infos"`
	}
	type jobsResp struct {
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}

	response := jobsResp{Msg: "ok", Data: data{jobInfos}}
	fmt.Println(response)
	c.IndentedJSON(http.StatusOK, response)
}

func getHistoryInfos(c *gin.Context) {

	type data struct {
		HistoryInfos []historyInfo `json:"historyInfos"`
	}
	type historyResp struct {
		Msg  string `json:"msg"`
		Data data   `json:"data"`
	}
	resData := data{historyInfos}
	response := historyResp{"ok", resData}

	c.IndentedJSON(http.StatusOK, response)
}

//员工信息box显示信息
type empInfo struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Department string `json:"department"`
}
type evaluation struct {
	Abilities []int  `json:"abilities"`
	Remark    string `json:"remark"`
}

var empInfos = []empInfo{
	{"5555555", "小黑", "普通员工", "开发部"},
	{"6666666", "小红", "主管", "设计部"},
	{"7777777", "小蓝", "经理", "产品部"},
	{"8888888", "小紫", "老板", "管理部"},
	{"9999999", "小白", "ceo", "法律部"},
	{"1010101", "小绿", "ipo", "市场部"},
	{"1212121", "小灰", "组长", "开发部"},
	{"1313131", "小黄", "领导", "开发部"},
	{"1414141", "小橙", "普通员工", "开发部"},
	{"1515151", "小青", "普通员工", "开发部"},
}

//求职市场
type jobInfo struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

var jobInfos = []jobInfo{
	{Id: "1", Title: "java web", Detail: "asdfasdfa"},
	{Id: "2", Title: "java工程师", Detail: "asdfasdfa"},
	{Id: "3", Title: "java 架构师", Detail: "asdfasdfa"},
}

//历史评价界面
type historyInfo struct {
	RoundIndex int    `json:"roundIndex"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	note       string `json:"note"`
}

var historyInfos = []historyInfo{
	{1, "2020-10-10", "2020-11-10", "第1次评价"},
	{2, "2021-10-10", "2021-11-10", "第2次评价"},
	{3, "2022-10-10", "2022-11-10", "第3次评价"},
	{4, "2023-10-10", "2023-11-10", "第4次评价"},
	{5, "2024-10-10", "2024-11-10", "第5次评价"},
	{6, "2025-10-10", "2025-11-10", "第6次评价"},
	{7, "2026-10-10", "2026-11-10", "第7次评价"},
	{8, "2027-10-10", "2027-11-10", "第8次评价"},
}
var allRounds = map[int][]empInfo{
	1: []empInfo{
		{"1111111", "小黑", "普通员工", "开发部"},
		{"2222222", "小红", "主管", "设计部"},
		{"3333333", "小蓝", "经理", "产品部"},
		{"4444444", "小紫", "老板", "管理部"},
		{"4444444", "小白", "ceo", "法律部"},
		{"4444444", "小绿", "ipo", "市场部"},
		{"4444444", "小灰", "组长", "开发部"},
		{"4444444", "小黄", "领导", "开发部"},
		{"4444444", "小橙", "普通员工", "开发部"},
		{"4444444", "小青", "普通员工", "开发部"},
	}, 2: []empInfo{
		{"5555555", "小黑", "普通员工", "开发部"},
		{"6666666", "小红", "主管", "设计部"},
		{"7777777", "小蓝", "经理", "产品部"},
		{"8888888", "小紫", "老板", "管理部"},
		{"9999999", "小白", "ceo", "法律部"},
		{"1010101", "小绿", "ipo", "市场部"},
		{"1212121", "小灰", "组长", "开发部"},
		{"1313131", "小黄", "领导", "开发部"},
		{"1414141", "小橙", "普通员工", "开发部"},
		{"1515151", "小青", "普通员工", "开发部"},
	},
	3: []empInfo{
		{"17171717", "小黑", "普通员工", "开发部"},
		{"18181818", "小红", "主管", "设计部"},
		{"19191919", "小蓝", "经理", "产品部"},
		{"20202020", "小紫", "老板", "管理部"},
		{"23232323", "小白", "ceo", "法律部"},
		{"24424242", "小绿", "ipo", "市场部"},
		{"25252525", "小灰", "组长", "开发部"},
	},
	4: []empInfo{
		{"44444444", "小黑", "普通员工", "开发部"},
	},
	5: []empInfo{
		{"55555555", "小黑", "普通员工", "开发部"},
	},
	6: []empInfo{
		{"6666666", "小黑", "普通员工", "开发部"},
	},
	7: []empInfo{
		{"7777777", "小黑", "普通员工", "开发部"},
	},
	8: []empInfo{
		{"8888888", "小黑", "普通员工", "开发部"},
	},
}
