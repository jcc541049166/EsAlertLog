package utils

import (
	"encoding/json"
	"io/ioutil"
)

const (
	TypeMatchPhase = "MatchPhase"
	TypeTerm       = "Term"
)

//ES连接信息
type EsInfo struct {
	Host   []string `json:"host"`
	User   string   `json:"user"`
	Passwd string   `json:"passwd"`
}

//告警规则信息
type RuleInfo struct {
	Ruleid      int      `json:"ruleid"`      //告警规则ID
	Ruleindex   string   `json:"ruleindex"`   //查询的索引
	Rulefield   string   `json:"rulefield"`   //查询关键字 字段 MatchPhraseQuery
	Timefield   string   `json:"timefield"`   //查询时间   字段 RangeQuery
	Rulekeyword string   `json:"rulekeyword"` //查询的关键字,bool查询,NewMatchPhraseQuery
	Recvmail    []string `json:"recvmail"`    //收件人信息
	Mailsubject string   `json:"mailsubject"` //邮件主题
	Mailbody    string   `json:"mailbody"`    //html 邮件模板路径
	Threshold   int      `json:"threshold"`   //告警阈值
	Interval    string   `json:"interval"`    //轮询间隔,轮询的查询时间间隔, 也是查询最近轮询间隔时间的数据
	Type        string   `json:"type"`        //查询类型
}

//解析告警规则信息
func NewEsInfo(fpath string) (Ei EsInfo, err error) {
	Ei = EsInfo{}
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &Ei)
	return
}

//解析告警规则信息
func NewRulesInfo(fpath string) (Rsi []RuleInfo, err error) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return
	}
	Rsi = []RuleInfo{}
	err = json.Unmarshal(b, &Rsi)
	return
}
