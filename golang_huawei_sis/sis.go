package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	//"io/ioutil"
	"net/http"
)

//网络请求
//https://www.jb51.net/article/210846.htm

//华为云
//https://support.huaweicloud.com/api-sis/sis_03_0111.html

func main() {
	//testGet()
	//testPost()

	postSisToken()
	postSisVoice()
	//sisVoice()
}

//配置项
const voicepersion = chinese_xiaosong_common

//小琪，标准女声发音人。
const chinese_xiaoqi_common = "chinese_xiaoqi_common"

//小宇，标准男声发音人。
const chinese_xiaoyu_common = "chinese_xiaoyu_common"

//小燕，温柔女声发音人。
const chinese_xiaoyan_common = "chinese_xiaoyan_common"

//小王，童声发音人。
const chinese_xiaowang_common = "chinese_xiaowang_common"

//小雯，柔美女声发音人。
const chinese_xiaowen_common = "chinese_xiaowen_common"

//小婧，俏皮女声发音人。
const chinese_xiaojing_common = "chinese_xiaojing_common"

//小宋，激昂男声发音人。
const chinese_xiaosong_common = "chinese_xiaosong_common"

//小夏，热情女声发音人。
const chinese_xiaoxia_common = "chinese_xiaoxia_common"

//小呆，呆萌童声发音人。
const chinese_xiaodai_common = "chinese_xiaodai_common"

//小倩，成熟女声发音人。
const chinese_xiaoqian_common = "chinese_xiaoqian_common"

//cameal，柔美女声英文发音人。
const english_cameal_common = "english_cameal_common"

//精品
//华小夏，热情女声发音人。
const chinese_huaxiaoxia_common = "chinese_huaxiaoxia_common"

//华晓刚，利落男声发音人。
const chinese_huaxiaogang_common = "chinese_huaxiaogang_common"

//华小璐，知性女声发音人。
const chinese_huaxiaolu_common = "chinese_huaxiaolu_common"

//华小舒，舒缓女声发音人。
const chinese_huaxiaoshu_common = "chinese_huaxiaoshu_common"

//华小唯，嗲柔女声发音人。
const chinese_huaxiaowei_common = "chinese_huaxiaowei_common"

//华小靓，嘹亮女声发音人。
const chinese_huaxiaoliang_common = "chinese_huaxiaoliang_common"

//华晓东，成熟男声发音人。
const chinese_huaxiaodong_common = "chinese_huaxiaodong_common"

//华小颜，严厉女声发音人。
const chinese_huaxiaoyan_common = "chinese_huaxiaoyan_common"

//华小萱，台湾女声发音人。
const chinese_huaxiaoxuan_common = "chinese_huaxiaoxuan_common"

//华小雯，柔美女声发音人。
const chinese_huaxiaowen_common = "chinese_huaxiaowen_common"

//华晓阳，朝气男声发音人。
const chinese_huaxiaoyang_common = "chinese_huaxiaoyang_common"

//华小闽，闽南女声发音人。
const chinese_huaxiaomin_common = "chinese_huaxiaomin_common"

//华女侠，武侠女生发音人，只支持16k的采样率。
const chinese_huanvxia_literature = "chinese_huanvxia_literature"

//华晓悬，悬疑男声发音人，只支持16k的采样率。
const chinese_huaxiaoxuan_literature = "chinese_huaxiaoxuan_literature"

//voice请求链接
const url_voice = "https://sis-ext.cn-north-4.myhuaweicloud.com/v1/0e7a89f0bf8091292fb2c01efc8f8be0/tts"

//token请求链接
const url_token = "https://iam.cn-north-4.myhuaweicloud.com/v3/auth/tokens"

//临时token值
var token = "MIIUfwYJKoZIhvcNAQcCoIIUcDCCFGwCAQExDTALBglghkgBZQMEAgEwghKRBgkqhkiG9w0BBwGgghKCBIISfnsidG9rZW4iOnsiZXhwaXJlc19hdCI6IjIwMjEtMTItMTFUMDk6MTc6MzEuMzM5MDAwWiIsIm1ldGhvZHMiOlsicGFzc3dvcmQiXSwiY2F0YWxvZyI6W10sInJvbGVzIjpbeyJuYW1lIjoidGVfYWRtaW4iLCJpZCI6IjAifSx7Im5hbWUiOiJ0ZV9hZ2VuY3kiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9lY3Nfc3BvdF9pbnN0YW5jZSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2l2YXNfdmNyX3ZjYSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2FfY24tc291dGgtNGMiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9lY3Nfa2FlMSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2V2c19lc3NkIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfZHdzX3BvYyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2Nicl9maWxlIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfZWNzX2tjMV91c2VyX2RlZmluZWQiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9tZWV0aW5nX2VuZHBvaW50X2J1eSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX21hcF9ubHAiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9zaXNfc2Fzcl9lbiIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX3NhZF9iZXRhIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfc2VydmljZXN0YWdlX21ncl9kdG1fZW4iLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9ldnNfdm9sdW1lX3JlY3ljbGVfYmluIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfZGNzX2RjczItZW50ZXJwcmlzZSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX3ZjcCIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2N2ciIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX21hcyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX211bHRpX2JpbmQiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9laXBfcG9vbCIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2FfYXAtc291dGhlYXN0LTNkIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfcHJvamVjdF9kZWwiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9zaGFyZUJhbmR3aWR0aF9xb3MiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9jY2lfb2NlYW4iLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9jZXNfcmVzb3VyY2Vncm91cF90YWciLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9ldnNfcmV0eXBlIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfZWNzX2lyM3giLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9lbGJfZ3VhcmFudGVlZCIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2FfY24tc291dGh3ZXN0LTJiIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfY2llIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfc2ZzdHVyYm8iLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF92cGNfbmF0IiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfdnBuX3Znd19pbnRsIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfaHZfdmVuZG9yIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfYV9jbi1ub3J0aC00ZSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2FfY24tbm9ydGgtNGQiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9kYXl1X2RsbV9jbHVzdGVyIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfaW50bF9jb25maWd1cmF0aW9uIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfY2NlX21jcF90aGFpIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfY29tcGFzcyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX3NlcnZpY2VzdGFnZV9tZ3JfZHRtIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfYV9jbi1ub3J0aC00ZiIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX3Vnb19wdWJsaWN0ZXN0IiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfY3BoIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfZWNzX2dwdV9nNXIiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF93a3Nfa3AiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9jY2lfa3VucGVuZyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX3JpX2R3cyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX3ZwY19mbG93X2xvZyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2F0YyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2FhZF9iZXRhX2lkYyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2NzYnNfcmVwX2FjY2VsZXJhdGlvbiIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2Vjc19kaXNrQWNjIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfYWlzX2FwaV9pbWFnZV9hbnRpX2FkIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfZHNzX21vbnRoIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfY3NnIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfZGVjX21vbnRoX3VzZXIiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9pZWZfZWRnZWF1dG9ub215IiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfdmlwX2JhbmR3aWR0aCIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX29zYyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2Vjc19vbGRfcmVvdXJjZSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX3dlbGlua2JyaWRnZV9lbmRwb2ludF9idXkiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9kY3NfZGNzMi1yZWRpczYtZ2VuZXJpYyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2llZi1pbnRsIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfZWNzX2FsZyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2Z1bmN0aW9uZ3JhcGhfdjJfaW50bCIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX3BzdG5fZW5kcG9pbnRfYnV5IiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfbWFwX29jciIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2Rsdl9vcGVuX2JldGEiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9pZXMiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9vYnNfZHVhbHN0YWNrIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfZWRjbSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2NzYnNfcmVzdG9yZSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2l2c2NzIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfZWNzX2M2YSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX3Zwbl92Z3ciLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9zbW5fY2FsbG5vdGlmeSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2NhZS1iZXRhIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfY2NlX2FzbV9oayIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2NzYnNfcHJvZ3Jlc3NiYXIiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9ldnNfcG9vbF9jYSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2JjZSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2Vjc19vZmZsaW5lX2Rpc2tfNCIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2ludGxfY29tcGFzcyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2VwcyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2NzYnNfcmVzdG9yZV9hbGwiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9sMmNnIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfaW50bF92cGNfbmF0IiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfZmNzX3BheSIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2wyY2dfaW50bCIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2FfYXAtc291dGhlYXN0LTFlIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfYV9ydS1tb3Njb3ctMWIiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9hX2FwLXNvdXRoZWFzdC0xZCIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2FfYXAtc291dGhlYXN0LTFmIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfb3BfZ2F0ZWRfbWVzc2FnZW92ZXI1ZyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2Vjc19jNyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX21hcF92aXNpb24iLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9lY3NfcmkiLCJpZCI6IjAifSx7Im5hbWUiOiJvcF9nYXRlZF9hX2FwLXNvdXRoZWFzdC0xYyIsImlkIjoiMCJ9LHsibmFtZSI6Im9wX2dhdGVkX2FfcnUtbm9ydGh3ZXN0LTJjIiwiaWQiOiIwIn0seyJuYW1lIjoib3BfZ2F0ZWRfaWVmX3BsYXRpbnVtIiwiaWQiOiIwIn1dLCJwcm9qZWN0Ijp7ImRvbWFpbiI6eyJuYW1lIjoiaHdfMDA4NjE1NjUyOTE5MjUyXzAxIiwiaWQiOiIwZDQwNDI3ZWJhODBmNTgwMGY1MWMwMWVhYjM4MzcwMCJ9LCJuYW1lIjoiY24tbm9ydGgtNCIsImlkIjoiMGU3YTg5ZjBiZjgwOTEyOTJmYjJjMDFlZmM4ZjhiZTAifSwiaXNzdWVkX2F0IjoiMjAyMS0xMi0xMFQwOToxNzozMS4zMzkwMDBaIiwidXNlciI6eyJkb21haW4iOnsibmFtZSI6Imh3XzAwODYxNTY1MjkxOTI1Ml8wMSIsImlkIjoiMGQ0MDQyN2ViYTgwZjU4MDBmNTFjMDFlYWIzODM3MDAifSwibmFtZSI6InN1biIsInBhc3N3b3JkX2V4cGlyZXNfYXQiOiIiLCJpZCI6ImViZmI5MzlhNjZmNzRjNTE4MDA2ZGY3YTk4ZjhkNGE5In19fTGCAcEwggG9AgEBMIGXMIGJMQswCQYDVQQGEwJDTjESMBAGA1UECAwJR3VhbmdEb25nMREwDwYDVQQHDAhTaGVuWmhlbjEuMCwGA1UECgwlSHVhd2VpIFNvZnR3YXJlIFRlY2hub2xvZ2llcyBDby4sIEx0ZDEOMAwGA1UECwwFQ2xvdWQxEzARBgNVBAMMCmNhLmlhbS5wa2kCCQDcsytdEGFqEDALBglghkgBZQMEAgEwDQYJKoZIhvcNAQEBBQAEggEAi8kjOEKRIgMICNEzfZrJh1jcM6yxm01nudvKSWkFv3qXGRA2d1-Z7hzTkepoH7wqwy87Mx-FAUswSSZq06gyjaIib81wTMKPvmHtXMF5YDxlulZmzUN2Bvqd+6XFZKsMG2tApGiycUlrAudxpbo9YbVkakkFlebM5FKF2JBJ4rLJ-5IGx80mAKqDfCLs+ApKierpihqg35KS9yJN-Q+wabxmrdMXImrZkB1qlgGtfn5m1rBp0EEfLHibTmzZjO+HwwdGD61pk4W4b8PFUrYxUoZT4rAIzOo7TpJDDPtL4A28ZyvzqaRqMXZ0ZGjTIs4sF9CUY99tG1tenhwB31LHiw=="

//token请求参数
const request_token = `{ 
    "auth": { 
        "identity": { 
            "methods": [ 
                "password" 
            ], 
            "password": { 
                "user": { 
                    "name": "sun", 
                    "password": "gh426486", 
                    "domain": { 
                        "name": "hw_008615652919252_01" 
                    } 
                } 
            } 
        }, 
        "scope": { 
            "project": { 
                "name": "cn-north-4" 
            } 
        } 
    } 
}`

//token请求参数
const request_voice = `{ 
  "text": "` + voiceText + `",
   "config": 
   { 
     "audio_format": "wav", 
     "sample_rate": "16000", 
     "property": "` + voicepersion + `",
     "speed": 10,
     "pitch": 10,
     "volume": 60
   }
 }`

//获取token
func postSisToken() {
	jsonStr := []byte(request_token)
	req, err := http.NewRequest("POST", url_token, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	//statuscode := resp.StatusCode
	head := resp.Header
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//fmt.Println(statuscode)
	//fmt.Println(head)
	fmt.Println(head.Get("X-Subject-Token"))
	token = head.Get("X-Subject-Token")
}

//500个字符以上的分页上传,由于需要断句,不是很合适,暂时不这样做,功能未完成
func sisVoice() {

	var limitCount = 500
	//总字数
	//countAll := utf8.RuneCountInString(voiceText)
	countAll := len(voiceText)
	fmt.Printf(`总长度:%d`, countAll)
	fmt.Printf("\n")

	//够几个500?
	count := countAll/limitCount + 1

	for i := 0; i < count; i++ {
		fmt.Printf(`i:%d,c:%d`, i, count)

		startIndex := limitCount * i
		var stopIndex = limitCount * (i + 1)
		if i == (count - 1) {
			stopIndex = countAll
		}

		textLen := voiceText[startIndex:stopIndex]
		fmt.Printf(`%d长度:%d`, i, len(textLen))
		fmt.Printf("\n")
		fmt.Println(textLen)
	}

}

//获取voice
func postSisVoice() {
	jsonStr := []byte(request_voice)
	req, err := http.NewRequest("POST", url_voice, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	//statuscode := resp.StatusCode
	//head := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	//fmt.Println(statuscode)
	//fmt.Println(head)
	//fmt.Println(head.Get("X-Subject-Token"))

	//var result map[string]interface{}
	//err = json.Unmarshal(body, &result)
	//
	//fmt.Println("------------")
	////fmt.Println(result["data"])
	//
	//var data1 = result["result"]
	//data1[]
	//var data = ""

	var data Baseentity
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return
	}

	decodeBytes, err := base64.StdEncoding.DecodeString(data.Result.Data)
	write(decodeBytes)

	//fmt.Println(decodeBytes)
	//fmt.Println(string(decodeBytes))

}

//写入文件
func write(byte []byte) {
	f, err := os.Create("test.wav")
	if err != nil {
		fmt.Println(err)
		return
	}
	n2, err := f.Write(byte)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

type Baseentity struct {
	Traceid string  `json:"trace_id"`
	Result  Reldata `json:"result"`
}

type Reldata struct {
	Data string `json:"data"`
}

//func testGet() {
//
//		resp, err := http.Get("http://www.hao123.com")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer resp.Body.Close()
//		body, err := ioutil.ReadAll(resp.Body)
//		fmt.Println(string(body))
//		fmt.Println(resp.StatusCode)
//		if resp.StatusCode == 200 {
//			fmt.Println("ok")
//		}
//
//}
//
//func testPost() {
//	urlValues := url.Values{}
//	urlValues.Add("broker_id","8323003")
//	//urlValues.Add("age","26")
//	resp, _ := http.PostForm("https://appapi.5i5j.com/im/brokermessage",urlValues)
//	body, _ := ioutil.ReadAll(resp.Body)
//	fmt.Println(string(body))
//}
