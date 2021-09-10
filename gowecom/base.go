package gowecom

import (
	"encoding/json"
	"errors"
	"github.com/guangdatech/go-wecom-message/library"
	"github.com/guangdatech/go-wecom-message/model"
)

type send struct {
	key string
}


func Set(key string) *send {
	entity := &send{
		key: key,
	}
	return entity
}

// SendMarkdown send a message
// 参见: https://work.weixin.qq.com/api/doc/90000/90136/91770
func (s *send) SendMarkdown(msg string) error {
	req := &model.RobotMarkdown{}
	req.Content = msg
	message := &model.RoBotMsgBody{
		MsgType:  "markdown",
		Markdown: req,
	}
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + s.key
	return Send(url,message)

}

//Send 发送请求
func Send(url string ,message *model.RoBotMsgBody) error{
	rsp,err:=library.Post(url, message,"application/json")

	if err!=nil{
		return err
	}

	sendRsp:=&model.SendRsp{}
	if err=json.Unmarshal([]byte(rsp),&sendRsp);err!=nil{
		return err
	}

	if sendRsp.ErrMsg == "ok" {
		return  nil
	} else {
		return  errors.New(sendRsp.ErrMsg)
	}

}


