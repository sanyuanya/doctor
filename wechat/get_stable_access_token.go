package wechat

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
)

type getStableAccessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type getStableAccessTokenRequest struct {
	AppID        string `json:"appid"`
	Secret       string `json:"secret"`
	GrantType    string `json:"grant_type"`
	ForceRefresh bool   `json:"force_refresh,omitempty"`
}

func GetStableAccessToken(appid string, secret string) (*getStableAccessTokenResp, error) {
	baseURL := "https://api.weixin.qq.com/cgi-bin/stable_token"

	payload := &getStableAccessTokenRequest{
		AppID:        appid,
		Secret:       secret,
		GrantType:    "client_credential",
		ForceRefresh: false,
	}

	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	// 创建一个自定义的 http.Client，跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Post(baseURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	getStableAccessTokenResp := &getStableAccessTokenResp{}
	if err := json.NewDecoder(resp.Body).Decode(&getStableAccessTokenResp); err != nil {
		log.Fatal(err)
	}

	return getStableAccessTokenResp, nil
}
