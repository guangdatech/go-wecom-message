# 介绍
本项目是利用企业微信群里的机器人，发送消息.
使用步骤：
1. 创建企业微信群
2. 给群添加一个聊天机器人
3. 获取到群机器人的webhook地址中的key，如:https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=693a91f6-7xxx-4bc4-97a0-0ec2sifa5aaa

# 安装
``` 
go get github.com/guangdatech/go-wecom-message
```

示例:
```
import (
	"fmt"
	"github.com/guangdatech/go-wecom-message/message"
	"testing"
)

func Test_Send(t *testing.T) {
	key := "xxxxx-8a9c-433e-82c8-xxxx" //企业微信群机器人的key
	if flag, err := message.Set(key).SendMarkdown("包测试"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("发送成功", flag)
	}
}

```

