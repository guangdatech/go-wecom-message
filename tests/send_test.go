package tests

import (
	"fmt"
	"github.com/guangdatech/go-wecom-message/gowecom"
	"testing"
)

func Test_Send(t *testing.T) {
	key := "f0c332c7-8a9c-433e-82c8-30469ca87xxx"
	if err := gowecom.Set(key).SendMarkdown("hello word"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("测试成功")
	}
}
