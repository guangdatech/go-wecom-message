package model

type SendRsp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type RoBotMsgBody struct {
	MsgType  string         `json:"msgtype"`
	Markdown *RobotMarkdown `json:"markdown"`
}

type RobotMarkdown struct {
	Content             string `json:"content"`
	MentionedList       string `json:"mentioned_list"`
	MentionedMobileList string `json:"mentioned_mobile_list"`
}
