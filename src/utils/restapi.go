package utils

import (
	"bytes"
	"configfile"
	"encoding/json"
	"fmt"
	"gopkg.in/macaron.v1"
	"io/ioutil"
	"net/http"
	"strconv"
	"utils/log"
)

func RestGet(url string, token string, jwt string) (response map[string]interface{}, code int, err error) {
	code = 500
	// Create Client
	client := &http.Client{}
	// Set request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err.Error())
		return
	}
	if token != "" {
		request.Header.Add("token", token)
	}
	if jwt != "" {
		request.Header.Add("authorization", jwt)
	}
	// Do request
	res, err := client.Do(request)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer res.Body.Close()
	code = res.StatusCode
	// Read body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err.Error())
		return
	}
	// Change to json
	//log.Debug(string(body))
	json.Unmarshal(body, &response)
	if res.StatusCode != 200 {
		if response["error"] != nil {
			err = fmt.Errorf(response["error"].(string))
			return
		}
		//log.Error(response)
		err = fmt.Errorf("status code is " + strconv.Itoa(res.StatusCode))
		return
	}
	//log.Debug(response)
	return
}

func RestPost(url string, data map[string]interface{}, token string, jwt string) (response map[string]interface{}, code int, err error) {
	// default value
	code = 500
	// Create client
	client := &http.Client{}
	// Change body to []byte
	b, err := json.Marshal(data)
	if err != nil {
		log.Error(err.Error())
		return
	}
	body := bytes.NewBuffer([]byte(b))
	// Set request
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Error(err.Error())
		return
	}
	// Set header
	if token != "" {
		request.Header.Add("token", token)
	}
	if jwt != "" {
		request.Header.Add("authorization", jwt)
	}
	request.Header.Set("Content-Type", "application/json")
	// Do request
	res, err := client.Do(request)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer res.Body.Close()
	code = res.StatusCode
	// read body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err.Error())
		return
	}

	json.Unmarshal(resBody, &response)
	if res.StatusCode != 200 {
		if response["error"] != nil {
			err = fmt.Errorf(response["error"].(string))
			return
		}
		log.Error(response)
	}
	//log.Debug(response)
	return
}

func RestPut(url string, data map[string]interface{}, token string, jwt string) (response map[string]interface{}, code int, err error) {
	code = 500
	// Create client
	client := &http.Client{}
	// Change body to []byte
	b, err := json.Marshal(data)
	if err != nil {
		log.Error(err.Error())
		return
	}
	body := bytes.NewBuffer([]byte(b))
	// Set request
	request, err := http.NewRequest("PUT", url, body)
	if err != nil {
		log.Error(err.Error())
		return
	}
	// Set header
	if token != "" {
		request.Header.Add("token", token)
	}
	if jwt != "" {
		request.Header.Add("authorization", jwt)
	}
	request.Header.Set("Content-Type", "application/json")
	// Do request
	res, err := client.Do(request)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer res.Body.Close()
	code = res.StatusCode
	// read body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err.Error())
		return
	}
	json.Unmarshal(resBody, &response)

	if res.StatusCode != 200 {
		if response["error"] != nil {
			err = fmt.Errorf(response["error"].(string))
			return
		}
		log.Error(response)
	}
	//log.Debug(response)
	return
}

func RestDelete(url string, data map[string]interface{}, token string, jwt string) (response map[string]interface{}, code int, err error) {
	code = 500
	// Create client
	client := &http.Client{}
	// Change body to []byte
	b, err := json.Marshal(data)
	if err != nil {
		log.Error(err.Error())
		return
	}
	body := bytes.NewBuffer([]byte(b))
	// Set request
	request, err := http.NewRequest("DELETE", url, body)
	if err != nil {
		log.Error(err.Error())
		return
	}
	// Set header
	if token != "" {
		request.Header.Add("token", token)
	}
	if jwt != "" {
		request.Header.Add("authorization", jwt)
	}
	request.Header.Set("Content-Type", "application/json")
	// Do request
	res, err := client.Do(request)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer res.Body.Close()
	code = res.StatusCode
	// read body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err.Error())
		return
	}
	json.Unmarshal(resBody, &response)
	if res.StatusCode != 200 {
		log.Debug(response["error"])
		if response["error"] != nil {
			err = fmt.Errorf(response["error"].(string))
			return
		}
		log.Error(response)
	}
	return
}

func ReturnErrorContext(ctx *macaron.Context, code int, err string) {
	Info := make(map[string]interface{}, 0)
	Info["error"] = err
	Info["success"] = false
	Info["data"] = nil
	log.Error(Info)
	ctx.JSON(code, Info)
	return
}

func ReturnSuccessContext(ctx *macaron.Context, code int, data map[string]interface{}) {
	Info := make(map[string]interface{}, 0)
	Info["error"] = nil
	Info["success"] = true
	Info["data"] = data
	ctx.JSON(code, Info)
	return
}

func GetPostInfo(ctx *macaron.Context) (args map[string]interface{}, err error) {
	args = make(map[string]interface{}, 0)
	bodyInByte, err := ctx.Req.Body().Bytes()
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyInByte, &args)
	return
}

func GenerateNewId(jwt string) (uuid string, err error) {
	resp, code, err := RestGet(configfile.IdentityUrl+"/auth/entity/newid", "", jwt)
	if err != nil {
		log.Error(err.Error())
		return
	}
	if code != 200 {
		log.Error("Code is not 200, current code is ", code)
		err = fmt.Errorf("Http code is not 200. ")
		return
	}
	uuid = resp["data"].(map[string]interface{})["uuid"].(string)
	return

}
