package Service

import (
	"EsAlertLog/utils"
	"context"
	"github.com/olivere/elastic/v7"
	"strconv"
	"time"
)

type MatchPhaseProcess struct {
	client *elastic.Client
	ResChan chan utils.ResultInfo
	Ri utils.RuleInfo
}

func (mpp *MatchPhaseProcess)MatchPhaseCron() (err error) {
	logger:=utils.CreateLogger()
	i, err := time.ParseDuration(mpp.Ri.Interval)
	if err != nil {
		logger.Panicf("Interval转time Duration失败,err: %s",err)
	}
	t := time.NewTicker(i)
	for {
		select {
		case _ = <-t.C: //定时任务执行
			mpp.MatchPhaseProcess(i)
		}
	}
}


//MatchPhase 查询Es
func (mpp *MatchPhaseProcess)MatchPhaseProcess(i time.Duration) {
	//定义查询条件,查询时间范围内,指定短语匹配
	logger:=utils.CreateLogger()
	boolq := elastic.NewBoolQuery()
	boolq.Must(elastic.NewMatchPhraseQuery(mpp.Ri.Rulefield, mpp.Ri.Rulekeyword))
	boolq.Must(elastic.NewRangeQuery(mpp.Ri.Timefield).Gte(time.Now().Add(-i).Format("2006/01/02 15:04:05")).Lte(time.Now().Format("2006/01/02 15:04:05")))
	//执行查询
	searchdata, err := mpp.client.Search().Index(mpp.Ri.Ruleindex).Query(boolq).Size(1).Do(context.Background())
	logger.Printf("MatchPhase Status: %d,TimOut: %v, ToTal Hits: %d, TookInMillis: %d\n",searchdata.Status, searchdata.TimedOut,searchdata.Hits.TotalHits.Value,searchdata.TookInMillis)
	if err != nil {
		logger.Panicf("MatchPhase Process查询搜索出错, err: %s ", err)
	}
	var res utils.ResultInfo
	c, err := strconv.Atoi(strconv.FormatInt(searchdata.TotalHits(), 2))
	if err != nil {
		logger.Panicf("TotalHits数据转换错误, err:%s ", err)
	}
	res.Count = c
	res.RuleInfo = mpp.Ri
	mpp.ResChan <- res
}
