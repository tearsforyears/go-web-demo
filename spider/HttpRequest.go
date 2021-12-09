package spider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GenSpider() {
	par := map[string]interface{}{}
	res := GetBody("https://www.tianqiapi.com/api?version=v1&appid=21375891&appsecret=fTYv7v5E")
	json.Unmarshal([]byte(res), &par)
	data := par["data"].([]interface{})[0]
	row := data.(map[string]interface{})
	fmt.Println(par["city"], "空气质量:", row["air_level"], "空气质量指数:", row["air"], "\ntips:", row["air_tips"])
	fmt.Println("温度:", row["tem1"], "~", row["tem2"])
	fmt.Println("日落 日出:", row["sunset"], " ", row["sunrise"])
	fmt.Println("风:", row["win_speed"], ",", row["win_meter"])
}

func GetBody(url string) string {
	res, _ := http.Get(url)
	r, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return string(r)
}
