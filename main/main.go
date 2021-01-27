package main

import (
	"EsAlertLog/Service"
	"EsAlertLog/utils"
	"flag"
	"fmt"
)

func Processor() {
	logger:=utils.CreateLogger()
	espath := flag.String("c", "", "Elasticsearch连接配置文件路径")
	rulepath := flag.String("f", "", "告警规则配置文件路径")
	mailpath := flag.String("m", "", "发件箱信息")
	flag.Parse()
	Ei, err := utils.NewEsInfo(*espath)
	if err != nil {
		logger.Panicln(err)
	}
	Rsi, err := utils.NewRulesInfo(*rulepath)
	if err != nil {
		logger.Panicln(err)
	}
	Mi, err := utils.ParseMailconf(*mailpath)
	if err != nil {
		logger.Panicln(err)
	}
	newclient, err := utils.NewEsClient(Ei)
	if err!=nil{
		logger.Panicln(err)
	}
	ResChan := make(chan utils.ResultInfo, 100)
	esprocessclient := &Service.ProcessClient{
		Pclient: newclient,
		ResChan:ResChan,
	}
	esprocessclient.Process(Rsi)
	Service.SendMail(ResChan, Mi)
}

func main() {
	//解析Flag参数
	fmt.Println("开始执行......")
	logger:=utils.CreateLogger()
	logger.Println("开始执行......")
	Processor()
}
