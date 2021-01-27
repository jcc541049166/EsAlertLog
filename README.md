### EsAlertLog

###### Usage:
应用程序日志收集到ES，查找相关关键字，达到多少笔进行邮件告警。

```
go run main.go -c D:\goproject\src\EsAlertLog\config\es.conf-f D:\goproject\src\EsAlertLog\config\alert.conf -m D:\goproject\src\EsAlertLog\config\mail.conf
go build -o esalertlog.exe .\EsAlertLog\main\main.go
```

###### es.conf:
```
{
    "host":["http://192.168.168.168:9200"],
    "user":"elastic",
    "passwd":"123456"
}
```
###### alert.conf
```
[
    {
        "ruleid": 1,                                           //规则ID
        "ruleindex": "index1",                                 //要检查得索引
        "rulefield": "msg",                                    //检查得字段
        "timefield": "time",                                   //要检查得时间字段
        "rulekeyword": "love",                                 //要检查得关键字
        "recvmail": ["abc@admin.com"],                         //收件人邮箱列表
        "mailsubject":"index1 alert",                          //邮件主题
        "mailbody": "D:/goproject/src/EsAlertLog/config/mailtemplate.html",           //邮件模板 
        "threshold": 2,                                        //关键字日志超过多少比，发生告警
        "interval": "30s",                                     //多久检查一次, 也是只检查距离当前时间多久得日志
        "type":"MatchPhase"                                    //检查类型，短语匹配
    }
]
```
###### mail.conf
```
{
		"user": "abc@admin.com",      //发件人邮箱
		"pass": "11111111",           //发件人邮箱密码
		"host": "smtp.admin.com",     //发件箱地址
		"port": "465"                 
}
```
###### log
```
./log
```