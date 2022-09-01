package test

import (
	"github.com/loebfly/klogs"
	"testing"
)

func TestConsole(t *testing.T) {
	err := klogs.Init("")
	if err != nil {
		t.Error(err)
		return
	}
	klogs.WillOutputUse(func(c, level string) map[string]int {
		if c == "HTTP" {
			return map[string]int{
				"[TraceId:73523242]": 2,
			}
		}
		return nil
	})
	klogs.C("HTTP").Error("测试错误输出")
	klogs.Warn("测试:{},{}", "第一个告警输出", 1)
	klogs.C("MYSQL").Debug("调试:{}", "数据库链接异常")
}
