package test

import (
	"fmt"
	"github.com/serge-xu/go-vue3-project/config"
)

func TestStart() {
	testLog()
}

func testLog() {
	fmt.Printf("testLog is nil %v\n", config.Logger == nil)
	config.Logger.Debugf("这是一条调试日志")
	config.Logger.Infof("这是一条信息日志")
	config.Logger.Warnf("这是一条警告日志")
	config.Logger.Errorf("这是一条错误日志")
}
