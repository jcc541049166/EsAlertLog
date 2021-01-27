package utils

import (
	"encoding/json"
	"gopkg.in/gomail.v2"
	"io/ioutil"
)

type Mailclient struct {
	mc *gomail.Dialer //mc  Mail Dialer Client
}

type MailInfo struct {
	User string `json:"user"` //发件人邮箱
	Pass string `json:"pass"` //发件人邮箱密码
	Host string `json:"host"` //发件邮箱
	Port string `json:"port"` //邮箱端口
}

func ParseMailconf(fpath string) (Mi MailInfo, err error) {
	Mi = MailInfo{}
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &Mi)
	return
}

func ParseMailbody(fpath string) (str string, err error) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return
	}
	str = string(b)
	return
}
