package Service

import (
	"fmt"
	"time"
)

func TermCron() {
	fmt.Println("可以根据MatchPhaseQuery查询，写TermQuery")
	time.Sleep(time.Second * 1000)
}
