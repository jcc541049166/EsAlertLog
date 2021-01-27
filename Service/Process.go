package Service

import (
	"EsAlertLog/utils"
	"github.com/olivere/elastic/v7"
)
type ProcessClient struct {
	Pclient *elastic.Client
	ResChan chan utils.ResultInfo
}

//控制任务分发
func (pc *ProcessClient)Process(Rsi []utils.RuleInfo) {
	logger:=utils.CreateLogger()
	for _, Ri := range Rsi {
		//根据查询规则(TypeMatchPhase,TypeTerm等),按协程分发任务
		switch Ri.Type {
		case utils.TypeMatchPhase:
			ep:=&MatchPhaseProcess{
				client: pc.Pclient,
				ResChan: pc.ResChan,
				Ri: Ri,
			}
			go ep.MatchPhaseCron()
		case utils.TypeTerm:
			go TermCron()
		default:
			logger.Println("存在查询规则中,搜索类型不存在.")
			continue
		}
	}

}
