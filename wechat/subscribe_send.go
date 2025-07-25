package wechat

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"time"
)

type SubscribeSend struct {
	TemplateId       string         `json:"template_id"`
	Page             string         `json:"page,omitempty"`
	ToUser           string         `json:"touser"`
	Data             map[string]any `json:"data"`
	MiniprogramState string         `json:"miniprogram_state"`
	Lang             string         `json:"lang,omitempty"`
}

type SubscribeSendResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func Subscribe(send *SubscribeSend, accessToken string) (*SubscribeSendResponse, error) {

	baseURL := "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=" + accessToken

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	encodedData, err := json.Marshal(send)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(encodedData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var response SubscribeSendResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
