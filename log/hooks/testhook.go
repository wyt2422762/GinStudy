package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// 实现自己的logrus hook 需要实现以下的接口
/* type Hook interface {
	Levels() []Level // 返回对应日志级别，输出其他日志时不会触发钩子
	Fire(*Entry) error // 日志输出前要执行大的操作
} */

type testHook struct {
	wyt string
}

func NewWytHook(wyt string) *testHook {
	return &testHook{
		wyt: wyt,
	}
}

func (th *testHook) Levels() []logrus.Level {
	//对应所有日志级别
	return logrus.AllLevels
}

func (th *testHook) Fire(entry *logrus.Entry) error {
	fmt.Println("test hook")
	// 日志中添加wyt字段
	entry.Data["wyt"] = th.wyt
	return nil
}
