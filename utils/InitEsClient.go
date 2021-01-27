package utils

import (
	"fmt"
	"github.com/olivere/elastic/v7"

)

type EsClient struct {
	client *elastic.Client
}

func NewEsClient(Ei EsInfo) (client *elastic.Client, err error) {
	logger:=CreateLogger()
	client, err = elastic.NewClient(elastic.SetURL(Ei.Host...), elastic.SetSniff(false),
		elastic.SetBasicAuth(Ei.User, Ei.Passwd))
	if err != nil {
		err = fmt.Errorf("NewClien err:%s", err)
		return
	}
	v, err := client.ElasticsearchVersion(Ei.Host[0])
	if err != nil {
		err = fmt.Errorf("NewClien err:%s", err)
		return
	}
	logger.Printf("Elasticsearch Version: %s\n", v)
	return client, err
}
