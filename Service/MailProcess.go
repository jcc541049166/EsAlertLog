package Service

import (
	"EsAlertLog/utils"
	"gopkg.in/gomail.v2"
	_ "io/ioutil"
	"strconv"
	"strings"
)

func MailProcess(i utils.ResultInfo, Mi utils.MailInfo) {
	logger:=utils.CreateLogger()
	//Mail Server
	mailConn := map[string]string{
		"user": Mi.User,
		"pass": Mi.Pass,
		"host": Mi.Host,
		"port": Mi.Port,
	}
	port, err := strconv.Atoi(mailConn["port"])
	if err != nil {
		logger.Panicf("转换port失败,err: %s", err)
	}
	mc := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	mm := gomail.NewMessage()
	//Mail Headers
	mm.SetHeaders(map[string][]string{
		"From":    {mm.FormatAddress(mailConn["user"], "【日志告警】")},
		"To":      i.RuleInfo.Recvmail,
		"Subject": {i.RuleInfo.Mailsubject},
	})
	mailbody, err := utils.ParseMailbody(i.RuleInfo.Mailbody)
	if err != nil {
		logger.Panicf("ParseMailbody 解析mail body 失败, err: %s", err)
	}
	newmailbody := strings.Replace(mailbody, "$COUNT", strconv.Itoa(i.Count), -1)
	newmailbody = strings.Replace(newmailbody, "$KEYWORD", i.RuleInfo.Rulekeyword, -1)
	mm.SetBody("text/html", newmailbody)
	// mm.Attach("D:\\aaa.txt")   //可以添加附件功能
	//Send Mail
	err = mc.DialAndSend(mm)
}

func SendMail(ResChan chan utils.ResultInfo, Mi utils.MailInfo) {
	logger:=utils.CreateLogger()
	for i := range ResChan {
		if i.Count >= i.RuleInfo.Threshold {
			logger.Printf("需要发送邮件,Count数量为:%d\n", i.Count)
			MailProcess(i, Mi)
			logger.Printf("规则ID: %d,规则索引: %s, 搜索关键字: %s,搜索阈值: %d\n", i.RuleInfo.Ruleid, i.RuleInfo.Ruleindex, i.RuleInfo.Rulekeyword, i.RuleInfo.Threshold)
		} else {
			logger.Printf("不需要发送邮件,Count数量为:%d\n", i.Count)
			logger.Printf("规则ID: %d,规则索引: %s, 搜索关键字: %s,搜索阈值: %d\n", i.RuleInfo.Ruleid, i.RuleInfo.Ruleindex, i.RuleInfo.Rulekeyword, i.RuleInfo.Threshold)
		}
	}
}
