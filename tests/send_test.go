package tests

import (
	"fmt"
	"github.com/guangdatech/go-wecom-message/util"
	"testing"
)

func Test_Send(t *testing.T) {
	key := "f0c332c7-8a9c-433e-82c8-304xxxxxx"
	if flag, err := util.Set(key).SendMarkdown("包测试"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("测试失败", flag)
	}
}
