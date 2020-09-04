package toLogCenter

import (
	"testing"
)


func TestAAA(t *testing.T) {
	//AAA()
	//fmt.Println(time.Now().Format(time.RFC3339Nano))
	logger := &Logger{
		Project:   "wikenTest",
		Module:    "test1",
		User:      "7921",
		Field:     nil,
	}
	defer func() {
		logger.Send()
	}()
	logger.LogURL = "http://log.ys.com/log/save"

	logger = logger.New()
	//logger.Print( "wikenvalue")
	//logger.Print("wikenvalue1")
	//logger.Print( "wikenvalue2")
	//logger.Print("wikenvalue3")
	//logger.Print( "wikenvalue5")
	//logger.Print(map[string]string{
	//	"aa": "aa",
	//	"bb": "aa",
	//	"cc": "aa",
	//	"add": "aa",
	//	"aas": "aa",
	//})
	//logger.Print("aaaaa", "bbb")
	logger.Printf("aaaaa%s", "bbb")
	logger.Print(map[string]string{
		"aa": "aa",
		"bb": "aa",
		"cc": "aa",
		"add": "aa",
		"aas": "aa",
	}, "请求异常")
	//logger.Printf("sss%s", "bbb", "cccc")
	//logger.PrintInput("aaaaaaaaaaaaaaaaaa")
	//logger.PrintReturn("bbbbbbbbbbbbbbbbbbb")
	//logger.level = "info"
	//logger.AddField(0, "AA")
	//logger.AddField(1, "BB")
	//logger.AddField(2, "CC")

}