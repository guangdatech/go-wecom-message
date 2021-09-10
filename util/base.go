package util

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/guangdatech/go-wecom-message/model"
)

type send struct {
	key string
}

var Send send

func Set(key string) *send {
	entity := &send{
		key: key,
	}
	return entity
}

// SendMarkdown send a message
func (s *send) SendMarkdown(msg string) (bool, error) {
	req := &model.RobotMarkdown{}
	req.Content = msg
	message := &model.RoBotMsgBody{
		MsgType:  "markdown",
		Markdown: req,
	}
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + s.key
	rsp, err := g.Client().Header(map[string]string{
		"Content-Type": "application/json",
	}).Post(url, gconv.String(message))
	sendRsp := &model.SendRsp{}
	if rsp != nil {
		defer rsp.Close()
	}
	if err != nil {
		return false, err
	}
	err = gconv.Scan(rsp.ReadAllString(), &sendRsp)
	if err != nil {
		return false, err
	}
	if sendRsp.ErrMsg == "ok" {
		return true, nil
	} else {
		return false, errors.New(sendRsp.ErrMsg)
	}

}
