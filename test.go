package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var x string
	x = "{\"Optcode\":\"test_play_music\",\"Value\":{\"Pattern1\":{\"Type\":\"PL\", \"Svolume\": [\"1\",\"2\"] },\"Pattern2\":{\"Type\":\"PL\", \"Svolume\": [1,2,3,4,5] }}}"
	y := make(map[string]interface{})
	err := json.Unmarshal([]byte(x), &y)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(x)
	fmt.Println(y)
	a := y["Value"].(map[string]interface{})
	b := a["Pattern1"].(map[string]interface{})
	c := b["Svolume"].([]interface{})
	d := b["Type"].(string)
	e, ok := a["key"].(string)
	if ok {
		fmt.Println(e)
	} else {
		fmt.Println("error")
	}
	zonelist := make([]string, len(c))
	fmt.Println(d)
	for key, value := range c {
		zonelist[key] = value.(string)
	}
	fmt.Println(zonelist)
	for _, value := range a {
		pattern := value.(map[string]interface{})
		fmt.Println(pattern)
	}
	/*data := `{"optCode": "Zone.getZoneIds", "value": ""}`
	request, _ := http.NewRequest("POST", "http://app-service-resource:49710/api/v2/trigger", strings.NewReader(data))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	//post数据并接收http响应
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("post data error:%v", err)
	} else {
		log.Println("post a data successful.")
		respBody, _ := ioutil.ReadAll(resp.Body)
		log.Printf("response data:%v", string(respBody))
		data := make(map[string]interface{})
		err = json.Unmarshal(respBody, &data)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(data)
		}
	}*/
}
