// 调用api
package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Api struct {
	baseUrl string
}

func NewApi(baseUrl string) Api {
	return Api{baseUrl}
}

func post(url string, param map[string]interface{}) []byte {
	client := &http.Client{}
	jParam, _ := json.Marshal(param)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(jParam)))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return buffer
}

func (api *Api) PostReturnMap(path string, param map[string]interface{}) map[string]interface{} {
	url := api.baseUrl + "/" + path
	buffer := post(url, param)

	respBody := map[string]interface{}{}
	json.Unmarshal(buffer, &respBody)
	return respBody
}

func (api *Api) PostReturnList(path string, param map[string]interface{}) []map[string]interface{} {
	url := api.baseUrl + "/" + path
	buffer := post(url, param)

	respBody := make([]map[string]interface{}, 0)
	json.Unmarshal(buffer, &respBody)
	return respBody
}
