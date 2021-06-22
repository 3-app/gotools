package estool

import (
	"github.com/webchen/gotools/base/conf"
	"github.com/webchen/gotools/help/logs"

	"github.com/elastic/go-elasticsearch/v8"
)

var esList = make(map[string]*elasticsearch.Client)

func init() {
	var es *elasticsearch.Client
	serverList := make(map[string]map[string]interface{})
	serverList = conf.GetConfig("es", serverList).(map[string]map[string]interface{})

	for k, v := range serverList {
		host := v["host"].([]interface{})
		var hostList []string
		for _, v := range host {
			hostList = append(hostList, v.(string))
		}
		user := v["user"].(string)
		password := v["password"].(string)
		cfg := elasticsearch.Config{
			Addresses: hostList,
			Username:  user,
			Password:  password,
			// ...
		}
		var err error
		es, err = elasticsearch.NewClient(cfg)
		if logs.ErrorProcess(err, "无法初始化ES") {
			continue
		}
		esList[k] = es
	}
}

/*
// WriteLog 往ES里面写LOG
func WriteLog(level string, message string, v ...interface{}) {
	index := (conf.GetConfig("es.index", "gateway_pub")).(string)
	go (func() {
		data := map[string]interface{}{
			"@timestamp": time.Now().Format(time.RFC3339Nano),
			"level":      level,
			"ip":         nettool.GetLocalFirstIPStr(),
			"message":    message,
			"content":    v,
		}
		body := jsontool.MarshalToString(data)
		req := esapi.IndexRequest{
			Index:   index,
			Body:    bytes.NewReader([]byte(body)),
			Refresh: "true",
		}
		res, err := req.Do(context.Background(), es)
		if err != nil || res == nil {
			log.SetPrefix("ESERROR")
			log.Printf("write log error [%+v] [%+v]", data, err)
			return
		}
		defer res.Body.Close()
		if strings.Contains(res.String(), "error") {
			log.Println(res.String())
		}
	})()
}
*/
