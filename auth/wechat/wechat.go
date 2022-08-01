package wechat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service struct {
	AppID     string
	AppSecret string
}

type OpenIDResponse struct {
	ErrCode    int32  `json:"errorcode"`
	ErrMsg     string `json:"errmsg"`
	UnionID    string `json:"unionid"`
	SessionKey string `json:"session_key"`
	OpenID     string `json:"openid"`
}

func (s *Service) Resolve(code string) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", s.AppID, s.AppSecret, code)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var respJSON OpenIDResponse
	err = json.Unmarshal(body, &respJSON)
	if err != nil {
		return "", err
	}

	if respJSON.OpenID == "" {
		return "", fmt.Errorf("response code: %v, response error: %v", respJSON.ErrCode, respJSON.ErrMsg)
	}

	//map[errcode:41002 errmsg:appid missing, rid: 62de9999-02bc2880-3a17b2c4]
	//map[errcode:40029 errmsg:invalid code, rid: 62de9a0b-7149b5f1-2eee4908]
	//map[openid:okQ-s5X1ILvDud0amynR1Psw-URs session_key:bJb7LAI5uK8OoMPUHlYmrA== Unionid:oWXxQ6CFM9_9FqIknCWfwshZ_bwQ]

	return respJSON.OpenID, nil
}
